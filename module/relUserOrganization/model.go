package relUserOrganization

import (
	"fmt"
	"time"
)

type RelUserOrganization struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrgID      int64     `gorm:"column:org_id;not null" json:"orgId"`
	UserID     int64     `gorm:"column:user_id;not null" json:"userId"`
	Role       int16     `gorm:"column:role;not null;default:0" json:"role"`
	CreateBy   int64     `gorm:"column:create_by" json:"createBy"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateBy   int64     `gorm:"column:update_by" json:"updateBy"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// TableName sets the insert table name for this struct type
func (RelUserOrganization) TableName() string {
	return "rel_user_organization"
}

const (
	MemberRole = 0
	OwnerRole  = 1
)

var UserOrgRoleMap = map[int]string{
	MemberRole: "member",
	OwnerRole:  "owner",
}

func RoleOf(role int) (string, error) {
	roleName, ok := UserOrgRoleMap[role]
	if !ok {
		return "", fmt.Errorf("role %d not found", role)
	}
	return roleName, nil
}
