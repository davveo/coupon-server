package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderItemDto struct {
	OrderItemId      string          `json:"orderItemId"`      // 订单明细项ID
	SkuId            string          `json:"skuId"`            // 商品/服务SKU ID
	ShopId           string          `json:"shopId"`           // 商家ID
	ShopName         string          `json:"shopName"`         // 商铺名称
	SkuCategory      string          `json:"skuCategory"`      // 商品/服务分类
	GoodsSalePrice   decimal.Decimal `json:"goodsSalePrice"`   // 商品/服务销售单价
	GoodsQuantity    int64           `json:"goodsQuantity"`    // 商品/服务数量
	TotalGoodsAmount decimal.Decimal `json:"totalGoodsAmount"` // 商品/服务原始金额
	PayGoodsAmount   decimal.Decimal `json:"payGoodsAmount"`   // 结算后实际应付金额
	UseCoupons       []*UseCouponDto `json:"useCoupons"`       // 所使用的优惠券
}

type OrderDto struct {
	OrderId           string          `json:"orderId"`           // 订单ID
	OrderNo           string          `json:"orderNo"`           // 订单编号
	TotalOrderAmount  decimal.Decimal `json:"totalOrderAmount"`  // 订单总金额(动态选劵后参与满减条件计算的虚拟总金额)
	OriginOrderAmount decimal.Decimal `json:"originOrderAmount"` // 订单原始(优惠前)总金额
	PayOrderAmount    decimal.Decimal `json:"payOrderAmount"`    // 结算后实际应付金额
	OrderDate         time.Time       `json:"orderDate"`         // 下单时间
	ReserveTime       time.Time       `json:"reserveTime"`       // 预约时间
	OrderItemList     []OrderItemDto  `json:"orderItemList"`     // 订单明细项
	UseCoupons        []UseCouponDto  `json:"useCoupons"`        // 订单所使用的优惠券
}
