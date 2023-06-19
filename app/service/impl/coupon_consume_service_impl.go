package impl

import (
	"context"
	"github.com/davveo/coupon-server/app/dao/model/bo"
	"github.com/davveo/coupon-server/app/service"
	"github.com/davveo/coupon-server/pkg/db"
	rdsV8 "github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
)

type CouponConsumeServiceImpl struct {
}

func (c *CouponConsumeServiceImpl) QueryOrderConsumedCoupon(ctx context.Context, orderId string) (*bo.CouponConsumeHistoryRespBo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) QueryMatchedCouponForPurchaseItem(ctx context.Context,
	memberReqBo *bo.MemberBo, orderItem *bo.OrderItemBo) (*bo.OrderItemAvailableCouponsRespBo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) QueryMatchedCouponForOrder(ctx context.Context,
	memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo) (*bo.OrderAvailableCouponsRespBo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) GetSelectCouponReceiveInfoList(ctx context.Context,
	memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo, orderItemReqBo *bo.OrderItemBo) ([]*bo.SelectCouponReceiveInfoBo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) GetTotalCouponValidRemainQuantity(ctx context.Context,
	memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo, orderItemReqBo *bo.OrderItemBo, couponId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) GetSelectCouponUseList(ctx context.Context, orderReqBo *bo.OrderBo,
	orderItemReqBo *bo.OrderItemBo, memberReqBo *bo.MemberBo, selectCouponReceiveInfoList []*bo.SelectCouponReceiveInfoBo) ([]*bo.SelectCouponUseBo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) SortAvailableMatchedCoupons(ctx context.Context,
	availableCouponList []*bo.AvailableCouponDetailsBo) ([]*bo.AvailableCouponDetailsBo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) ConsumeCoupon(ctx context.Context,
	memberReqBo *bo.MemberBo, orderReqBo *bo.OrderBo) ([]*bo.CouponConsumeBo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponConsumeServiceImpl) ValidateCouponAvailable(ctx context.Context, memberReqBo *bo.MemberBo,
	orderReqBo *bo.OrderBo, isSettleCheck bool, couponReduceAmountMap map[*bo.CouponReduceInfoBo]map[int64]decimal.Decimal) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewCouponConsumeService(db *db.Datastore, redis *rdsV8.Client) service.CouponConsumeService {
	return &CouponConsumeServiceImpl{}
}
