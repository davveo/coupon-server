package service

import "context"

type AccountRuleService interface {
	ValidateAccountRuleMatch(ctx context.Context, couponId, couponQuantity int64, originPayloadValue float64) (bool error)
	CheckAccountRuleMatchResult(ctx context.Context) (bool, error)
}
