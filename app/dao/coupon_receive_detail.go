package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type CouponReceiveDetailDao struct {
	db *db.Datastore
}

func NewCouponReceiveDetailDao(db *db.Datastore) *CouponReceiveDetailDao {
	return &CouponReceiveDetailDao{db: db}
}

func (c *CouponReceiveDetailDao) Create(ctx context.Context, mo *model.CouponReceiveDetail) error {
	return c.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (c *CouponReceiveDetailDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return c.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.CouponReceiveDetail{}, id).Updates(updates).Error
}

func (c *CouponReceiveDetailDao) GetById(ctx context.Context, id int64) (*model.CouponReceiveDetail, error) {
	couponReceiveDetail := &model.CouponReceiveDetail{}
	err := c.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(couponReceiveDetail).Error
	return couponReceiveDetail, err
}

func (c *CouponReceiveDetailDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.CouponReceiveDetail, error) {
	var couponReceiveDetails []*model.CouponReceiveDetail
	if err := c.db.DB(ctx).Model(&model.CouponReceiveDetail{}).Where(condition).Find(&couponReceiveDetails).Error; err != nil {
		return nil, err
	}
	return couponReceiveDetails, nil
}
