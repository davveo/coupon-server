package bo

import "github.com/shopspring/decimal"

type AccountRuleReqBo struct {
	PagerBaseReqBo
	ID               int64           `json:"ID"`               // 结算规则ID
	RuleName         string          `json:"ruleName"`         // 结算规则名称
	RuleDesc         string          `json:"ruleDesc"`         // 规则描述
	ThresholdType    int             `json:"thresholdType"`    // 满减条件类型 0:无条件;1:满金额;2:满数量;3:优惠价
	RewardThreshold  float64         `json:"rewardThreshold"`  // 满减条件阈值, 0表示无条件阈值
	RewardType       int             `json:"rewardType"`       // 扣减类型 0:积分;1:抵扣金额;2:折扣;3:其他赠品
	RewardAmount     decimal.Decimal `json:"rewardAmount"`     // 扣减数,包括扣减积分数,金额,折扣率
	RewardProduct    string          `json:"rewardProduct"`    // 赠送商品或服务ID
	RewardActivityID int64           `json:"rewardActivityID"` // 赠送优惠券ID
	RewardDesc       string          `json:"rewardDesc"`       // 赠品描述
	PromotionPrice   decimal.Decimal `json:"promotionPrice"`   // 优惠价
	ExchangeRatio    decimal.Decimal `json:"exchangeRatio"`    // 金额兑换积分比例,积分除以该比例等于金额
}
