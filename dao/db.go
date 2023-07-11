package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	driverName := viper.GetString("datasource.driverName")
	name := viper.GetString("datasource.name")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	parseTime := viper.GetString("datasource.parseTime")
	loc := viper.GetString("datasource.loc")

	mysqlArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		name,
		password,
		host,
		port,
		database,
		charset,
		parseTime,
		loc)

	switch driverName {
	case "mysql":
		db, err = gorm.Open(mysql.Open(mysqlArgs), &gorm.Config{})
	default:
		db, err = gorm.Open(mysql.Open(mysqlArgs), &gorm.Config{})
	}
	if err != nil {
		panic("failed to connect sql,err:" + err.Error())
	}
	//对于所有结构体都需要在这注册
	err = db.AutoMigrate(&Admin{}, &Level{})
	//回调函数，每次更新数据时执行
	//_ = db.Callback().Update().Before("gorm:update").Register("my_plugin:before_update", beforeUpdate)
}

// 回调函数
func beforeUpdate(tx *gorm.DB) {
	if tx.Statement.Changed("DeletedAt") {
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
