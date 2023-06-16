package router

import (
	"github.com/davveo/coupon-server/app/controller/coupon_receive"
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

func CouponReceiveRouterGroup(router *gin.RouterGroup, db *db.Datastore, conf *config.Config, redis *rdsV8.Client) {
	couponReceiveGroup := router.Group("/couponReceive")
	couponReceiveController := coupon_receive.NewController(conf, db, redis)
	{
		//单次领券接口
		couponReceiveGroup.POST("/produceCoupon", couponReceiveController.ProduceCoupon)
	}
}
