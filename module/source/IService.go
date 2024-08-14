package source

import (
	"davinci/module/user"
	"davinci/pkg/SqlUtils"
)

type Service interface {
	GetSources(projectId int64, u *user.User) (*[]Source, error)
	GetSourceDetail(u *user.User, id int64) (*SourceDetail, error)
	ExecuteSql(u *user.User, id int64, sql string, variables any, Limit int64) (map[string]interface{}, error)
	GetSourceDatabases(u *user.User, id int64) (*[]string, error)
	GetSourceTables(u *user.User, id int64, database string) (*[]string, error)
	GetSourceTableColumns(u *user.User, id int64, database string, table string) (*[]SqlUtils.Column, error)
}
