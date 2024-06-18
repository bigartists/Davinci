package role

import (
	"time"
)

type Role struct {
	ID          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrgID       int64     `json:"orgId" gorm:"column:org_id;not null"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Description string    `json:"description" gorm:"column:description"`
	CreateBy    int64     `json:"createBy" gorm:"column:create_by"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateBy    int64     `json:"updateBy" gorm:"column:update_by"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time"`
	Avatar      string    `json:"avatar" gorm:"column:avatar"`
}

func (*Role) TableName() string {
	return "role"
}
