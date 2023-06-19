package impl

import (
	"context"
	"github.com/davveo/coupon-server/app/biz"
	"github.com/davveo/coupon-server/app/entity"
	"github.com/davveo/coupon-server/app/service"
	"github.com/davveo/coupon-server/app/service/impl"
	"github.com/davveo/coupon-server/pkg/db"
	rdsV8 "github.com/go-redis/redis/v8"
)

type couponConsumeBiz struct {
	couponConsumeService service.CouponConsumeService
}

func (c couponConsumeBiz) QueryOrderConsumedCoupon(ctx context.Context, req *entity.QueryOrderConsumedCouponReqDto) {
	//TODO implement me
	panic("implement me")
}

func (c couponConsumeBiz) QueryMatchedCouponForPurchaseItem(ctx context.Context, req *entity.QueryItemMatchedCouponReqDto) {
	//TODO implement me
	panic("implement me")
}

func (c couponConsumeBiz) QueryMatchedCouponForOrder(ctx context.Context, req *entity.QueryOrderMatchedCouponDto) {
	//TODO implement me
	panic("implement me")
}

func (c couponConsumeBiz) SortAvailableMatchedCoupons(ctx context.Context, req *entity.SortAvailableMatchedCouponsReqDto) {
	//TODO implement me
	panic("implement me")
}

func (c couponConsumeBiz) ConsumeCoupon(ctx context.Context, req *entity.ConsumeCouponReqDto) {
	//TODO implement me
	panic("implement me")
}

func NewCouponConsumeBiz(db *db.Datastore, redis *rdsV8.Client) biz.CouponConsumeBiz {
	return &couponConsumeBiz{
		couponConsumeService: impl.NewCouponConsumeService(db, redis),
	}
}
