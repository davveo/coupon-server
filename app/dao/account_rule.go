package dao

import (
	"context"
	"github.com/davveo/coupon-server/app/constant"
	"github.com/davveo/coupon-server/app/dao/model"
	"github.com/davveo/coupon-server/pkg/db"
	"gorm.io/gorm/clause"
)

type AccountRuleDao struct {
	db *db.Datastore
}

func NewAccountRuleDao(db *db.Datastore) *AccountRuleDao {
	return &AccountRuleDao{db: db}
}

func (a *AccountRuleDao) Create(ctx context.Context, mo *model.AccountRule) error {
	return a.db.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mo).Error
}

func (a *AccountRuleDao) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	return a.db.DB(ctx).Where("is_delete = ?", constant.NotDeleted).
		First(&model.AccountRule{}, id).Updates(updates).Error
}

func (a *AccountRuleDao) GetById(ctx context.Context, id int64) (*model.AccountRule, error) {
	ar := &model.AccountRule{}
	err := a.db.DB(ctx).Where("id = ? and status = ?",
		id, constant.NotDeleted).First(ar).Error
	return ar, err
}

func (a *AccountRuleDao) Find(ctx context.Context, condition map[string]interface{}) ([]*model.AccountRule, error) {
	var ars []*model.AccountRule
	if err := a.db.DB(ctx).Model(&model.AccountRule{}).Where(condition).Find(&ars).Error; err != nil {
		return nil, err
	}
	return ars, nil
}
