package model

type Coupon struct {
}

func (Coupon) TableName() string {
	return "t_coupon"
}
