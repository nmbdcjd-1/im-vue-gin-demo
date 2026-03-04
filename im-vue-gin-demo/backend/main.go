package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"im-vue-gin-demo/backend/handlers"
	"im-vue-gin-demo/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MySQLDSN  = "root:wch123456.@tcp(8.129.109.155:3306)/im_db?charset=utf8mb4&parseTime=True&loc=Local"
	RedisAddr = "8.129.109.155:6379"
	RedisPass = "wch123456"
)

var (
	DB       *gorm.DB
	RDB      *redis.Client
	ctx      = context.Background()
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Message)
	mutex     = &sync.Mutex{}
)

func initDB() {
	var err error
	tempDSN := "root:wch123456.@tcp(8.129.109.155:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	tempDB, err := gorm.Open(mysql.Open(tempDSN), &gorm.Config{})
	if err == nil {
		tempDB.Exec("CREATE DATABASE IF NOT EXISTS im_db")
	}

	DB, err = gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("MySQL connect error: %v", err)
	}
	log.Println("MySQL connected")

	DB.AutoMigrate(&models.User{}, &models.Friend{}, &models.Message{}, &models.Group{}, &models.GroupMember{})
	handlers.DB = DB

	RDB = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPass,
		DB:       0,
	})

	if _, err := RDB.Ping(ctx).Result(); err != nil {
		log.Printf("Redis connect error: %v", err)
	} else {
		log.Println("Redis connected")
	}
}

func handleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	initDB()
	go handleMessages()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}

	r.GET("/ws", handleWebSocket)

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
