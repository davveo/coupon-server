package bo

import (
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/shopspring/decimal"
)

type CalcTotalReduceByEachCouponReqBo struct {
	CouponQuantity  int64              `json:"couponQuantity"`  // 优惠券使用数量
	AccountRuleInfo *model.AccountRule `json:"accountRuleInfo"` // 结算规则信息
	GoodsSalePrice  decimal.Decimal    `json:"goodsSalePrice"`  // 商品/服务销售原单价
	TmpTotalAmount  decimal.Decimal    `json:"tmpTotalAmount"`  // 临时变量(当前商品或订单的总金额)
}
