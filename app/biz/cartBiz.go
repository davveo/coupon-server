package biz

import (
	"context"
	"github.com/davveo/coupon-server/app/dao/model/bo"
	"github.com/davveo/coupon-server/app/entity"
)

type CartBiz interface {
	AddGoodsToCartList(ctx context.Context, req *entity.AddCartDto) error
	GetCartListFromRedis(ctx context.Context, userId string) ([]*bo.CartBo, error)
	MergeCartList(ctx context.Context, cartDtoList []*entity.CartDto, memberDto *entity.MemberDto) error
	GetAvailableCouponDetailList(ctx context.Context, memberDto *entity.MemberDto) ([]*bo.OrderItemAvailableCouponsRespBo, error)
}
