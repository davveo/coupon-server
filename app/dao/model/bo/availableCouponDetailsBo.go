package bo

import (
	"github.com/shopspring/decimal"
	"time"
)

type AvailableCouponDetailsBo struct {
	CouponId                    int64           `json:"couponId"`                    // 消费优惠券(模板)记录ID
	CouponReceiveId             int64           `json:"couponReceiveId"`             // 优惠券领劵主记录ID
	CouponReceiveRemainQuantity int64           `json:"couponReceiveRemainQuantity"` // 优惠券领劵剩余数量
	CouponReceiveStatus         int             `json:"couponReceiveStatus"`         // 优惠券领劵记录状态(0-已领取未使用,1-已使用,2-使用中,3-已过期)
	CouponReceiveDetailId       int64           `json:"couponReceiveDetailId"`       // 消费优惠券领劵明细记录ID
	CouponReceiveDetailStatus   int             `json:"couponReceiveDetailStatus"`   // 优惠券领劵明细记录状态(0-已领取未使用,1-已使用,2-使用中,3-已过期)
	CouponNo                    int64           `json:"couponNo"`                    // 优惠券编码,发券时按规则生成
	CouponAmount                decimal.Decimal `json:"couponAmount"`                // 券额(积分数,金额,折扣率等)
	CouponQrCodeUrl             string          `json:"couponQrCodeUrl"`             // 优惠券二维码地址URL
	CouponWeight                int             `json:"couponWeight"`                // 权重值,数值小表示优先
	CouponEndDate               time.Time       `json:"couponEndDate"`               // 优惠券使用截至日期
}
