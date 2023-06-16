package router

import (
	"github.com/davveo/coupon-server/app/controller/coupon_consume"
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

func CouponConsumeRouterGroup(router *gin.RouterGroup, db *db.Datastore, conf *config.Config, redis *rdsV8.Client) {
	couponConsumeGroup := router.Group("/couponConsume")
	couponConsumeController := coupon_consume.NewController(conf, db, redis)
	{
		couponConsumeGroup.POST("/order/history/query", couponConsumeController.QueryOrderConsumedCoupon)
		couponConsumeGroup.POST("/item/available/query", couponConsumeController.QueryMatchedCouponForPurchaseItem)
		couponConsumeGroup.POST("/order/available/query", couponConsumeController.QueryMatchedCouponForOrder)
		couponConsumeGroup.POST("/available/sort", couponConsumeController.SortAvailableMatchedCoupons)
		couponConsumeGroup.POST("/consume", couponConsumeController.ConsumeCoupon)
	}
}
