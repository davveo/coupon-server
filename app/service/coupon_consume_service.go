package service

import (
	"context"
	"github.com/davveo/coupon-server/app/dao/model/bo"
	"github.com/shopspring/decimal"
)

// CouponConsumeService 优惠券使用接口
type CouponConsumeService interface {
	// QueryOrderConsumedCoupon 查询订单关联的优惠券使用记录
	QueryOrderConsumedCoupon(ctx context.Context, orderId string) (*bo.CouponConsumeHistoryRespBo, error)
	// QueryMatchedCouponForPurchaseItem 查询会员购买单项商品/服务可使用的优惠券
	QueryMatchedCouponForPurchaseItem(ctx context.Context,
		memberReqBo *bo.MemberBo, orderItem *bo.OrderItemBo) (*bo.OrderItemAvailableCouponsRespBo, error)
	// QueryMatchedCouponForOrder 查询会员下单可使用的优惠券
	QueryMatchedCouponForOrder(ctx context.Context,
		memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo) (*bo.OrderAvailableCouponsRespBo, error)
	// GetSelectCouponReceiveInfoList 查询订单或商品(各活动)可用优惠券结果集合
	GetSelectCouponReceiveInfoList(ctx context.Context,
		memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo, orderItemReqBo *bo.OrderItemBo) ([]*bo.SelectCouponReceiveInfoBo, error)
	// GetTotalCouponValidRemainQuantity 获取订单或商品(各活动)优惠券总剩余数量
	GetTotalCouponValidRemainQuantity(ctx context.Context, memberReqBo *bo.MemberBo,
		orderReqBo *bo.OrderBo, orderItemReqBo *bo.OrderItemBo, couponId int64) (int64, error)
	// GetSelectCouponUseList 将用户优惠券可选列表排序返回
	GetSelectCouponUseList(ctx context.Context, orderReqBo *bo.OrderBo, orderItemReqBo *bo.OrderItemBo,
		memberReqBo *bo.MemberBo, selectCouponReceiveInfoList []*bo.SelectCouponReceiveInfoBo) ([]*bo.SelectCouponUseBo, error)
	// SortAvailableMatchedCoupons 将可用的优惠券按优先使用策略排序
	SortAvailableMatchedCoupons(ctx context.Context,
		availableCouponList []*bo.AvailableCouponDetailsBo) ([]*bo.AvailableCouponDetailsBo, error)
	// ConsumeCoupon 用券相关操作 校验用券信息,每一种优惠券按使用策略顺序更新领券记录状态,并更新用户中心优惠券信息
	ConsumeCoupon(ctx context.Context,
		memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo) ([]*bo.CouponConsumeBo, error)
	// ValidateCouponAvailable 校验当前用户下单和商品/服务优惠券是否可用
	ValidateCouponAvailable(ctx context.Context, memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo,
		isSettleCheck bool, couponReduceAmountMap map[*bo.CouponReduceInfoBo]map[int64]decimal.Decimal) (bool, error)
}
