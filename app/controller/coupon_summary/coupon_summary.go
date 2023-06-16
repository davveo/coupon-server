package coupon_summary

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

func (c *Controller) ProduceCoupon(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}

func (c *Controller) ProduceCouponDetail(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}
