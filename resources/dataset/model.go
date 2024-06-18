package dataset

import (
	"time"
)

// CronJob represents the cron_job table
type CronJob struct {
	Id             int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name           string    `json:"name" gorm:"column:name" binding:"required"`
	ProjectId      int64     `json:"project_id" gorm:"column:project_id" binding:"required"`
	JobType        string    `json:"job_type" gorm:"column:job_type" binding:"required"`
	JobStatus      string    `json:"job_status" gorm:"column:job_status" binding:"required"`
	CronExpression string    `json:"cron_expression" gorm:"column:cron_expression" binding:"required"`
	StartDate      time.Time `json:"start_date" gorm:"column:start_date" binding:"required"`
	EndDate        time.Time `json:"end_date" gorm:"column:end_date" binding:"required"`
	Config         string    `json:"config" gorm:"column:config" binding:"required"`
	Description    string    `json:"description" gorm:"column:description"`
	ExecLog        string    `json:"exec_log" gorm:"column:exec_log"`
	CreateBy       int64     `json:"-" gorm:"column:create_by"`
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateBy       int64     `json:"-" gorm:"column:update_by"`
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
	ParentId       int64     `json:"parent_id" gorm:"column:parent_id"`
	FullParentId   string    `json:"full_parent_id" gorm:"column:full_parent_id"`
	IsFolder       int       `json:"is_folder" gorm:"column:is_folder"`
	Index          int       `json:"index" gorm:"column:index"`
}

// Dashboard represents the dashboard table
type Dashboard struct {
	Id                int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name              string    `json:"name" gorm:"column:name" binding:"required"`
	DashboardPortalId int64     `json:"dashboard_portal_id" gorm:"column:dashboard_portal_id" binding:"required"`
	Type              int       `json:"type" gorm:"column:type" binding:"required"`
	Index             int       `json:"index" gorm:"column:index" binding:"required"`
	ParentId          int64     `json:"parent_id" gorm:"column:parent_id" binding:"required"`
	Config            string    `json:"config" gorm:"column:config"`
	FullParentId      string    `json:"full_parent_id" gorm:"column:full_parent_id"`
	CreateBy          int64     `json:"-" gorm:"column:create_by"`
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateBy          int64     `json:"-" gorm:"column:update_by"`
	UpdateTime        time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
}

// DashboardPortal represents the dashboard_portal table
type DashboardPortal struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"column:name" binding:"required"`
	Description string    `json:"description" gorm:"column:description"`
	ProjectId   int64     `json:"project_id" gorm:"column:project_id" binding:"required"`
	Avatar      string    `json:"avatar" gorm:"column:avatar"`
	Publish     int       `json:"publish" gorm:"column:publish"`
	CreateBy    int64     `json:"-" gorm:"column:create_by"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateBy    int64     `json:"-" gorm:"column:update_by"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
}

// Display represents the display table

type Display struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"column:name" binding:"required"`
	Description string    `json:"description" gorm:"column:description"`
	ProjectId   int64     `json:"project_id" gorm:"column:project_id" binding:"required"`
	Avatar      string    `json:"avatar" gorm:"column:avatar"`
	Publish     int       `json:"publish" gorm:"column:publish"`
	CreateBy    int64     `json:"-" gorm:"column:create_by"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateBy    int64     `json:"-" gorm:"column:update_by"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
}

// DisplaySlide represents the display_slide table
type DisplaySlide struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	DisplayId  int64     `json:"display_id" gorm:"column:display_id" binding:"required"`
	Index      int       `json:"index" gorm:"column:index" binding:"required"`
	Config     string    `json:"config" gorm:"column:config"`
	CreateBy   int64     `json:"-" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateBy   int64     `json:"-" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
}

