package cron

import (
	"context"
	"github.com/davveo/market-coupon/app/service"
	"github.com/davveo/market-coupon/pkg/crontab"
	"github.com/davveo/market-coupon/pkg/db"
	"github.com/davveo/market-coupon/pkg/distributed"
	"github.com/davveo/market-coupon/pkg/logger"
)

func Run(service *service.Service) {
	ctx := context.Background()
	logger.Infof(ctx, "run worker start...")

	serviceDB := service.Db()

	srv := distributed.Init(service.Conf().Meta,
		distributed.WithMysql(&db.DB{DB: *serviceDB.DB(ctx)}))
	crontab.Init(crontab.WithSingle(true), crontab.WithDistributed(srv))
	crontab.AddCommand(service.Conf().TestCron, test(ctx))
}

func test(ctx context.Context) crontab.CmdFunc {
	return func() {
		logger.Infof(ctx, "this is a test...")
	}
}
