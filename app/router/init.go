package router

import (
	"github.com/davveo/market-coupon/config"
	"github.com/davveo/market-coupon/pkg/db"
	ginLogger "github.com/davveo/market-coupon/pkg/gin/logger"
	"github.com/davveo/market-coupon/pkg/gin/panic_writer"
	"github.com/davveo/market-coupon/pkg/logger"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
	"net/http"
)

func Init(router *gin.Engine, db *db.Datastore, conf *config.Config, redis *rdsV8.Client) {
	initTraceLogger(router)

	publicGroup := router.Group("api")
	{
		publicGroup.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ok": true})
		})
		CouponRouterGroup(publicGroup, db, conf, redis)
	}
}

func initTraceLogger(g *gin.Engine) {
	g.ContextWithFallback = true
	g.Use(gin.RecoveryWithWriter(&panic_writer.PanicWriter{
		Logger: logger.Errorf,
	}))
	g.Use(ginLogger.LogWithWriter())
}
