package biz

import (
	"context"
	"github.com/davveo/coupon-server/app/entity"
)

type CouponConsumeBiz interface {
	// QueryOrderConsumedCoupon 查询订单关联的优惠券使用记录
	QueryOrderConsumedCoupon(ctx context.Context, req *entity.QueryOrderConsumedCouponReqDto)
	// QueryMatchedCouponForPurchaseItem 查询会员购买单项商品/服务可使用的优惠券
	QueryMatchedCouponForPurchaseItem(ctx context.Context, req *entity.QueryItemMatchedCouponReqDto)
	// QueryMatchedCouponForOrder 查询会员/用户下单可使用的优惠券
	QueryMatchedCouponForOrder(ctx context.Context, req *entity.QueryOrderMatchedCouponDto)
	// SortAvailableMatchedCoupons 将可用优惠券明细列表按优先使用策略排序
	SortAvailableMatchedCoupons(ctx context.Context, req *entity.SortAvailableMatchedCouponsReqDto)
	// ConsumeCoupon 校验用券信息,每一种优惠券按使用策略顺序更新领券记录状态,并更新用户中心优惠券信息
	ConsumeCoupon(ctx context.Context, req *entity.ConsumeCouponReqDto)
}
