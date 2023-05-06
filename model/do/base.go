package do

import (
	"gorm.io/gorm"
	"time"
)

type BaseDO struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdateAt  time.Time `gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt
}
