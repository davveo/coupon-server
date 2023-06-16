package bo

import "github.com/shopspring/decimal"

type UserCouponConsumeProcessBo struct {
	eUserId                string                   `json:"userId"`                 // 用户ID
	UserCode               string                   `json:"userCode"`               // 用户编号
	ShopId                 string                   `json:"shopId"`                 // 商家店铺ID
	OrderId                string                   `json:"orderId"`                // 订单ID
	OrderNo                string                   `json:"orderNo"`                // 订单编号
	OrderItemId            string                   `json:"orderItemId"`            // 订单明细项ID
	SkuId                  string                   `json:"skuId"`                  // 商品/服务ID
	GoodsQuantity          int64                    `json:"goodsQuantity"`          // 订单项/商品购买数量
	OriginalAmount         decimal.Decimal          `json:"originalAmount"`         // 原订单或商品项金额
	FinalPayAmount         decimal.Decimal          `json:"finalPayAmount"`         // 优惠后订单或商品项实际支付金额
	CouponReduceAmountList []*CouponsReduceAmountBo `json:"couponReduceAmountList"` // 各种优惠券扣减金额信息
}
