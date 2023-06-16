package model

import "time"

type CouponConsumeDetail struct {
	Id                     uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	CouponConsumeId        int64     `gorm:"column:coupon_consume_id;default:NULL;comment:'消费优惠券ID'"`
	CouponReceiveId        int64     `gorm:"column:coupon_receive_id;default:NULL;comment:'发放优惠券ID'"`
	CouponReceiveDetailIds string    `gorm:"column:coupon_receive_detail_ids;default:NULL;comment:'发放优惠券明细ID集合'"`
	ActivityId             int64     `gorm:"column:activity_id;default:NULL;comment:'促销活动ID'"`
	ConsumeQuantity        int64     `gorm:"column:consume_quantity;default:1;comment:'本次用券数量'"`
	Version                int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	IsDelete               int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	CreatorId              int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt              time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	ModifierId             int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt             time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (CouponConsumeDetail) TableName() string {
	return "t_coupon_consume_detail"
}
