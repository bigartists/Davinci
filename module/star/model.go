package star

import "time"

type Star struct {
	ID       int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Target   string    `gorm:"column:target;type:varchar(20);not null" json:"target"`
	TargetID int64     `gorm:"column:target_id;not null" json:"target_id"`
	UserID   int64     `gorm:"column:user_id;not null" json:"user_id"`
	StarTime time.Time `gorm:"column:star_time;not null;autoUpdateTime" json:"star_time"`
}

func (u *Star) TableName() string {
	return "star"
}
