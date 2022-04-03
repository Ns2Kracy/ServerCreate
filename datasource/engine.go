package datasource

import (
	"ServerCreate/config"
	"ServerCreate/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"log"

	"xorm.io/xorm"
)

func NewEngine() *xorm.Engine {
	initConfig := config.InitConfig()
	if initConfig == nil {
		return nil
	}

	database := initConfig.MySQL

	//root:NK20021001wdnmd.@tcp(139.224.19.236)/servercreate?charset = UTF-8
	dataSourceName := database.User + ":" + database.Pwd + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.Database + "?charset=utf8"

	//engine, err := xorm.NewEngine(database.Drive, "root:NK20021001wdnmd.@tcp(139.224.19.236)/servercreate?charset = UTF-8")
	engine, err := xorm.NewEngine(database.Drive, dataSourceName)

	iris.New().Logger().Info(dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	err = engine.Sync2(new(model.Users))

	err = engine.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	engine.ShowSQL(true)

	return engine

}
