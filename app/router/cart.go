package router

import (
	"github.com/davveo/coupon-server/app/controller/cart"
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

func CartRouterGroup(router *gin.RouterGroup, db *db.Datastore, conf *config.Config, redis *rdsV8.Client) {
	cartGroup := router.Group("/cart")
	cartController := cart.NewController(conf, db, redis)
	{
		// 添加商品到购物车列表
		cartGroup.POST("/addGoodsToCartList", cartController.AddGoodsToCartList)
		// 获取指定用户的购物车列表
		cartGroup.POST("/getCartDtoList", cartController.GetCartDtoList)
		// 获得用户可用优惠券列表
		cartGroup.POST("/getAvailableCouponDetailList", cartController.GetAvailableCouponDetailList)
	}
}
