package bo

import (
	"github.com/shopspring/decimal"
)

type CouponConsumeBo struct {
	OrderID                 string                   `json:"orderID"`                 // OrderID 订单ID
	OrderNo                 string                   `json:"orderNo"`                 // OrderNo 订单编号
	OriginalAmount          decimal.Decimal          `json:"originalAmount"`          // OriginalAmount 原订单金额
	FinalPayAmount          decimal.Decimal          `json:"finalPayAmount"`          // FinalPayAmount 优惠后订单实际支付金额
	CouponConsumeDetailList []*CouponConsumeDetailBo `json:"couponConsumeDetailList"` // CouponConsumeDetailList 使用优惠券明细记录列表
}
