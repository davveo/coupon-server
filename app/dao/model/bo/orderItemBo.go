package bo

import "github.com/shopspring/decimal"

type OrderItemBo struct {
	OrderItemId      string          `json:"orderItemId"`      // 订单明细项ID
	SkuId            string          `json:"skuId"`            // 商品/服务SKU ID
	ShopId           string          `json:"shopId"`           // 商家ID
	ShopName         string          `json:"shopName"`         // 商铺名称
	SkuCategory      string          `json:"skuCategory"`      // 商品/服务分类
	GoodsSalePrice   decimal.Decimal `json:"goodsSalePrice"`   // 商品/服务销售单价
	TotalGoodsAmount decimal.Decimal `json:"totalGoodsAmount"` // 商品/服务原始金额
	GoodsQuantity    int64           `json:"goodsQuantity"`    // 商品/服务数量
	PayGoodsAmount   decimal.Decimal `json:"payGoodsAmount"`   // 结算后实际应付金额
	UseCoupons       []UseCouponBo   `json:"useCoupons"`       // 所使用的优惠券
}
