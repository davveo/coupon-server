package model

import "time"

type CouponConsume struct {
	Id              uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	UserId          string    `gorm:"column:user_id;default:NULL;comment:'用户买家ID'"`
	UserCode        string    `gorm:"column:user_code;default:NULL;comment:'用户买家编号'"`
	ShopId          string    `gorm:"column:shop_id;default:NULL;comment:'卖家商户供应商ID'"`
	CouponId        int64     `gorm:"column:coupon_id;default:NULL;comment:'基础优惠券ID,使用积分时设置'"`
	OrderId         string    `gorm:"column:order_id;default:NULL;comment:'订单ID'"`
	OrderNo         string    `gorm:"column:order_no;default:NULL;comment:'订单编号'"`
	OrderItemId     string    `gorm:"column:order_item_id;default:NULL;comment:'订单明细项ID'"`
	OriginalAmount  float64   `gorm:"column:original_amount;default:0.00;comment:'原订单或商品项金额'"`
	CouponAmount    float64   `gorm:"column:coupon_amount;default:0.00;comment:'券额,包含积分数,金额,折扣率'"`
	ConsumeQuantity int64     `gorm:"column:consume_quantity;default:1;comment:'用券数量'"`
	FinalPayAmount  float64   `gorm:"column:final_pay_amount;default:0.00;comment:'优惠后订单或商品项实际支付金额'"`
	Status          int32     `gorm:"column:status;default:0;comment:'消费状态: 默认为0已支付使用劵完毕，退回优惠券为1'"`
	Version         int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	IsDelete        int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	CreatorId       int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	ModifierId      int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt      time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (CouponConsume) TableName() string {
	return "t_coupon_consume"
}
