package bo

type CouponConsumeDetailBo struct {
	CouponConsumeId        int64  `json:"couponConsumeId"`        // 消费优惠券ID
	CouponReceiveId        int64  `json:"couponReceiveId"`        // 发放优惠券ID
	CouponReceiveDetailIds string `json:"couponReceiveDetailIds"` // 发放优惠券明细ID集合(逗号间隔
	ActivityId             int64  `json:"activityId"`             // 促销活动ID
	ConsumeQuantity        int64  `json:"consumeQuantity"`        // 本次使用数量
}
