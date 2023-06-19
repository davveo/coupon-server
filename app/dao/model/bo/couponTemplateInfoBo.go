package bo

type CouponTemplateInfoBo struct {
	CouponId     int64   `json:"couponId"`     // 消费优惠券(模板)记录ID
	CouponName   string  `json:"couponName"`   // 消费优惠券(模板)名称
	CouponDesc   string  `json:"couponDesc"`   // 消费优惠券(模板)描述
	CouponIcon   string  `json:"couponIcon"`   // 消费优惠券(模板)图标URL
	CouponAmount float64 `json:"couponAmount"` // 券额(积分数,金额,折扣率等)
	CouponWeight int     `json:"couponWeight"` // 权重值,数值小表示优先
}
