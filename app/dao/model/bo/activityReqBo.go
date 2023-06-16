package bo

import "time"

type ActivityReqBo struct {
	ID                int64     `json:"id"`                // 促销活动ID
	CouponID          int64     `json:"couponId"`          // 基础优惠券id
	ActivityName      string    `json:"activityName"`      // 促销活动名称
	ActivityDesc      string    `json:"activityDesc"`      // 促销活动描述
	ActivityStart     time.Time `json:"activityStart"`     // 优惠券活动开始日期
	ActivityEnd       time.Time `json:"activityEnd"`       // 优惠券活动结束日期
	EffectDate        time.Time `json:"effectDate"`        // 优惠券生效日期
	EffectDays        int       `json:"effectDays"`        // 优惠券有效天数,用于发券开始计算有效截至期
	StartTime         time.Time `json:"startTime"`         // 优惠券使用开始时间
	EndTime           time.Time `json:"endTime"`           // 优惠券使用结束时间
	QuantityLimit     bool      `json:"quantityLimit"`     // 数量限制，无限制则不需要发券，自动用券
	IssueQuantity     int64     `json:"issueQuantity"`     // 总发放数量
	ReceivedQuantity  int64     `json:"receivedQuantity"`  // 已发放数量
	LimitQuantity     int64     `json:"limitQuantity"`     // 每人最多领取数量
	AutoIssueQuantity int       `json:"autoIssueQuantity"` // 自动发放数量
	AutoIssueTime     time.Time `json:"autoIssueTime"`     // 自动发放时间
	Status            int       `json:"status"`            // 状态，0表示未开始，1表示进行中，2表示结束
}
