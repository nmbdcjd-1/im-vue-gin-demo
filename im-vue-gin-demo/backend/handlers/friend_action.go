package handlers

import (
	"im-backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetFriendRequests 获取收到的好友申请
func GetFriendRequests(c *gin.Context) {
	currentUserID := uint(1) // 演示暂定
	var requests []models.Friend
	DB.Where("friend_id = ? AND status = ?", currentUserID, 0).Find(&requests)

	type ReqInfo struct {
		ID        uint   `json:"id"`
		FromEmail string `json:"from_email"`
		Nickname  string `json:"nickname"`
	}
	var res []ReqInfo
	for _, r := range requests {
		var u models.User
		DB.First(&u, r.UserID)
		res = append(res, ReqInfo{
			ID:        r.ID,
			FromEmail: u.Email,
			Nickname:  u.Nickname,
		})
	}
	c.JSON(http.StatusOK, res)
}

// HandleFriendRequest 处理好友申请 (同意/拒绝)
func HandleFriendRequest(c *gin.Context) {
	var input struct {
		RequestID uint `json:"id" binding:"required"`
		Action    int  `json:"action" binding:"required"` // 1: 同意, 2: 拒绝
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误喵"})
		return
	}

	var req models.Friend
	if err := DB.First(&req, input.RequestID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该申请喵"})
		return
	}

	if input.Action == 1 {
		req.Status = 1
		DB.Save(&req)
		// 双向添加好友
		DB.Create(&models.Friend{
			UserID:   req.FriendID,
			FriendID: req.UserID,
			Status:   1,
		})
		c.JSON(http.StatusOK, gin.H{"message": "已同意好友申请喵~"})
	} else {
		DB.Delete(&req)
		c.JSON(http.StatusOK, gin.H{"message": "已拒绝好友申请喵"})
	}
}
