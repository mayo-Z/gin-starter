package main

import (
	"gin-starter/dao"
	"gin-starter/route"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	route.InitConfig()
	dao.InitDB()
	r := route.InitRouter()
	port := viper.GetString("server.port")
	r.Run(":" + port)
}
