package cart

import (
	"github.com/davveo/coupon-server/config"
	"github.com/davveo/coupon-server/pkg/db"
	"github.com/davveo/coupon-server/pkg/gin/wrapper"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

type Controller struct {
}

func NewController(conf *config.Config, db *db.Datastore, redis *rdsV8.Client) *Controller {
	return &Controller{}
}

func (c *Controller) AddGoodsToCartList(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}

func (c *Controller) GetCartDtoList(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}

func (c *Controller) GetAvailableCouponDetailList(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}
