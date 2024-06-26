package view

import (
	"time"
)

type View struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description  string    `gorm:"column:description;type:varchar(255)" json:"description,omitempty"`
	ProjectID    int64     `gorm:"column:project_id;not null" json:"projectId"`
	SourceID     int64     `gorm:"column:source_id;not null" json:"sourceId"`
	SQL          string    `gorm:"column:sql;type:text" json:"sql,omitempty"`
	Model        string    `gorm:"column:model;type:text" json:"model,omitempty"`
	Variable     string    `gorm:"column:variable;type:text" json:"variable,omitempty"`
	Config       string    `gorm:"column:config;type:text" json:"config"`
	CreateBy     int64     `gorm:"column:create_by" json:"-"`
	CreateTime   time.Time `gorm:"column:create_time" json:"-"`
	UpdateBy     int64     `gorm:"column:update_by" json:"-"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"-"`
	ParentID     int64     `gorm:"column:parent_id" json:"parent_id,omitempty"`
	FullParentID string    `gorm:"column:full_parent_id;type:varchar(255)" json:"full_parent_id,omitempty"`
	IsFolder     bool      `gorm:"column:is_folder;type:tinyint(1)" json:"is_folder,omitempty"`
	Index        int       `gorm:"column:index" json:"index,omitempty"`
}

func (*View) TableName() string {
	return "view"
}
