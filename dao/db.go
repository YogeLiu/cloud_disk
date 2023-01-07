package dao

import (
	"fmt"
	"time"

	"github.com/YogeLiu/CloudDisk/model"
	"github.com/YogeLiu/CloudDisk/pkg/conf"
	"github.com/YogeLiu/CloudDisk/pkg/secret"
	"github.com/YogeLiu/CloudDisk/pkg/util"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	util.Log().Info("initialize database")
	var (
		db  *gorm.DB
		err error
	)
	if gin.Mode() == gin.TestMode {
		// 测试模式下，使用内存数据库
		db, err = gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	} else {
		switch conf.DatabaseConfig.Type {
		case "UNSET", "sqlite", "sqlite3":
			db, err = gorm.Open(sqlite.Open(conf.DatabaseConfig.DBFile), &gorm.Config{})
		case "postgres":
			db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Name,
				conf.DatabaseConfig.Port)), &gorm.Config{})
		case "mysql", "mssql":
			var host string
			if conf.DatabaseConfig.UnixSocket {
				host = fmt.Sprintf("unix(%s)", conf.DatabaseConfig.Host)
			} else {
				host = fmt.Sprintf("tcp(%s:%d)",
					conf.DatabaseConfig.Host,
					conf.DatabaseConfig.Port)
			}
			db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True&loc=Local",
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				host,
				conf.DatabaseConfig.Name,
				conf.DatabaseConfig.Charset)), &gorm.Config{})
		default:
			util.Log().Panic("Unsupported database type %q.", conf.DatabaseConfig.Type)
		}
	}
	if err != nil {
		util.Log().Panic(err.Error())
	}
	d, err := db.DB()
	if err != nil {
		util.Log().Panic(err.Error())
	}
	err = d.Ping()
	if err != nil {
		util.Log().Panic(err.Error())
	}
	d.SetMaxIdleConns(50)
	d.SetConnMaxLifetime(time.Second * 30)
	if conf.DatabaseConfig.Type == "UNSET" || conf.DatabaseConfig.Type == "sqlite" || conf.DatabaseConfig.Type == "sqlite3" {
		d.SetMaxOpenConns(1)
	} else {
		d.SetConnMaxLifetime(30)
	}
	DB = db
	migration()
}

func migration() {
	var err error
	tables := map[string]interface{}{
		"user":   &User{},
		"group":  &Group{},
		"policy": &Policy{},
		"file":   &model.File{},
	}
	for key, table := range tables {
		if !DB.Migrator().HasTable(table) {
			err = DB.Migrator().CreateTable(table)
			if err != nil {
				util.Log().Panic(err.Error())
			}
			switch key {
			case "user":
				password, err := secret.SetPassword("admin")
				if err != nil {
					util.Log().Panic(err.Error())
				}
				err = DB.Model(table).Create(&User{ID: 1, Name: "admin", Password: password}).Error
				if err != nil {
					util.Log().Panic(err.Error())
				}
			case "group":
				err = DB.Model(table).Create(&Group{ID: 1, Name: "管理员"}).Error
				if err != nil {
					util.Log().Panic(err.Error())
				}
			case "policy":
				err = DB.Model(table).Create(&Policy{ID: 1, Name: "本地存储"}).Error
				if err != nil {
					util.Log().Panic(err.Error())
				}
			}
		}
	}
}
