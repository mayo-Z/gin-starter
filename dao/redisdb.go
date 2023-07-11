package dao

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
)

func GetSessionStore() (sessions.Store, error) {
	// 1. 配置Redis
	network := viper.GetString("redis.network")
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	size := viper.GetInt("redis.size")
	key := viper.GetString("redis.key")

	return redis.NewStore(size, network, address, password, []byte(key))

}
