package model

import "time"

type PromotionRule struct {
	Id              int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'自动增加ID'"`
	ActivityId      int64     `gorm:"column:activity_id;default:NULL;comment:'促销活动ID'"`
	RuleName        string    `gorm:"column:rule_name;default:NULL;comment:'规则名称'"`
	RuleType        int8      `gorm:"column:rule_type;default:0;comment:'规则类型 0:发放规则;1:使用规则'"`
	UserCondition   string    `gorm:"column:user_condition;default:NULL;comment:'用户适用条件 json字符串'"`
	SkuCondition    string    `gorm:"column:sku_condition;default:NULL;comment:'商品或服务适用条件 json字符串'"`
	OrderCondition  string    `gorm:"column:order_condition;default:NULL;comment:'订单适用条件 json字符串'"`
	TimeCondition   string    `gorm:"column:time_condition;default:NULL;comment:'适用时间范围条件 条件表达式'"`
	ExtrasCondition string    `gorm:"column:extras_condition;default:NULL;comment:'其他使用范围条件 json字符串'"`
	IsDelete        int       `gorm:"column:is_delete;default:b;comment:'是否生效'"`
	Version         int64     `gorm:"column:version;default:NULL;comment:'乐观锁标识'"`
	CreatorId       int64     `gorm:"column:creator_id;default:NULL;comment:'创建人'"`
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	ModifierId      int64     `gorm:"column:modifier_id;default:NULL;comment:'修改人'"`
	ModifiedAt      time.Time `gorm:"column:modified_at;default:CURRENT_TIMESTAMP;comment:'修改时间'"`
}

func (PromotionRule) TableName() string {
	return "promotion_rule"
}
