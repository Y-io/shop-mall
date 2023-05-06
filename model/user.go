package model

import (
	"time"
)

type UserGetListItem struct {
	ID        int32
	CreatedAt time.Time
	UpdateAt  time.Time
	Mobile    string
	NickName  string
	Birthday  uint64
	Gender    string
	Role      int
}

type UserGetListInput struct {
	Page     int32 `json:"page"`     // 分页码
	PageSize int32 `json:"pageSize"` // 分页数量
}

type UserGetListOutput struct {
	Page     int32             `json:"page"`     // 分页码
	PageSize int32             `json:"pageSize"` // 分页数量
	Total    int32             `json:"total"`    // 数据总数
	List     []UserGetListItem `json:"list"`
}

type PassWordLoginInput struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
