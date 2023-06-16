package model

type AccountRule struct {
}

func (AccountRule) TableName() string {
	return "t_coupon_account_rule"
}
