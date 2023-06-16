package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type CouponConsumeDao struct {
	db *db.Datastore
}

func NewCouponConsumeDao(db *db.Datastore) *CouponConsumeDao {
	return &CouponConsumeDao{db: db}
}

func (c *CouponConsumeDao) Create(ctx context.Context, mo *model.CouponConsume) error {
	return c.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (c *CouponConsumeDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return c.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.CouponConsume{}, id).Updates(updates).Error
}

func (c *CouponConsumeDao) GetById(ctx context.Context, id int64) (*model.CouponConsume, error) {
	couponConsume := &model.CouponConsume{}
	err := c.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(couponConsume).Error
	return couponConsume, err
}

func (c *CouponConsumeDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.CouponConsume, error) {
	var couponConsumes []*model.CouponConsume
	if err := c.db.DB(ctx).Model(&model.CouponConsume{}).Where(condition).Find(&couponConsumes).Error; err != nil {
		return nil, err
	}
	return couponConsumes, nil
}
