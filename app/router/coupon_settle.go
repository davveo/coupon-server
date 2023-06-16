package router

import (
	"github.com/davveo/coupon-server/app/controller/coupon_settle"
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

func CouponSettleRouterGroup(router *gin.RouterGroup, db *db.Datastore, conf *config.Config, redis *rdsV8.Client) {
	couponSettleGroup := router.Group("/couponSettle")
	couponSettleController := coupon_settle.NewController(conf, db, redis)
	{
		// 订单结算
		couponSettleGroup.POST("/order/settlement", couponSettleController.OrderSettlement)
		// 动态返回用户选中的优惠券数据及其金额
		couponSettleGroup.POST("/dynamicSelect", couponSettleController.SelectCoupon)
	}
}
