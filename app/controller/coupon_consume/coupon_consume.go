package coupon_consume

import (
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/davveo/coupon-server/pkg/gin/wrapper"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

type Controller struct {
}

func NewController(conf *config.Config, db *db.Datastore, redis *rdsV8.Client) *Controller {
	return &Controller{}
}

// QueryOrderConsumedCoupon 查询订单关联的优惠券使用记录
func (c *Controller) QueryOrderConsumedCoupon(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}

// QueryMatchedCouponForPurchaseItem 查询会员购买单项商品/服务可使用的优惠券
func (c *Controller) QueryMatchedCouponForPurchaseItem(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}

// QueryMatchedCouponForOrder 查询会员/用户下单可使用的优惠券
func (c *Controller) QueryMatchedCouponForOrder(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}

// SortAvailableMatchedCoupons 将可用优惠券明细列表按优先使用策略排序
func (c *Controller) SortAvailableMatchedCoupons(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}

// ConsumeCoupon 校验用券信息,每一种优惠券按使用策略顺序更新领券记录状态,并更新用户中心优惠券信息
func (c *Controller) ConsumeCoupon(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}
