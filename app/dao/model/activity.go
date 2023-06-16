package model

type Activity struct {
}

func (Activity) TableName() string {
	return "t_promotion_activity"
}
