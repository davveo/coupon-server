package router

import (
	"github.com/davveo/market-coupon/app/controller/coupon"
	"github.com/davveo/market-coupon/config"
	"github.com/davveo/market-coupon/pkg/db"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

func CouponRouterGroup(router *gin.RouterGroup, db *db.Datastore, conf *config.Config, redis *rdsV8.Client) {
	couponGroup := router.Group("/v1/coupon")
	couponController := coupon.NewController(conf, db, redis)
	{
		couponGroup.POST("/list", couponController.Page)
	}
}