// DownloadRecord represents the download_record table
type DownloadRecord struct {
	Id               int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name             string    `json:"name" gorm:"column:name;not null"`
	UserId           int64     `json:"user_id" gorm:"column:user_id;not null"`
	Path             string    `json:"path" gorm:"column:path"`
	Status           int16     `json:"status" gorm:"column:status;not null"`
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time;not null"`
	LastDownloadTime time.Time `json:"last_download_time" gorm:"column:last_download_time"`
}

// Favorite represents the favorite table
type Favorite struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;not null"`
	ProjectId  int64     `json:"project_id" gorm:"column:project_id;not null"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null;autoCreateTime"`
}

// MemDashboardWidget represents the mem_dashboard_widget table
type MemDashboardWidget struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Alias       string    `json:"alias" gorm:"column:alias"`
	DashboardId int64     `json:"dashboard_id" gorm:"column:dashboard_id;not null"`
	WidgetId    int64     `json:"widget_id" gorm:"column:widget_id"`
	X           int       `json:"x" gorm:"column:x;not null"`
	Y           int       `json:"y" gorm:"column:y;not null"`
	Width       int       `json:"width" gorm:"column:width;not null"`
	Height      int       `json:"height" gorm:"column:height;not null"`
	Polling     bool      `json:"polling" gorm:"column:polling;not null;default:0"`
	Frequency   int       `json:"frequency" gorm:"column:frequency"`
	Config      string    `json:"config" gorm:"column:config"`
	CreateBy    int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy    int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

// MemDisplaySlideWidget represents the mem_display_slide_widget table
type MemDisplaySlideWidget struct {
	Id             int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	DisplaySlideId int64     `json:"display_slide_id" gorm:"column:display_slide_id;not null"`
	WidgetId       int64     `json:"widget_id" gorm:"column:widget_id"`
	Name           string    `json:"name" gorm:"column:name;not null"`
	Params         string    `json:"params" gorm:"column:params;not null"`
	Type           int16     `json:"type" gorm:"column:type;not null"`
	SubType        int16     `json:"sub_type" gorm:"column:sub_type"`
	Index          int       `json:"index" gorm:"column:index;not null;default:0"`
	CreateBy       int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy       int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time"`
}

// Organization represents the organization table
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

// Platform represents the platform table
type Platform struct {
	Id               int64  `json:"id" gorm:"column:id;primaryKey"`
	Name             string `json:"name" gorm:"column:name;not null"`
	Platform         string `json:"platform" gorm:"column:platform;not null"`
	Code             string `json:"code" gorm:"column:code;not null"`
	CheckCode        string `json:"check_code" gorm:"column:check_code"`
	CheckSystemToken string `json:"check_system_token" gorm:"column:check_system_token"`
	CheckUrl         string `json:"check_url" gorm:"column:check_url"`
	AlternateField1  string `json:"alternate_field1" gorm:"column:alternate_field1"`
	AlternateField2  string `json:"alternate_field2" gorm:"column:alternate_field2"`
	AlternateField3  string `json:"alternate_field3" gorm:"column:alternate_field3"`
	AlternateField4  string `json:"alternate_field4" gorm:"column:alternate_field4"`
	AlternateField5  string `json:"alternate_field5" gorm:"column:alternate_field5"`
}

