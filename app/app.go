package app

import (
	"fmt"
	"github.com/davveo/market-coupon/app/router"
	"github.com/davveo/market-coupon/app/service"
	"github.com/davveo/market-coupon/config"
	"github.com/davveo/market-coupon/cron"
	"github.com/davveo/market-coupon/pkg/constant"
	"github.com/davveo/market-coupon/pkg/db"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

type Server struct {
	c     *config.Config
	db    *db.Datastore
	redis *rdsV8.Client
}

func NewServer(conf *config.Config, db *db.Datastore, redis *rdsV8.Client) *Server {
	return &Server{c: conf, db: db, redis: redis}
}

func (s *Server) Init() {
	if s.c.Meta.Env != constant.EnvDev {
		gin.SetMode(gin.ReleaseMode)
	}

	// 启动定时任务
	go cron.Run(service.NewService(s.c, s.db, s.redis))

	route := gin.New()
	route.Use(gin.Recovery())
	router.Init(route, s.db, s.c, s.redis)
	err := route.Run(fmt.Sprintf("%s:%d", s.c.Host, s.c.Port))
	if err != nil {
		fmt.Println(fmt.Errorf("fail to init server: %w", err))
		panic(err)
	}
}
