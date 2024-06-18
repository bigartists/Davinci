package relProjectAdmin

import "time"

type RelProjectAdmin struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ProjectID  int64     `gorm:"column:project_id;not null" json:"projectId"`
	UserID     int64     `gorm:"column:user_id;not null" json:"userId"`
	CreateBy   int64     `gorm:"column:create_by" json:"createBy"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateBy   int64     `gorm:"column:update_by" json:"updateBy"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// TableName overrides the table name used by GORM
func (RelProjectAdmin) TableName() string {
	return "rel_project_admin"
}
