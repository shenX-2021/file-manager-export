package model

import (
	"file-manager-export/args"
	"file-manager-export/exit"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(args args.Args) *gorm.DB {
	if DB != nil {
		return DB
	}

	// 判断sqlite文件是否存在
	_, err := os.Open(args.DbPath)
	if err != nil {
		if os.IsNotExist(err) {
			exit.Error("sqlite文件不存在")
		}
		exit.Error(err.Error())
	}

	// gorm初始化
	DB, err = gorm.Open(sqlite.Open(args.DbPath))
	if err != nil {
		exit.Error("sqlite数据库无法打开，请检查" + args.DbPath + "文件是否存在")
	}

	return DB
}
