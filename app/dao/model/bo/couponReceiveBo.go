package bo

import "github.com/davveo/coupon-server/app/dao/model"

type CouponReceiveBo struct {
	CompanyID            int                          `json:"companyId"`
	CouponReceive        *model.CouponReceive         `json:"couponReceive"`
	CouponReceiveDetails []*model.CouponReceiveDetail `json:"couponReceiveDetails"`
}
