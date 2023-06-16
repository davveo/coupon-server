package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type CouponDao struct {
	db *db.Datastore
}

func NewCouponDao(db *db.Datastore) *CouponDao {
	return &CouponDao{db: db}
}

func (c *CouponDao) Create(ctx context.Context, mo *model.Coupon) error {
	return c.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (c *CouponDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return c.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.Coupon{}, id).Updates(updates).Error
}

func (c *CouponDao) GetById(ctx context.Context, id int64) (*model.Coupon, error) {
	coupon := &model.Coupon{}
	err := c.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(coupon).Error
	return coupon, err
}

func (c *CouponDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.Coupon, error) {
	var coupons []*model.Coupon
	if err := c.db.DB(ctx).Model(&model.Coupon{}).Where(condition).Find(&coupons).Error; err != nil {
		return nil, err
	}
	return coupons, nil
}
