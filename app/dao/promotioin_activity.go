package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type PromotionActivityDao struct {
	db *db.Datastore
}

func NewPromotionActivityDao(db *db.Datastore) *PromotionActivityDao {
	return &PromotionActivityDao{db: db}
}

func (p *PromotionActivityDao) Create(ctx context.Context, mo *model.PromotionActivity) error {
	return p.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (p *PromotionActivityDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return p.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.PromotionActivity{}, id).Updates(updates).Error
}

func (p *PromotionActivityDao) GetById(ctx context.Context, id int64) (*model.PromotionActivity, error) {
	promotionActivity := &model.PromotionActivity{}
	err := p.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(promotionActivity).Error
	return promotionActivity, err
}

func (p *PromotionActivityDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.PromotionActivity, error) {
	var promotionActivities []*model.PromotionActivity
	if err := p.db.DB(ctx).Model(&model.PromotionActivity{}).Where(condition).Find(&promotionActivities).Error; err != nil {
		return nil, err
	}
	return promotionActivities, nil
}
