package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	IsActive  bool   `gorm:"default:false" json:"is_active"`
}

type Friend struct {
	gorm.Model
	UserID   uint `gorm:"uniqueIndex:idx_user_friend" json:"user_id"`
	FriendID uint `gorm:"uniqueIndex:idx_user_friend" json:"friend_id"`
	Status   int  `json:"status"` // 0: pending, 1: accepted
}

type Message struct {
	gorm.Model
	FromUserID uint   `json:"from_user_id"`
	ToUserID   uint   `json:"to_user_id"`
	GroupID    uint   `json:"group_id"`
	Content    string `json:"content"`
	Type       int    `json:"type"` // 0: text, 1: image, 2: voice
}

type Group struct {
	gorm.Model
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
	Avatar  string `json:"avatar"`
}

type GroupMember struct {
	gorm.Model
	GroupID uint `json:"group_id"`
	UserID  uint `json:"user_id"`
	Role    int  `json:"role"` // 0: member, 1: owner
}
