package model

type CouponConsume struct {
}

func (CouponConsume) TableName() string {
	return "t_coupon_consume"
}
