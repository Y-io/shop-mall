package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
	"shop-mall/config"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
	Trans        ut.Translator
)
