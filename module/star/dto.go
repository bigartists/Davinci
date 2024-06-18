package star

type OrganizationInfo struct {
	AllowCreateProject bool   `json:"allow_create_project" gorm:"column:allow_create_project;default:1"`
	Avatar             string `json:"avatar" gorm:"column:avatar"`
	Description        string `json:"description" gorm:"column:description"`
	Id                 int64  `json:"id" `
	MemberNum          int    `json:"member_num" gorm:"column:member_num;default:0"`
	MemberPermission   int16  `json:"member_permission" gorm:"column:member_permission;not null;default:0"`
	Name               string `json:"name" `
	ProjectNum         int    `json:"project_num" gorm:"column:project_num;default:0"`
	Role               int    `json:"role"`
	RoleNum            int    `json:"role_num" gorm:"column:role_num;default:0"`
}
