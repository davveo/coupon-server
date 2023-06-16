package bo

import "github.com/shopspring/decimal"

type CouponsReduceAmountBo struct {
	CouponId               int64           `json:"couponId"`               // 优惠ID
	CouponReduceAmount     decimal.Decimal `json:"couponReduceAmount"`     // 扣减优惠券金额
	PaymentAmount          decimal.Decimal `json:"paymentAmount"`          // 扣减优惠券后支付金额
	RewardCouponActivityId int64           `json:"rewardCouponActivityId"` // 赠送优惠券活动Id
	RewardProductId        string          `json:"rewardProductId"`        // 赠送商品或服务ID
	RewardDesc             string          `json:"rewardDesc"`             // 赠品描述
}
