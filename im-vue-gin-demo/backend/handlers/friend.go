package handlers

import (
	"im-backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// AddFriend 发送好友申请
func AddFriend(c *gin.Context) {
	var input struct {
		TargetEmail string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的邮箱喵~"})
		return
	}

	var target models.User
	if err := DB.Where("email = ?", input.TargetEmail).First(&target).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到该用户喵~"})
		return
	}

	// 这里假设从 context 获取当前用户 ID，演示先写死
	currentUserID := uint(1) 

	friendReq := models.Friend{
		UserID:   currentUserID,
		FriendID: target.ID,
		Status:   0, // pending
	}

	if err := DB.Create(&friendReq).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "申请发送失败喵~"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "申请已发送，等待对方同意喵~"})
}

// GetFriends 获取好友列表
func GetFriends(c *gin.Context) {
	currentUserID := uint(1)
	var friends []models.Friend
	DB.Where("user_id = ? AND status = ?", currentUserID, 1).Find(&friends)

	var friendUsers []models.User
	for _, f := range friends {
		var u models.User
		DB.First(&u, f.FriendID)
		friendUsers = append(friendUsers, u)
	}

	c.JSON(http.StatusOK, friendUsers)
}
