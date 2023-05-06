package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	ID        int32          `json:"id" description:"UID"`
	CreatedAt time.Time      `json:"createdAt" description:"创建时间"`
	UpdateAt  time.Time      `json:"updatedAt" description:"更新时间"`
	DeleteAt  gorm.DeletedAt `json:"deletedAt" description:"删除时间"`
}
