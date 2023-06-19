package bo

// SelectCouponUseBo 按模板分组的可使用优惠券列表
type SelectCouponUseBo struct {
	CouponTotalRemainQuantity   int64                        `json:"couponTotalRemainQuantity"`   // 用户此类优惠券总共剩余数量
	SelectCouponReceiveInfoList []*SelectCouponReceiveInfoBo `json:"selectCouponReceiveInfoList"` //用户此类优惠券各活动领劵记录列表
}
