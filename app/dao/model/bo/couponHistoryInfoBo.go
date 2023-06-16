package bo

import (
	"github.com/shopspring/decimal"
	"time"
)

type CouponHistoryInfoBo struct {
	ActivityId              int64           `json:"activity_id"`
	ActivityName            string          `json:"activity_name"`
	ActivityDesc            string          `json:"activity_desc"`
	ActivityStart           time.Time       `json:"activity_start"`
	ActivityEnd             time.Time       `json:"activity_end"`
	StartTime               time.Time       `json:"start_time"`
	EndTime                 time.Time       `json:"end_time"`
	CouponId                int64           `json:"coupon_id"`
	StartDate               time.Time       `json:"start_date"`
	EndDate                 time.Time       `json:"end_date"`
	CouponAmount            decimal.Decimal `json:"coupon_amount"`
	ConsumeUsedQuantity     int64           `json:"consume_used_quantity"`
	Remark                  string          `json:"remark"`
	ReceivedQuantity        int64           `json:"received_quantity"`
	ConsumedQuantity        int64           `json:"consumed_quantity"`
	RewardAmount            decimal.Decimal `json:"reward_amount"`
	CouponAvailableQuantity int64           `json:"coupon_available_quantity"`
}
