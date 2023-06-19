package bo

type GetAvailableCouponSelectListBo struct {
	SelectCouponUseList []*SelectCouponUseBo `json:"selectCouponUseList"` //可使用的优惠券列表,按模板分组,可查看详情用户在每次活动领劵情况
}
