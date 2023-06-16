package router

import (
	"github.com/davveo/coupon-server/app/controller/coupon"
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
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