// Project represents the project table
type Project struct {
	Id           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"column:name;not null"`
	Description  string    `json:"description" gorm:"column:description"`
	Pic          string    `json:"pic" gorm:"column:pic"`
	OrgId        int64     `json:"org_id" gorm:"column:org_id;not null"`
	UserId       int64     `json:"user_id" gorm:"column:user_id;not null"`
	Visibility   bool      `json:"visibility" gorm:"column:visibility;default:true"`
	StarNum      int       `json:"star_num" gorm:"column:star_num;default:0"`
	IsTransfer   bool      `json:"is_transfer" gorm:"column:is_transfer;not null;default:false"`
	InitialOrgId int64     `json:"initial_org_id" gorm:"column:initial_org_id;not null"`
	CreateBy     int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy     int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelProjectAdmin represents the rel_project_admin table
type RelProjectAdmin struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProjectId  int64     `json:"project_id" gorm:"column:project_id;not null"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;not null"`
	CreateBy   int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy   int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRoleDashboard represents the rel_role_dashboard table
type RelRoleDashboard struct {
	RoleId      int64     `json:"role_id" gorm:"column:role_id;primaryKey"`
	DashboardId int64     `json:"dashboard_id" gorm:"column:dashboard_id;primaryKey"`
	Visible     bool      `json:"visible" gorm:"column:visible;default:false"`
	CreateBy    int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy    int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRoleDisplay represents the rel_role_display table
type RelRoleDisplay struct {
	RoleId     int64     `json:"role_id" gorm:"column:role_id;primaryKey"`
	DisplayId  int64     `json:"display_id" gorm:"column:display_id;primaryKey"`
	Visible    bool      `json:"visible" gorm:"column:visible;default:false"`
	CreateBy   int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy   int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRolePortal represents the rel_role_portal table
type RelRolePortal struct {
	RoleId     int64     `json:"role_id" gorm:"column:role_id;primaryKey"`
	PortalId   int64     `json:"portal_id" gorm:"column:portal_id;primaryKey"`
	Visible    bool      `json:"visible" gorm:"column:visible;default:false"`
	CreateBy   int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy   int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRoleProject represents the rel_role_project table
type RelRoleProject struct {
	Id                 int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProjectId          int64     `json:"project_id" gorm:"column:project_id;not null"`
	RoleId             int64     `json:"role_id" gorm:"column:role_id;not null"`
	SourcePermission   int       `json:"source_permission" gorm:"column:source_permission;default:1"`
	ViewPermission     int       `json:"view_permission" gorm:"column:view_permission;default:1"`
	WidgetPermission   int       `json:"widget_permission" gorm:"column:widget_permission;default:1"`
	VizPermission      int       `json:"viz_permission" gorm:"column:viz_permission;default:1"`
	SchedulePermission int       `json:"schedule_permission" gorm:"column:schedule_permission;default:1"`
	SharePermission    bool      `json:"share_permission" gorm:"column:share_permission;default:false"`
	DownloadPermission bool      `json:"download_permission" gorm:"column:download_permission;default:false"`
	CreateBy           int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy           int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime         time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRoleSlide represents the rel_role_slide table
type RelRoleSlide struct {
	RoleId     int64     `json:"role_id" gorm:"column:role_id;primaryKey"`
	SlideId    int64     `json:"slide_id" gorm:"column:slide_id;primaryKey"`
	Visible    bool      `json:"visible" gorm:"column:visible;default:false"`
	CreateBy   int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy   int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRoleUser represents the rel_role_user table
type RelRoleUser struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;not null"`
	RoleId     int64     `json:"role_id" gorm:"column:role_id;not null"`
	CreateBy   int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy   int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRoleView represents the rel_role_view table
type RelRoleView struct {
	ViewId     int64     `json:"view_id" gorm:"column:view_id;primaryKey"`
	RoleId     int64     `json:"role_id" gorm:"column:role_id;primaryKey"`
	RowAuth    string    `json:"row_auth" gorm:"column:row_auth;type:text"`
	ColumnAuth string    `json:"column_auth" gorm:"column:column_auth;type:text"`
	CreateBy   int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy   int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelUserOrganization represents the rel_user_organization table
type RelUserOrganization struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrgId      int64     `json:"org_id" gorm:"column:org_id;not null"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;not null"`
	Role       int       `json:"role" gorm:"column:role;default:0"`
	CreateBy   int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy   int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// Role represents the role table
type Role struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrgId       int64     `json:"org_id" gorm:"column:org_id;not null"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Description string    `json:"description" gorm:"column:description"`
	CreateBy    int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy    int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
	Avatar      string    `json:"avatar" gorm:"column:avatar"`
}

// Source represents the source table
type Source struct {
	Id           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"column:name;not null"`
	Description  string    `json:"description" gorm:"column:description"`
	Config       string    `json:"config" gorm:"column:config;type:text;not null"`
	Type         string    `json:"type" gorm:"column:type;not null"`
	ProjectId    int64     `json:"project_id" gorm:"column:project_id;not null"`
	CreateBy     int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy     int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
	ParentId     int64     `json:"parent_id" gorm:"column:parent_id"`
	FullParentId string    `json:"full_parent_id" gorm:"column:full_parent_id"`
	IsFolder     bool      `json:"is_folder" gorm:"column:is_folder"`
	Index        int       `json:"index" gorm:"column:index"`
}

// Star represents the star table
type Star struct {
	Id       int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Target   string    `json:"target" gorm:"column:target;not null"`
	TargetId int64     `json:"target_id" gorm:"column:target_id;not null"`
	UserId   int64     `json:"user_id" gorm:"column:user_id;not null"`
	StarTime time.Time `json:"star_time" gorm:"column:star_time;not null;autoUpdateTime"`
}

// View represents the view table
type View struct {
	Id           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"column:name;not null"`
	Description  string    `json:"description" gorm:"column:description"`
	ProjectId    int64     `json:"project_id" gorm:"column:project_id;not null"`
	SourceId     int64     `json:"source_id" gorm:"column:source_id;not null"`
	SQL          string    `json:"sql" gorm:"column:sql;type:text"`
	Model        string    `json:"model" gorm:"column:model;type:text"`
	Variable     string    `json:"variable" gorm:"column:variable;type:text"`
	Config       string    `json:"config" gorm:"column:config;type:text"`
	CreateBy     int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy     int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
	ParentId     int64     `json:"parent_id" gorm:"column:parent_id"`
	FullParentId string    `json:"full_parent_id" gorm:"column:full_parent_id"`
	IsFolder     bool      `json:"is_folder" gorm:"column:is_folder"`
	Index        int       `json:"index" gorm:"column:index"`
}

// Widget represents the widget table
type Widget struct {
	Id           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"column:name;not null"`
	Description  string    `json:"description" gorm:"column:description"`
	ViewId       int64     `json:"view_id" gorm:"column:view_id;not null"`
	ProjectId    int64     `json:"project_id" gorm:"column:project_id;not null"`
	Type         int64     `json:"type" gorm:"column:type;not null"`
	Publish      bool      `json:"publish" gorm:"column:publish;not null"`
	Config       string    `json:"config" gorm:"column:config;type:longtext;not null"`
	CreateBy     int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy     int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
	ParentId     int64     `json:"parent_id" gorm:"column:parent_id"`
	FullParentId string    `json:"full_parent_id" gorm:"column:full_parent_id"`
	IsFolder     bool      `json:"is_folder" gorm:"column:is_folder"`
	Index        int       `json:"index" gorm:"column:index"`
}

// RelRoleDisplaySlideWidget represents the rel_role_display_slide_widget table
type RelRoleDisplaySlideWidget struct {
	RoleId                  int64     `json:"role_id" gorm:"column:role_id;primaryKey"`
	MemDisplaySlideWidgetId int64     `json:"mem_display_slide_widget_id" gorm:"column:mem_display_slide_widget_id;primaryKey"`
	Visible                 bool      `json:"visible" gorm:"column:visible;default:0"`
	CreateBy                int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime              time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy                int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime              time.Time `json:"update_time" gorm:"column:update_time"`
}

// RelRoleDashboardWidget represents the rel_role_dashboard_widget table
type RelRoleDashboardWidget struct {
	RoleId               int64     `json:"role_id" gorm:"column:role_id;primaryKey"`
	MemDashboardWidgetId int64     `json:"mem_dashboard_widget_id" gorm:"column:mem_dashboard_widget_id;primaryKey"`
	Visible              bool      `json:"visible" gorm:"column:visible;not null;default:0"`
	CreateBy             int64     `json:"create_by" gorm:"column:create_by"`
	CreateTime           time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateBy             int64     `json:"update_by" gorm:"column:update_by"`
	UpdateTime           time.Time `json:"update_time" gorm:"column:update_time"`
}

// DavinciStatisticTerminal represents the davinci_statistic_terminal table
type DavinciStatisticTerminal struct {
	Id              int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId          int64     `json:"user_id" gorm:"column:user_id"`
	Email           string    `json:"email" gorm:"column:email"`
	BrowserName     string    `json:"browser_name" gorm:"column:browser_name"`
	BrowserVersion  string    `json:"browser_version" gorm:"column:browser_version"`
	EngineName      string    `json:"engine_name" gorm:"column:engine_name"`
	EngineVersion   string    `json:"engine_version" gorm:"column:engine_version"`
	OsName          string    `json:"os_name" gorm:"column:os_name"`
	OsVersion       string    `json:"os_version" gorm:"column:os_version"`
	DeviceModel     string    `json:"device_model" gorm:"column:device_model"`
	DeviceType      string    `json:"device_type" gorm:"column:device_type"`
	DeviceVendor    string    `json:"device_vendor" gorm:"column:device_vendor"`
	CpuArchitecture string    `json:"cpu_architecture" gorm:"column:cpu_architecture"`
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
}

// DavinciStatisticDuration represents the davinci_statistic_duration table
type DavinciStatisticDuration struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId      int64     `json:"user_id" gorm:"column:user_id"`
	Email       string    `json:"email" gorm:"column:email"`
	OrgId       int64     `json:"org_id" gorm:"column:org_id"`
	ProjectId   int64     `json:"project_id" gorm:"column:project_id"`
	ProjectName string    `json:"project_name" gorm:"column:project_name"`
	VizType     string    `json:"viz_type" gorm:"column:viz_type"`
	VizId       int64     `json:"viz_id" gorm:"column:viz_id"`
	VizName     string    `json:"viz_name" gorm:"column:viz_name"`
	SubVizId    int64     `json:"sub_viz_id" gorm:"column:sub_viz_id"`
	SubVizName  string    `json:"sub_viz_name" gorm:"column:sub_viz_name"`
	StartTime   time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime     time.Time `json:"end_time" gorm:"column:end_time"`
}

type DavinciStatisticVisitorOperation struct {
	Id          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId      int64     `json:"user_id" gorm:"column:user_id"`
	Email       string    `json:"email" gorm:"column:email"`
	Action      string    `json:"action" gorm:"column:action"`
	OrgId       int64     `json:"org_id" gorm:"column:org_id"`
	ProjectId   int64     `json:"project_id" gorm:"column:project_id"`
	ProjectName string    `json:"project_name" gorm:"column:project_name"`
	VizType     string    `json:"viz_type" gorm:"column:viz_type"`
	VizId       int64     `json:"viz_id" gorm:"column:viz_id"`
	VizName     string    `json:"viz_name" gorm:"column:viz_name"`
	SubVizId    int64     `json:"sub_viz_id" gorm:"column:sub_viz_id"`
	SubVizName  string    `json:"sub_viz_name" gorm:"column:sub_viz_name"`
	WidgetId    int64     `json:"widget_id" gorm:"column:widget_id"`
	WidgetName  string    `json:"widget_name" gorm:"column:widget_name"`
	Variables   string    `json:"variables" gorm:"column:variables"`
	Filters     string    `json:"filters" gorm:"column:filters"`
	Groups      string    `json:"groups" gorm:"column:groups"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}

// ShareDownloadRecord represents the share_download_record table
type ShareDownloadRecord struct {
	Id               int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Uuid             string    `json:"uuid" gorm:"column:uuid"`
	Name             string    `json:"name" gorm:"column:name;not null"`
	Path             string    `json:"path" gorm:"column:path"`
	Status           int16     `json:"status" gorm:"column:status;not null"`
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time;not null"`
	LastDownloadTime time.Time `json:"last_download_time" gorm:"column:last_download_time"`
}
