package impl

import (
	"context"
	"github.com/davveo/coupon-server/app/biz"
	"github.com/davveo/coupon-server/app/dao/model/bo"
	"github.com/davveo/coupon-server/app/entity"
	"github.com/davveo/coupon-server/app/service"
	"github.com/davveo/coupon-server/app/service/impl"
	"github.com/davveo/coupon-server/pkg/db"
	rdsV8 "github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
)

type cartBiz struct {
	cartService service.CartService
}

func NewCartBiz(db *db.Datastore, redis *rdsV8.Client) biz.CartBiz {
	return &cartBiz{
		cartService: impl.NewCartService(db, redis),
	}
}

// GetCartListFromRedis 获得指定用户的购物车列表
func (c *cartBiz) GetCartListFromRedis(ctx context.Context, userId string) ([]*bo.CartBo, error) {
	return c.cartService.GetCartListFromRedis(ctx, userId)
}

// AddGoodsToCartList 添加商品到购物车列表
func (c *cartBiz) AddGoodsToCartList(ctx context.Context, req *entity.AddCartDto) error {
	memberBo := &bo.MemberBo{}
	if err := copier.Copy(memberBo, req.Member); err != nil {
		return err
	}
	productBo := &bo.ProductBo{}
	if err := copier.Copy(productBo, req.Product); err != nil {
		return err
	}
	if err := c.cartService.AddGoodsToCartList(ctx, productBo, memberBo); err != nil {
		return err
	}
	return nil
}

// MergeCartList 将本地购物车合并到用户购物车当中
func (c *cartBiz) MergeCartList(ctx context.Context, localCartDtoList []*entity.CartDto, memberDto *entity.MemberDto) error {
	memberBo := &bo.MemberBo{}
	if err := copier.Copy(memberBo, memberDto); err != nil {
		return err
	}
	var localCartBoList []*bo.CartBo
	for _, cartDto := range localCartDtoList {
		cartBo := &bo.CartBo{}
		if err := copier.Copy(cartBo, cartDto); err != nil {
			return err
		}
		localCartBoList = append(localCartBoList, cartBo)
	}
	return c.cartService.MergeCartList(ctx, localCartBoList, memberBo)
}

// GetAvailableCouponDetailList 获取指定用户的可用优惠券详情列表
func (c *cartBiz) GetAvailableCouponDetailList(ctx context.Context,
	memberDto *entity.MemberDto) ([]*bo.OrderItemAvailableCouponsRespBo, error) {
	memberBo := &bo.MemberBo{}
	if err := copier.Copy(memberBo, memberDto); err != nil {
		return nil, err
	}
	return c.cartService.GetAvailableCouponDetailList(ctx, memberBo)
}
