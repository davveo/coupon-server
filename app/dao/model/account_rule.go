package model

import "time"

type AccountRule struct {
	Id               int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	RuleName         string    `gorm:"column:rule_name;default:NULL;comment:' 结算规则名称'"`
	RuleDesc         string    `gorm:"column:rule_desc;default:NULL;comment:' 规则描述'"`
	ThresholdType    int8      `gorm:"column:threshold_type;default:NULL;comment:'满减条件类型 0:无条件;1:满金额;2:满数量;3:优惠价'"`
	RewardThreshold  float64   `gorm:"column:reward_threshold;default:0.00;comment:'满减条件阈值, 0表示无条件阈值'"`
	RewardType       int8      `gorm:"column:reward_type;default:0;comment:'扣减类型 0:积分;1:抵扣金额;2:折扣;3:其他赠品'"`
	RewardAmount     float64   `gorm:"column:reward_amount;default:0.00;comment:'扣减数,包括扣减积分数,金额,折扣率'"`
	RewardProduct    string    `gorm:"column:reward_product;default:0;comment:'赠送商品或服务ID'"`
	RewardActivityId int64     `gorm:"column:reward_activity_id;default:0;comment:'赠送优惠券活动ID'"`
	RewardDesc       string    `gorm:"column:reward_desc;default:NULL;comment:'赠品描述'"`
	PromotionPrice   float64   `gorm:"column:promotion_price;default:0.00;comment:'优惠价'"`
	ExchangeRatio    float64   `gorm:"column:exchange_ratio;default:1.00;comment:'金额兑换积分比例,积分除以该比例等于金额'"`
	IsDelete         int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	Version          int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	CreatorId        int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt        time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	ModifierId       int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt       time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (AccountRule) TableName() string {
	return "t_coupon_account_rule"
}
