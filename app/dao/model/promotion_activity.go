package model

import "time"

type PromotionActivity struct {
	Id                int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	CouponId          int64     `gorm:"column:coupon_id;default:NULL;comment:'基础优惠券ID'"`
	ActivityName      string    `gorm:"column:activity_name;default:NULL;comment:'促销活动名称'"`
	ActivityDesc      string    `gorm:"column:activity_desc;default:NULL;comment:'促销活动描述'"`
	ActivityStart     time.Time `gorm:"column:activity_start;default:NULL;comment:'优惠券活动开始日期'"`
	ActivityEnd       time.Time `gorm:"column:activity_end;default:NULL;comment:'优惠券活动结束日期'"`
	EffectDate        time.Time `gorm:"column:effect_date;default:NULL;comment:'优惠券生效日期'"`
	EffectDays        int32     `gorm:"column:effect_days;default:NULL;comment:'优惠券有效天数,用于发券开始计算有效截至期'"`
	StartTime         time.Time `gorm:"column:start_time;default:NULL;comment:'优惠券使用开始时间'"`
	EndTime           time.Time `gorm:"column:end_time;default:NULL;comment:'优惠券使用结束时间'"`
	IsQuantityLimit   int       `gorm:"column:is_quantityLimit;default:b;comment:'数量限制，无限制则不需要发券，自动用券'"`
	IssueQuantity     int64     `gorm:"column:issue_quantity;default:0;comment:'总发放数量'"`
	ReceivedQuantity  int64     `gorm:"column:received_quantity;default:0;comment:'已发放数量'"`
	LimitQuantity     int64     `gorm:"column:limit_quantity;default:NULL;comment:'每人最多领取数量'"`
	AutoIssueQuantity int32     `gorm:"column:auto_issue_quantity;default:0;comment:'自动发放数量'"`
	AutoIssueTime     time.Time `gorm:"column:auto_issue_time;default:NULL;comment:'自动发放时间'"`
	Status            int8      `gorm:"column:status;default:NULL;comment:'状态，0表示未开始，1表示进行中，2表示结束'"`
	IsDelete          int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	Version           int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	CreatorId         int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt         time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	ModifierId        int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt        time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (PromotionActivity) TableName() string {
	return "t_promotion_activity"
}
