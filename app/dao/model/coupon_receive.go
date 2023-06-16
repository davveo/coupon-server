package model

type CouponReceive struct {
}

func (CouponReceive) TableName() string {
	return "t_coupon_receive"
}
