package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type CouponSummaryDao struct {
	db *db.Datastore
}

func NewCouponSummaryDao(db *db.Datastore) *CouponSummaryDao {
	return &CouponSummaryDao{db: db}
}

func (c *CouponSummaryDao) Create(ctx context.Context, mo *model.CouponSummary) error {
	return c.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (c *CouponSummaryDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return c.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.CouponSummary{}, id).Updates(updates).Error
}

func (c *CouponSummaryDao) GetById(ctx context.Context, id int64) (*model.CouponSummary, error) {
	couponSummary := &model.CouponSummary{}
	err := c.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(couponSummary).Error
	return couponSummary, err
}

func (c *CouponSummaryDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.CouponSummary, error) {
	var couponSummaries []*model.CouponSummary
	if err := c.db.DB(ctx).Model(&model.CouponSummary{}).Where(condition).Find(&couponSummaries).Error; err != nil {
		return nil, err
	}
	return couponSummaries, nil
}
