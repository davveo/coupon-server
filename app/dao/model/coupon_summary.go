package model

type CouponSummary struct {
}

func (CouponSummary) TableName() string {
	return "t_coupon_summary"
}
