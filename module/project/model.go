package project

import (
	"time"
)

type Project struct {
	Id           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"column:name;not null"`
	Description  string    `json:"description" gorm:"column:description"`
	Pic          string    `json:"pic" gorm:"column:pic"`
	OrgId        int64     `json:"orgId" gorm:"column:org_id;not null"`
	UserId       int64     `json:"userId" gorm:"column:user_id;not null"`
	Visibility   bool      `json:"visibility" gorm:"column:visibility;default:true"`
	StarNum      int       `json:"starNum" gorm:"column:star_num;default:0"`
	IsTransfer   bool      `json:"isTransfer" gorm:"column:is_transfer;not null;default:false"`
	InitialOrgId int64     `json:"initialOrgId" gorm:"column:initial_org_id;not null"`
	CreateBy     int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy     int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
}

func (*Project) TableName() string {
	return "project"
}
