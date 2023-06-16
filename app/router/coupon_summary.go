package router

import (
	"github.com/davveo/coupon-server/app/controller/coupon_summary"
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

func CouponSummaryRouterGroup(router *gin.RouterGroup, db *db.Datastore, conf *config.Config, redis *rdsV8.Client) {
	couponSummaryGroup := router.Group("/couponSummary")
	couponSummaryController := coupon_summary.NewController(conf, db, redis)
	{
		//获取用户优惠券列表
		couponSummaryGroup.POST("/getUserCoupon", couponSummaryController.ProduceCoupon)
		// 获取用户优惠券领取详情
		couponSummaryGroup.POST("/getUserCouponDetail", couponSummaryController.ProduceCouponDetail)
	}
}
