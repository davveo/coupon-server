package bo

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderBo struct {
	OrderId           string          `json:"orderId"`           // 订单ID
	OrderNo           string          `json:"orderNo"`           // 订单编号
	TotalOrderAmount  decimal.Decimal `json:"totalOrderAmount"`  // 订单参与计算总金额(动态选劵后参与满减条件计算的虚拟总金额)
	OriginOrderAmount decimal.Decimal `json:"originOrderAmount"` // 订单原始(优惠前)总金额
	PayOrderAmount    decimal.Decimal `json:"payOrderAmount"`    // 结算后实际应付金额
	OrderDate         time.Time       `json:"orderDate"`         // 下单时间
	ReserveTime       time.Time       `json:"reserveTime"`       // 预约时间
	OrderItemList     []*OrderItemBo  `json:"orderItemList"`     // 订单明细项
	UseCoupons        []*UseCouponBo  `json:"useCoupons"`        // 订单所使用的优惠券
}
