package source

import (
	"davinci/module/project"
	"davinci/module/user"
	"davinci/pkg/SqlUtils"
	"fmt"
)

var ServiceGetter Service

func init() {
	ServiceGetter = NewServiceGetterImpl()
}

type ServiceGetterImpl struct{}

func (s ServiceGetterImpl) GetDB(u *user.User, id int64) (SqlUtils.IDatabase, error) {
	source, err := s.getSource(id)
	if err != nil {
		return nil, err
	}
	dbType, err := source.GetDBTypeFromJdbcUrl()
	if err != nil {
		return nil, err
	}
	dbUser := source.GetUsername()
	password := source.GetPassword()
	address, err := source.GetIpAddress()
	if err != nil {
		return nil, err
	}
	port, err := source.GetPort()

	if err != nil {
		return nil, err
	}

	db, err := SqlUtils.GetDatabaseInstance(dbType, dbUser, password, address, port)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		return nil, err
	}
	return db, nil
}

func (s ServiceGetterImpl) GetSourceTables(u *user.User, id int64, database string) (*[]string, error) {
	db, err := s.GetDB(u, id)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		return nil, err
	}

	tables, err := db.GetSourceTables(database)
	if err != nil {
		fmt.Println("Error getting tables:", err)
		return nil, err
	}

	fmt.Println("tables", tables)

	return &tables, nil
}

func (s ServiceGetterImpl) GetSourceTableColumns(u *user.User, id int64, database string, table string) (*[]SqlUtils.Column, error) {
	db, err := s.GetDB(u, id)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		return nil, err
	}
	colums, err := db.GetSourceTableColumns(database, table)
	if err != nil {
		fmt.Println("Error getting columns:", err)
		return nil, err
	}
	fmt.Println("columns", colums)
	return &colums, nil
}

func (s ServiceGetterImpl) GetSourceDatabases(u *user.User, id int64) (*[]string, error) {

	db, err := s.GetDB(u, id)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		return nil, err
	}

	databases, err := db.GetSourceDatabases()
	if err != nil {
		fmt.Println("Error getting databases:", err)
		return nil, err
	}

	fmt.Println("databases", databases)

	return &databases, nil
}

func (s ServiceGetterImpl) ExecuteSql(u *user.User, id int64, sql string, variables any, Limit int64) (map[string]interface{}, error) {
	db, err := s.GetDB(u, id)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		return nil, err
	}
	rows, err := db.ExecuteSql(sql)
	if err != nil {
		fmt.Println("Error getting tables:", err)
		return nil, err
	}
	//fmt.Println("rows", rows)
	return rows, nil
}

func (s ServiceGetterImpl) GetSourceDetail(u *user.User, id int64) (*SourceDetail, error) {
	source, err := s.getSource(id)
	if err != nil {
		return nil, err
	}
	// TODO ： 校验当前用户是否有查看权限 缺失了，后续补充上

	sourceDto := ConvertSourcesToSourceDetail(source)
	return sourceDto, nil
}

func (s ServiceGetterImpl) getSource(id int64) (*Source, error) {
	source, err := DaoGetter.GetSourceById(id)
	return source, err
}

func (s ServiceGetterImpl) GetSources(projectId int64, u *user.User) (*[]Source, error) {
	projectDetail, _ := project.ServiceGetter.GetProjectDetail(projectId, u, false)

	fmt.Print(projectDetail)

	sources, err := DaoGetter.GetSourcesByProject(projectId)

	return sources, err
}

func NewServiceGetterImpl() *ServiceGetterImpl {
	return &ServiceGetterImpl{}
}
