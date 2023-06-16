package model

import "time"

type CouponReceiveDetail struct {
	Id              int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	CouponReceiveId int64     `gorm:"column:coupon_receive_id;NOT NULL;comment:'优惠券领取ID'"`
	CouponNo        string    `gorm:"column:coupon_no;default:NULL;comment:'优惠券编码,发券时按规则生成'"`
	CouponAmount    string    `gorm:"column:coupon_amount;default:0.00;comment:'券额,包含积分数,金额,折扣率'"`
	CouponQrCodeUrl string    `gorm:"column:coupon_qr_code_url;default:NULL;comment:'优惠券二维码地址URL'"`
	Status          int32     `gorm:"column:status;default:0;comment:'状态，0为已领取未使用，1为已使用，2为使用中，3为已过期'"`
	Version         int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	IsDelete        int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	CreatorId       int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'领取时间'"`
	ModifierId      int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt      time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (CouponReceiveDetail) TableName() string {
	return "t_coupon_receive_detail"
}
