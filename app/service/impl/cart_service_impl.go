package impl

import (
	"context"
	"encoding/json"
	"github.com/davveo/coupon-server/app/dao/model/bo"
	"github.com/davveo/coupon-server/app/service"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/davveo/coupon-server/pkg/logger"
	rdsV8 "github.com/go-redis/redis/v8"
)

type cartServiceImpl struct {
	redisClient *rdsV8.Client
}

// GetCartListFromRedis 获得指定用户的购物车列表
func (c cartServiceImpl) GetCartListFromRedis(ctx context.Context, userId string) ([]*bo.CartBo, error) {
	return c.getCartListFromRedis(ctx, userId)
}

// MergeCartList 将本地购物车合并到用户购物车当中
func (c cartServiceImpl) MergeCartList(ctx context.Context, localCartBoList []*bo.CartBo, memberBo *bo.MemberBo) error {
	logger.Infof(ctx, "MergeCartList start, cartBoList: %+v, memberBo: %+v", localCartBoList, memberBo)
	for _, cartBo := range localCartBoList {
		for _, productBo := range cartBo.ProductBoList {
			if err := c.addGoodsToCartList(ctx, productBo, memberBo); err != nil {
				return err
			}
		}
	}
	return nil
}

// AddGoodsToCartList 添加商品到购物车列表
func (c cartServiceImpl) AddGoodsToCartList(ctx context.Context, productBo *bo.ProductBo, memberBo *bo.MemberBo) error {
	return c.addGoodsToCartList(ctx, productBo, memberBo)
}

// GetAvailableCouponDetailList 获得用户可用优惠券列表
func (c cartServiceImpl) GetAvailableCouponDetailList(ctx context.Context, memberBo *bo.MemberBo) ([]*bo.OrderItemAvailableCouponsRespBo, error) {
	var orderItemAvailableCoupons []*bo.OrderItemAvailableCouponsRespBo
	redisByte, err := c.redisClient.HGet(ctx, "orderItemAvailableCoupons", memberBo.UserID).Bytes()
	if err != nil {
		return orderItemAvailableCoupons, err
	}
	if err = json.Unmarshal(redisByte, &orderItemAvailableCoupons); err != nil {
		return orderItemAvailableCoupons, err
	}
	return orderItemAvailableCoupons, nil
}

func (c cartServiceImpl) searchProductBoBySkuId(ctx context.Context, cartBo *bo.CartBo, skuId string) *bo.ProductBo {
	for _, productBo := range cartBo.ProductBoList {
		if skuId == productBo.SkuId {
			return productBo
		}
	}
	return nil
}

func (c cartServiceImpl) searchCartBoByShopId(ctx context.Context, cartBoList []*bo.CartBo, shopId string) *bo.CartBo {
	for _, cartBo := range cartBoList {
		if cartBo.ShopId == shopId {
			return cartBo
		}
	}
	return nil
}

func (c cartServiceImpl) getCartListFromRedis(ctx context.Context, userId string) ([]*bo.CartBo, error) {
	var cartList []*bo.CartBo
	redisByte, err := c.redisClient.HGet(ctx, "cartList", userId).Bytes()
	if err != nil {
		return cartList, err
	}
	if err = json.Unmarshal(redisByte, &cartList); err != nil {
		return cartList, err
	}
	return cartList, nil
}

func (c cartServiceImpl) saveCartListToRedis(ctx context.Context, userId string, cartBoList []*bo.CartBo) error {
	cartBoBytes, err := json.Marshal(&cartBoList)
	if err != nil {
		return err
	}
	return c.redisClient.HSet(ctx,
		"cartList", userId, cartBoBytes).Err()
}

func (c cartServiceImpl) createCartBo(ctx context.Context, productBo *bo.ProductBo) *bo.CartBo {
	return &bo.CartBo{
		ShopId:        productBo.ShopId,
		ShopName:      productBo.ShopName,
		ProductBoList: []*bo.ProductBo{productBo},
	}
}

// addGoodsToCartList 添加商品到购物车列表
func (c cartServiceImpl) addGoodsToCartList(ctx context.Context, productBo *bo.ProductBo, memberBo *bo.MemberBo) error {
	//从redis中取出该用户购物车列表
	cartBoList, err := c.getCartListFromRedis(ctx, memberBo.UserID)
	if err != nil {
		return err
	}
	if len(cartBoList) > 0 {
		//在购物车列表中搜索该商家购物车
		cartBo := c.searchCartBoByShopId(ctx, cartBoList, productBo.ShopId)
		if cartBo != nil { // 存在该商家购物车
			//在购物车中搜索指定商品
			productBoFromCart := c.searchProductBoBySkuId(ctx, cartBo, productBo.SkuId)
			//已存在该商品
			if productBoFromCart != nil {
				productBoFromCart.Quantity = productBoFromCart.Quantity + productBo.Quantity
			} else { //不存在该商品的订单项
				cartBo.ProductBoList = append(cartBo.ProductBoList, productBo)
			}
		} else { //不存在该商家购物车
			cartBoList = append(cartBoList, c.createCartBo(ctx, productBo))
		}
	} else { //购物车列表为空
		cartBoList = append(cartBoList, c.createCartBo(ctx, productBo))
	}

	if err = c.saveCartListToRedis(ctx, memberBo.UserID, cartBoList); err != nil {
		return err
	}
	return nil
}

func NewCartService(db *db.Datastore, redis *rdsV8.Client) service.CartService {
	return &cartServiceImpl{
		redisClient: redis,
	}
}
