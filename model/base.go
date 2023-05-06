package model

import (
	_ "gorm.io/gorm"
)

type Pagination struct {
	Page int32 `json:"page"` // 分页码
	Size int32 `json:"size"` // 分页数量
}

type PaginationTotal struct {
	Total int32 `json:"total"` // 数据总数
}
