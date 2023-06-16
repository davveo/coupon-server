package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type CouponReceiveDao struct {
	db *db.Datastore
}

func NewCouponReceiveDao(db *db.Datastore) *CouponReceiveDao {
	return &CouponReceiveDao{db: db}
}

func (c *CouponReceiveDao) Create(ctx context.Context, mo *model.CouponReceive) error {
	return c.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (c *CouponReceiveDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return c.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.CouponReceive{}, id).Updates(updates).Error
}

func (c *CouponReceiveDao) GetById(ctx context.Context, id int64) (*model.CouponReceive, error) {
	couponReceive := &model.CouponReceive{}
	err := c.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(couponReceive).Error
	return couponReceive, err
}

func (c *CouponReceiveDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.CouponReceive, error) {
	var couponReceives []*model.CouponReceive
	if err := c.db.DB(ctx).Model(&model.CouponReceive{}).Where(condition).Find(&couponReceives).Error; err != nil {
		return nil, err
	}
	return couponReceives, nil
}
