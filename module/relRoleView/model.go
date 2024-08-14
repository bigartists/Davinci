package relRoleView

import "time"

type RelRoleView struct {
	ViewID     int64     `gorm:"column:view_id;primaryKey" json:"view_id"`
	RoleID     int64     `gorm:"column:role_id;primaryKey" json:"role_id"`
	RowAuth    string    `gorm:"column:row_auth;type:text" json:"row_auth"`
	ColumnAuth string    `gorm:"column:column_auth;type:text" json:"column_auth"`
	CreateBy   int64     `gorm:"column:create_by" json:"create_by,omitempty"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time,omitempty"`
	UpdateBy   int64     `gorm:"column:update_by" json:"update_by,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time,omitempty"`
}

// TableName sets the insert table name for this struct type
func (r *RelRoleView) TableName() string {
	return "rel_role_view"
}
