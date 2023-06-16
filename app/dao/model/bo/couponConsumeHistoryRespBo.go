package bo

import "github.com/shopspring/decimal"

type CouponConsumeHistoryRespBo struct {
	// CompanyId        int64
	OrderId                 string                 `json:"orderId"`                 // 订单ID
	OrderNo                 string                 `json:"orderNo"`                 // 订单编号
	ShopId                  string                 `json:"shopId"`                  // 卖家商户供应商ID
	OriginalAmount          decimal.Decimal        `json:"originalAmount"`          // 原订单或商品项金额
	FinalPayAmount          decimal.Decimal        `json:"finalPayAmount"`          // 优惠后订单或商品项实际支付金额
	CouponHistoryInfoBoList []*CouponHistoryInfoBo `json:"couponHistoryInfoBoList"` // 使用优惠券模板相关明细列表
}
