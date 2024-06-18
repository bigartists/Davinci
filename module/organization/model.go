package organization

import (
	"time"
)

type Organization struct {
	Id                 int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name               string    `json:"name" gorm:"column:name;not null"`
	Description        string    `json:"description" gorm:"column:description"`
	Avatar             string    `json:"avatar" gorm:"column:avatar"`
	UserId             int64     `json:"user_id" gorm:"column:user_id;not null"`
	ProjectNum         int       `json:"project_num" gorm:"column:project_num;default:0"`
	MemberNum          int       `json:"member_num" gorm:"column:member_num;default:0"`
	RoleNum            int       `json:"role_num" gorm:"column:role_num;default:0"`
	AllowCreateProject bool      `json:"allow_create_project" gorm:"column:allow_create_project;default:1"`
	MemberPermission   int16     `json:"member_permission" gorm:"column:member_permission;not null;default:0"`
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP"`
	CreateBy           int64     `json:"create_by" gorm:"column:create_by;not null;default:0"`
	UpdateTime         time.Time `json:"update_time" gorm:"column:update_time"`
	UpdateBy           int64     `json:"update_by" gorm:"column:update_by"`
}

func (*Organization) TableName() string {
	return "Organization"
}

type UserWithOrgRole struct {
	Id       int64  `json:"id"`
	Role     int16  `json:"role"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type OrganizationWithUser struct {
	Id   int64            `json:"id"`
	User *UserWithOrgRole `json:"user" gorm:"embedded;embeddedPrefix:user."`
}
