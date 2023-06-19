package service

import (
	"context"
	"github.com/davveo/coupon-server/app/dao/model/bo"
)

type CartService interface {
	// GetCartListFromRedis 获得指定用户的购物车列表
	GetCartListFromRedis(ctx context.Context, userId string) ([]*bo.CartBo, error)
	// MergeCartList 将本地购物车合并到用户购物车当中
	MergeCartList(ctx context.Context, cartBoList []*bo.CartBo, memberBo *bo.MemberBo) error
	// AddGoodsToCartList 添加商品到购物车列表
	AddGoodsToCartList(ctx context.Context, productBo *bo.ProductBo, memberBo *bo.MemberBo) error
	// GetAvailableCouponDetailList 获得用户可用优惠券列表
	GetAvailableCouponDetailList(ctx context.Context, memberBo *bo.MemberBo) ([]*bo.OrderItemAvailableCouponsRespBo, error)
}
