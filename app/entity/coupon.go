package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type UseCouponDto struct {
	CouponId       int64 `json:"couponId"`
	CouponQuantity int64 `json:"couponQuantity"`
}

type QueryOrderConsumedCouponReqDto struct {
	OrderId string `json:"orderId"`
}

type QueryItemMatchedCouponReqDto struct {
	MemberDto    `json:"memberDto"`
	OrderItemDto `json:"orderItemDto"`
}

type QueryOrderMatchedCouponDto struct {
	MemberDto `json:"memberDto"`
	OrderDto  `json:"orderDto"`
}

type ConsumeCouponReqDto struct {
	MemberDto `json:"memberDto"`
	OrderDto  `json:"orderDto"`
}

type AvailableCouponDetailsDto struct {
	CouponId        int64           `json:"couponId"`        // 消费优惠券(模板)记录ID
	CouponDetailId  int64           `json:"couponDetailId"`  // 消费优惠券领劵明细记录ID
	CouponNo        int64           `json:"couponNo"`        // 优惠券编码,发券时按规则生成
	CouponAmount    decimal.Decimal `json:"couponAmount"`    // 券额(积分数,金额,折扣率等)
	CouponQrCodeUrl string          `json:"couponQrCodeUrl"` // 优惠券二维码地址URL
	CouponWeight    int             `json:"couponWeight"`    // 权重值,数值小表示优先
	CouponEndDate   time.Time       `json:"couponEndDate"`   // 优惠券使用截至日期
}

type SortAvailableMatchedCouponsReqDto struct {
	AvailableCouponDetailsDtoList []*AvailableCouponDetailsDto `json:"availableCouponDetailsDtoList"`
}
