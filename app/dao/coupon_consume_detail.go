package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type CouponConsumeDetailDao struct {
	db *db.Datastore
}

func NewCouponConsumeDetailDao(db *db.Datastore) *CouponConsumeDetailDao {
	return &CouponConsumeDetailDao{db: db}
}

func (c *CouponConsumeDetailDao) Create(ctx context.Context, mo *model.CouponConsumeDetail) error {
	return c.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (c *CouponConsumeDetailDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return c.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.CouponConsumeDetail{}, id).Updates(updates).Error
}

func (c *CouponConsumeDetailDao) GetById(ctx context.Context, id int64) (*model.CouponConsumeDetail, error) {
	couponConsumeDetail := &model.CouponConsumeDetail{}
	err := c.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(couponConsumeDetail).Error
	return couponConsumeDetail, err
}

func (c *CouponConsumeDetailDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.CouponConsumeDetail, error) {
	var couponConsumeDetails []*model.CouponConsumeDetail
	if err := c.db.DB(ctx).Model(&model.CouponConsumeDetail{}).Where(condition).Find(&couponConsumeDetails).Error; err != nil {
		return nil, err
	}
	return couponConsumeDetails, nil
}
