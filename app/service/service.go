package service

import (
	"github.com/davveo/market-coupon/config"
	"github.com/davveo/market-coupon/pkg/db"
	rdsV8 "github.com/go-redis/redis/v8"
)

type Service struct {
	c     *config.Config
	db    *db.Datastore
	redis *rdsV8.Client
}

func NewService(c *config.Config, db *db.Datastore, redis *rdsV8.Client) *Service {
	return &Service{c: c, db: db, redis: redis}
}

func (s *Service) Conf() *config.Config {
	return s.c
}

func (s *Service) Db() *db.Datastore {
	return s.db
}
