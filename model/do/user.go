package do

import "time"

type User struct {
	BaseDO
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female 女,male 男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1:普通用户, 2:管理员'"`
}