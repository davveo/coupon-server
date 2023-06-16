package bo

import "github.com/shopspring/decimal"

type CalcCouponReduceResultBo struct {
	CalcCouponReduceSucceed bool            `json:"calcCouponReduceSucceed"` // 校验每种优惠券是否匹配结算递减的满减条件是否成功
	CalcResultTotalAmount   decimal.Decimal `json:"calcResultTotalAmount"`   // 参与下一个劵种计算匹配递减规则后的总金额
	FinalTotalResultAmount  decimal.Decimal `json:"finalTotalResultAmount"`  // 最终计算每种优惠券匹配递减后的实际总金额
}
