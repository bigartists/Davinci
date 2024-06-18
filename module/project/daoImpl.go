package project

import (
	"davinci/client"
	"fmt"
	"gorm.io/gorm"
)

var DaoGetter Dao

type DaoGetterImpl struct {
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

func (i *DaoGetterImpl) GetProjectDetail(id int64) (*ProjectDetail, error) {
	var projectDetail ProjectDetail
	fmt.Println("id: ", id)
	err := client.Orm.Table("project p").
		Select(`
			p.*, 
			u.id AS 'createBy.id', 
			IF(u.name IS NULL, u.username, u.name) AS 'createBy.username', 
			u.avatar AS 'createBy.avatar', 
			o.id AS 'organization.id', 
			o.name AS 'organization.name', 
			o.description AS 'organization.description', 
			o.avatar AS 'organization.avatar', 
			o.user_id AS 'organization.userId', 
			o.project_num AS 'organization.projectNum', 
			o.member_num AS 'organization.memberNum', 
			o.role_num AS 'organization.teamNum', 
			o.allow_create_project AS 'organization.allowCreateProject', 
			o.member_permission AS 'organization.memberPermission', 
			o.create_time AS 'organization.createTime', 
			o.create_by AS 'organization.createBy', 
			o.update_time AS 'organization.updateTime', 
			o.update_by AS 'organization.updateBy'
		`).
		Joins("LEFT JOIN organization o ON o.id = p.org_id").
		Joins("LEFT JOIN user u ON u.id = p.user_id").
		Where("p.id = ?", id).
		Scan(&projectDetail).Error

	fmt.Print("projectDetail: ", projectDetail, err)

	if err != nil {
		return nil, err
	}
	return &projectDetail, nil
}

func (*DaoGetterImpl) GetProjectsByUser(userId int64) ([]ProjectWithCreateBy, error) {
	var projects []ProjectWithCreateBy
	db := client.Orm
	baseQuery := selectProjectByUserBaseSql(db, userId)

	err := db.Table("project p").
		Select("p.*, IF(s.id IS NULL, FALSE, TRUE) AS isStar, u.id AS 'createBy.id', IF(u.name IS NULL, u.username, u.name) AS 'createBy.username', u.avatar AS 'createBy.avatar', u.email AS 'createBy.email'").
		Joins("LEFT JOIN user u ON u.id = p.user_id").
		Joins("LEFT JOIN star s ON (s.target_id = p.id AND s.target = 'project' AND s.user_id = ?)", userId).
		Where("p.id IN (?)", baseQuery).
		Order("p.id").
		Scan(&projects).Error

	return projects, err
}

func (*DaoGetterImpl) GetFavoriteProjects(userId int64) ([]ProjectWithCreateBy, error) {
	var projects []ProjectWithCreateBy
	db := client.Orm
	baseQuery := selectProjectByUserBaseSql(db, userId)
	err := db.Raw(`
		SELECT t.* 
		FROM (SELECT p.*, IF(s.id IS NULL, FALSE, TRUE) AS isStar, u.id AS 'createBy.id', 
						IF(u.name IS NULL, u.username, u.name) AS 'createBy.username', u.avatar AS 'createBy.avatar', u.email AS 'createBy.email'
			  FROM project p
			  LEFT JOIN user u ON u.id = p.user_id
			  LEFT JOIN star s ON (s.target_id = p.id AND s.target = 'project' AND s.user_id = ?)
			  WHERE p.id IN (?)
			 ) t 
		LEFT JOIN favorite f ON t.id = f.project_id
		WHERE f.user_id = ?
		ORDER BY t.id`, userId, baseQuery, userId).Scan(&projects).Error
	return projects, err
}

func selectProjectByUserBaseSql(db *gorm.DB, userId int64) *gorm.DB {

	subQuery1 := db.Model(&Project{}).
		Select("DISTINCT p.id").
		Joins("LEFT JOIN rel_project_admin rpa ON rpa.project_id = p.id").
		Where("p.user_id = ? OR rpa.user_id = ?", userId, userId)

	subQuery2 := db.Model(&Project{}).
		Select("DISTINCT p.id").
		Joins("LEFT JOIN rel_role_project rrp ON rrp.project_id = p.id").
		Joins("LEFT JOIN rel_role_user rru ON rru.role_id = rrp.role_id").
		Where("rru.user_id = ?", userId)

	subQuery3 := db.Model(&Project{}).
		Select("DISTINCT p.id").
		Joins("LEFT JOIN rel_user_organization ruo ON ruo.org_id = p.org_id").
		Joins("LEFT JOIN organization o ON o.id = p.org_id").
		Where("o.user_id = ?", userId)

	subQuery4 := db.Model(&Project{}).
		Select("DISTINCT p.id").
		Joins("LEFT JOIN rel_user_organization ruo ON ruo.org_id = p.org_id").
		Joins("LEFT JOIN organization o ON o.id = p.org_id").
		Where("ruo.user_id = ? AND (ruo.role = 1 OR (p.visibility = 1 AND o.member_permission = 1))", userId)

	return db.Raw("(?) UNION ALL (?) UNION ALL (?) UNION ALL (?)", subQuery1, subQuery2, subQuery3, subQuery4).Select("id")
}
