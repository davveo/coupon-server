package cart

import (
	"github.com/davveo/coupon-server/app/biz"
	"github.com/davveo/coupon-server/app/biz/impl"
	"github.com/davveo/coupon-server/app/entity"
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/davveo/coupon-server/pkg/gin/code"
	"github.com/davveo/coupon-server/pkg/gin/wrapper"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

type Controller struct {
	cartBiz biz.CartBiz
}

func NewController(conf *config.Config, db *db.Datastore, redis *rdsV8.Client) *Controller {
	return &Controller{
		cartBiz: impl.NewCartBiz(db, redis),
	}
}

// AddGoodsToCartList 添加商品到购物车列表
func (c *Controller) AddGoodsToCartList(ctx *gin.Context) {
	var addCart entity.AddCartDto
	if err := ctx.ShouldBindJSON(&addCart); err != nil {
		wrapper.ReplyErr(ctx, code.ErrBadParams, err.Error())
		return
	}
	if err := c.cartBiz.AddGoodsToCartList(ctx, &addCart); err != nil {
		wrapper.ReplyErr(ctx, code.ErrInternal, err.Error())
		return
	}
	wrapper.ReplyOK(ctx)
}

// GetCartDtoList 获取指定用户的购物车列表
func (c *Controller) GetCartDtoList(ctx *gin.Context) {
	var addCart entity.CartDtoList
	if err := ctx.ShouldBindJSON(&addCart); err != nil {
		wrapper.ReplyErr(ctx, code.ErrBadParams, err.Error())
		return
	}
	if len(addCart.LocalCartList) != 0 {
		if err := c.cartBiz.MergeCartList(ctx,
			addCart.LocalCartList, addCart.MemberDto); err != nil {
			wrapper.ReplyErr(ctx, code.ErrInternal, err.Error())
			return
		}
	}
	res, err := c.cartBiz.GetCartListFromRedis(ctx, addCart.MemberDto.UserID)
	if err != nil {
		wrapper.ReplyErr(ctx, code.ErrInternal, err.Error())
		return
	}
	wrapper.Reply(ctx, res)
}

// GetAvailableCouponDetailList 获得用户可用优惠券列表
func (c *Controller) GetAvailableCouponDetailList(ctx *gin.Context) {
	var memberDto entity.MemberDto
	if err := ctx.ShouldBindJSON(&memberDto); err != nil {
		wrapper.ReplyErr(ctx, code.ErrBadParams, err.Error())
		return
	}

	res, err := c.cartBiz.GetAvailableCouponDetailList(ctx, &memberDto)
	if err != nil {
		wrapper.ReplyErr(ctx, code.ErrInternal, err.Error())
		return
	}
	wrapper.Reply(ctx, res)
}
