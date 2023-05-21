package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int32          `json:"id" description:"UID"`
	Password  string         `json:"password"  description:"MD5密码"`
	CreatedAt time.Time      `json:"createdAt" description:"创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" description:"更新时间"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" description:"删除时间"`
	Mobile    string         `json:"mobile"`
	NickName  string         `json:"nickName"`
	Birthday  *time.Time     `json:"birthday"`
	Gender    string         `json:"gender"`
	Role      int            `json:"role"`
}
