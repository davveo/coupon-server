package coupon

import (
	"github.com/davveo/market-coupon/config"
	"github.com/davveo/market-coupon/pkg/db"
	"github.com/davveo/market-coupon/pkg/gin/wrapper"
	"github.com/gin-gonic/gin"
	rdsV8 "github.com/go-redis/redis/v8"
)

type Controller struct {
}

func NewController(conf *config.Config, db *db.Datastore, redis *rdsV8.Client) *Controller {
	return &Controller{}
}

func (c *Controller) Page(ctx *gin.Context) {
	wrapper.ReplyOK(ctx)
}
