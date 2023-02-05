package main

import (
	"fmt"
	"gin_blog/config"
	"gin_blog/dao"
	"gin_blog/logger"
	"gin_blog/routers"
)

func main() {
	// load config from  conf/conf.json
	//if len(os.Args) < 2 {
	//	return
	//}
	//if err := config.Init(os.Args[1]); err != nil {
	//	fmt.Printf("config.Init failed, err:%v\n", err)
	//	return
	//}
	s := `{
	"server": {
	  "port": 8080
	},
	"mysql": {
	  "host": "127.0.0.1",
	  "port": 3306,
	  "db": "gin_blog",
	  "username": "root",
	  "password": "961125"
	},
	"redis": {
	  "host": "127.0.0.1",
	  "port": 6379,
	  "db": 0,
	  "password": ""
	},
	"log":{
		"level": "debug",
	  "filename": "log/gin_blog.log",
	  "maxsize": 500,
	  "max_age": 7,
	  "max_backups": 10
	}
}`
	if err := config.InitFromStr(s); err != nil {
		fmt.Printf("config.Init failed, err:%v\n", err)
		return
	}
	// init logger
	//err := config.InitFromIni("conf/conf.ini")
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(config.Conf.ServerConfig.Port, *config.Conf.LogConfig)

	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// init MySQL
	if err := dao.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	// init redis
	if err := dao.InitRedis(config.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	logger.Logger.Info("start project...")
	//err := dao.InitMySQL(config.Conf.MySQLConfig)
	//if err != nil {
	//	fmt.Printf("init MySQL failed, err:%v\n", err)
	//	logger.Logger.Error("init MySQL failed", zap.Any("error", err))
	//	return
	//}
	r := routers.SetupRouter() // 初始化路由
	r.Run()
}
