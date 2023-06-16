package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type PromotionRuleDao struct {
	db *db.Datastore
}

func NewPromotionRuleDao(db *db.Datastore) *PromotionRuleDao {
	return &PromotionRuleDao{db: db}
}

func (p *PromotionRuleDao) Create(ctx context.Context, mo *model.PromotionRule) error {
	return p.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (p *PromotionRuleDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return p.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.PromotionRule{}, id).Updates(updates).Error
}

func (p *PromotionRuleDao) GetById(ctx context.Context, id int64) (*model.PromotionRule, error) {
	promotionRule := &model.PromotionRule{}
	err := p.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(promotionRule).Error
	return promotionRule, err
}

func (p *PromotionRuleDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.PromotionRule, error) {
	var promotionRules []*model.PromotionRule
	if err := p.db.DB(ctx).Model(&model.PromotionRule{}).Where(condition).Find(&promotionRules).Error; err != nil {
		return nil, err
	}
	return promotionRules, nil
}
