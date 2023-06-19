package bo

import (
	"github.com/shopspring/decimal"
)

type CouponReduceInfoBo struct {
	UserId                 string          `json:"userId"`
	CouponReduceTargetFlag int             `json:"couponReduceTargetFlag"` // 优惠券扣减目标类型(1-订单,2-商品)
	OrderId                string          `json:"orderId"`
	OrderNo                string          `json:"orderNo"`
	OrderItemId            string          `json:"orderItemId"`
	SkuId                  string          `json:"skuId"`
	SkuCategory            string          `json:"skuCategory"`
	GoodsSalePrice         decimal.Decimal `json:"goodsSalePrice"`
	GoodsQuantity          int64           `json:"goodsQuantity"`
	ShopId                 string          `json:"shopId"`
}
