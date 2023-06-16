package model

import "time"

type CouponReceive struct {
	Id              int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	UserId          string    `gorm:"column:user_id;default:NULL;comment:'用户买家ID'"`
	UserCode        string    `gorm:"column:user_code;default:NULL;comment:'用户买家编号'"`
	CouponId        int64     `gorm:"column:coupon_id;default:NULL;comment:'优惠券编号ID'"`
	ActivityId      int64     `gorm:"column:activity_id;default:NULL;comment:'促销活动ID'"`
	CouponAmount    float64   `gorm:"column:coupon_amount;default:0.00;comment:'券额,包含积分数,金额,折扣率'"`
	ReceiveQuantity int64     `gorm:"column:receive_quantity;default:0;comment:'领券数量'"`
	RemainQuantity  int64     `gorm:"column:remain_quantity;default:0;comment:'使用后剩余数量'"`
	StartDate       string    `gorm:"column:start_date;default:NULL;comment:'优惠券生效日期'"`
	EndDate         string    `gorm:"column:end_date;default:NULL;comment:'优惠券截至日期'"`
	Remark          string    `gorm:"column:remark;default:NULL;comment:'优惠券领取场景描述'"`
	OrderId         string    `gorm:"column:order_id;default:NULL;comment:'领券关联订单ID'"`
	OrderNo         string    `gorm:"column:order_no;default:NULL;comment:'订单编号'"`
	OrderItemId     string    `gorm:"column:order_item_id;default:NULL;comment:'订单明细项ID'"`
	Status          int32     `gorm:"column:status;default:0;comment:'状态，0为已领取未使用，1为已使用，2为部分使用, 3为已过期'"`
	Version         int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	IsDelete        int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	CreatorId       int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'领取时间'"`
	ModifierId      int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt      time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (CouponReceive) TableName() string {
	return "t_coupon_receive"
}
