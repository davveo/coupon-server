package model

import "time"

type Coupon struct {
	Id              int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	AccountRuleId   int64     `gorm:"column:account_rule_id;default:NULL;comment:'结算规则ID'"`
	CompanyId       int32     `gorm:"column:company_id;default:NULL;comment:'公司,租户ID'"`
	CouponType      int8      `gorm:"column:coupon_type;default:NULL;comment:'所属类型(0:积分;1:满减(包括现金抵扣和折扣,参考满减规则),2:赠送赠品)'"`
	IsAccumulated   int       `gorm:"column:is_accumulated;default:b;comment:'使用该优惠券产生的扣减是否计入满减条件的支付总量'"`
	CouponName      string    `gorm:"column:coupon_name;default:NULL;comment:'优惠券名称'"`
	CouponDesc      string    `gorm:"column:coupon_desc;default:NULL;comment:'优惠券活动描述'"`
	Icon            string    `gorm:"column:icon;default:NULL;comment:'优惠券图标的URL地址'"`
	ShopId          string    `gorm:"column:shop_id;default:NULL;comment:'适用商户供应商ID'"`
	SkuCategory     int32     `gorm:"column:sku_category;default:NULL;comment:'对象范围,适用商品或服务分类ID'"`
	UserCategory    int32     `gorm:"column:user_category;default:NULL;comment:'顾客资格,适用用户分类'"`
	SkuId           int32     `gorm:"column:sku_id;default:NULL;comment:'商品或服务SKU id'"`
	AccumulateType  int8      `gorm:"column:accumulate_type;default:0;comment:'优惠券使用规则之间关系,0:排他;1:并存(累计);2:择优'"`
	Weight          int8      `gorm:"column:weight;default:0;comment:'权重值,数值小表示优先'"`
	IsSupportReturn int       `gorm:"column:is_supportReturn;default:b;comment:'退款时是否支持退回'"`
	Status          int8      `gorm:"column:status;default:0;comment:'状态(0-已启用,1-未启用)'"`
	IsDelete        int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	Version         int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	CreatorId       int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	ModifierId      int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt      time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (Coupon) TableName() string {
	return "t_coupon"
}
