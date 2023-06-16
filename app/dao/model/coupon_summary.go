package model

import "time"

type CouponSummary struct {
	Id               int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	UserId           string    `gorm:"column:user_id;default:NULL;comment:'用户买家ID'"`
	UserCode         string    `gorm:"column:user_code;default:NULL;comment:'用户买家编号'"`
	CouponId         int64     `gorm:"column:coupon_id;default:NULL;comment:'优惠券编号ID'"`
	CouponAmount     float64   `gorm:"column:coupon_amount;default:0.00;comment:'券额,金额,折扣率'"`
	ReceivedQuantity int64     `gorm:"column:received_quantity;default:0;comment:'历史领取总数'"`
	ConsumedQuantity int64     `gorm:"column:consumed_quantity;default:0;comment:'历史使用总数'"`
	RewardAmount     float64   `gorm:"column:reward_amount;default:0.00;comment:'历史回馈扣减总数'"`
	CouponQuantity   int64     `gorm:"column:coupon_quantity;default:0;comment:'可用优惠券张数或积分余额'"`
	IsDelete         int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	Version          int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	CreatorId        int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt        time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	ModifierId       int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt       time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (CouponSummary) TableName() string {
	return "t_coupon_summary"
}
