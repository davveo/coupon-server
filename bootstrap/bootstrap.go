package bootstrap

import (
	"fmt"
	"github.com/davveo/market-coupon/app"
	"github.com/davveo/market-coupon/config"
	"github.com/davveo/market-coupon/pkg/db"
	"github.com/davveo/market-coupon/pkg/logger"
	"github.com/davveo/market-coupon/pkg/redis"
	"github.com/davveo/market-coupon/pkg/signal"
)

func Bootstrap() {
	conf, err := config.Init()
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger(conf.Log)
	defer logger.Close()
	if err != nil {
		panic(err)
	}

	dbIns, err := db.Init(&conf.MySql)
	defer dbIns.Close()
	if err != nil {
		panic(err)
	}

	redisIns, err := redis.Init(&conf.Redis)
	defer redisIns.Close()
	if err != nil {
		panic(err)
	}
	server := app.NewServer(conf, dbIns, redisIns)
	server.Init()
	fmt.Println("init api server success")

	signal.Wait()
}
