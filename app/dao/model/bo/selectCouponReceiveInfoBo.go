package bo

import "time"

type SelectCouponReceiveInfoBo struct {
	CouponTemplateInfoBo
	CouponReceiveId             int64     `json:"couponReceiveId"`             // 优惠券领劵主记录ID
	CouponReceiveRemainQuantity int64     `json:"couponReceiveRemainQuantity"` // 优惠券领劵剩余数量
	CouponReceiveStatus         int       `json:"couponReceiveStatus"`         // 优惠券领劵记录状态(0-已领取未使用,1-已使用,2-使用中,3-已过期)
	CouponEndDate               time.Time `json:"couponEndDate"`               // 优惠券使用截至日期
}
