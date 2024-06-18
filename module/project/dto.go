package project

import (
	"davinci/module/organization"
	"davinci/module/user"
)

type ProjectPure struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Pic          string `json:"pic"`
	OrgId        int64  `json:"orgId"`
	UserId       int64  `json:"userId"`
	Visibility   bool   `json:"visibility"`
	StarNum      int    `json:"starNum"`
	IsTransfer   bool   `json:"isTransfer"`
	InitialOrgId int64  `json:"initialOrgId"`
}

type ProjectWithCreateBy struct {
	*ProjectPure
	IsStar   bool               `json:"isStar"`
	CreateBy user.UserAsCreator `json:"createBy" gorm:"embedded;embeddedPrefix:createBy."` // todo 嵌套是个很实用的技法；这里的CreateBy是一个嵌套结构，它包含了user.UserAsCreator结构体的所有字段
}

type ProjectInfo struct {
	*ProjectWithCreateBy
	ProjectPermission ProjectPermission `json:"permission"`
}

type ProjectPermission struct {
	SourcePermission   int8 `json:"sourcePermission"`
	ViewPermission     int8 `json:"viewPermission"`
	WidgetPermission   int8 `json:"widgetPermission"`
	VizPermission      int8 `json:"vizPermission"`
	SchedulePermission int8 `json:"schedulePermission"`
	SharePermission    bool `json:"sharePermission"`
	DownloadPermission bool `json:"downloadPermission"`
}

type ProjectDetail struct {
	*Project
	CreateBy     user.UserAsCreator        `json:"createBy" gorm:"embedded;embeddedPrefix:createBy."` // todo 嵌套是个很实用的技法；这里的CreateBy是一个嵌套结构，它包含了user.UserAsCreator结构体的所有字段
	Organization organization.Organization `json:"organization" gorm:"embedded;embeddedPrefix:organization."`
}
