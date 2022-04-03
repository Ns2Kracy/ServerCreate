package datasource

import (
	"ServerCreate/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

func NewRedis() *redis.Database {
	var database *redis.Database

	cmsConfig := config.InitConfig()
	if cmsConfig != nil {
		Redis := cmsConfig.Redis
		iris.New().Logger().Info(Redis)
		database = redis.New(redis.Config{
			Network:   Redis.NetWork,
			Addr:      Redis.Addr + ":" + Redis.Port, //139.224.19.236:6379
			Password:  "",
			Database:  "",
			MaxActive: 10,
			Prefix:    Redis.Prefix,
		})
	} else {
		iris.New().Logger().Info("error")
	}
	iris.New().Logger().Info(database)
	defer database.Close()
	return database
}
