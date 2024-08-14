package view

import "davinci/module/relRoleView"

type ViewBaseInfo struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SourceName  string `json:"sourceName"`
}

type ViewDetail struct{}

type ViewWithSource struct {
	*View
	Source *SourceBaseInfo `json:"source" gorm:"embedded;embeddedPrefix:source."`
}

type ViewWithSourceBaseInfo struct {
	*ViewWithSource
	Roles *[]relRoleView.RelRoleView `json:"roles"`
}

type SourceBaseInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

//type RelRoleView struct {
//	ViewId     int64  `json:"viewId"`
//	RoleId     int64  `json:"roleId"`
//	RowAuth    string `json:"rowAuth"`
//	ColumnAuth string `json:"columnAuth"`
//}
