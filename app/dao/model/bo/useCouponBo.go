package bo

type UseCouponBo struct {
	CouponId       int64 `json:"couponId"`       // 消费优惠券(模板)ID
	CouponQuantity int64 `json:"couponQuantity"` // 本次消费优惠券数量
}
