package source

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

type Source struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description  string    `gorm:"column:description;type:varchar(255)" json:"description"`
	Config       string    `gorm:"column:config;type:text;not null" json:"config"`
	Type         string    `gorm:"column:type;type:varchar(10);not null" json:"type"`
	ProjectID    int64     `gorm:"column:project_id;not null" json:"project_id"`
	CreateBy     int64     `gorm:"column:create_by" json:"create_by"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateBy     int64     `gorm:"column:update_by" json:"update_by"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	ParentID     int64     `gorm:"column:parent_id" json:"parent_id"`
	FullParentID string    `gorm:"column:full_parent_id;type:varchar(255)" json:"full_parent_id"`
	IsFolder     bool      `gorm:"column:is_folder;type:tinyint(1)" json:"is_folder"`
	Index        int       `gorm:"column:index" json:"index"`
}

func (*Source) TableName() string {
	return "source"
}

const JDBC_DATASOURCE_DEFAULT_VERSION = "default"

// "url":"jdbc:mysql://101.43.2.58"
func SourceUtilsGetDataSourceName(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

type Config struct {
	Ext        bool   `json:"ext"`
	Password   string `json:"password"`
	Version    string `json:"version"`
	Properties []Dict `json:"properties"`
	Url        string `json:"url"`
	Username   string `json:"username"`
}

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//{"ext":false,"password":"Alex_13306","version":"","properties":[],"url":"jdbc:mysql://101.43.2.58","username":"root"}

func (s *Source) GetConfig() (*Config, error) {
	var config Config
	if err := json.Unmarshal([]byte(s.Config), &config); err != nil {
		log.Printf("Error parsing config: %v", err)
		return nil, err
	}
	return &config, nil
}

func (s *Source) GetJdbcUrl() string {
	config, err := s.GetConfig()
	if err != nil {
		return ""
	}
	return config.Url
}

func (s *Source) GetUsername() string {
	config, err := s.GetConfig()
	if err != nil {
		return ""
	}
	return config.Username
}

func (s *Source) GetPassword() string {
	config, err := s.GetConfig()
	if err != nil {
		return ""
	}
	return config.Password
}

func (s *Source) GetDatabase() string {
	// Assuming SourceUtils.GetDataSourceName is a method that extracts the database name from the URL
	return SourceUtilsGetDataSourceName(s.GetJdbcUrl())
}

func (s *Source) GetDBTypeFromJdbcUrl() (string, error) {
	jdbcUrl := s.GetJdbcUrl()
	re := regexp.MustCompile(`jdbc:([^:]+)://`)
	matches := re.FindStringSubmatch(jdbcUrl)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid JDBC URL: %s", jdbcUrl)
	}
	return matches[1], nil
}

func (s *Source) GetIpAddress() (string, error) {
	jdbcUrl := s.GetJdbcUrl()
	re := regexp.MustCompile(`jdbc:[^:]+://([^:/]+)`)
	matches := re.FindStringSubmatch(jdbcUrl)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid JDBC URL: %s", jdbcUrl)
	}
	return matches[1], nil
}

// getPort extracts the port from the JDBC URL
func (s *Source) GetPort() (string, error) {
	jdbcUrl := s.GetJdbcUrl()
	re := regexp.MustCompile(`jdbc:[^:]+://[^:/]+:(\d+)`)
	matches := re.FindStringSubmatch(jdbcUrl)
	if len(matches) < 2 {
		// No port specified, provide default port based on DB type
		dbType, err := s.GetDBTypeFromJdbcUrl()
		if err != nil {
			return "", err
		}
		switch dbType {
		case "mysql":
			return "3306", nil
		case "postgresql":
			return "5432", nil
		// Add other database types and their default ports as needed
		default:
			return "", fmt.Errorf("port not specified in JDBC URL and default port for %s is unknown", dbType)
		}
	}
	return matches[1], nil
}

func (s *Source) GetDbVersion() string {
	config, err := s.GetConfig()
	if err != nil {
		return ""
	}
	if config.Version == JDBC_DATASOURCE_DEFAULT_VERSION {
		return ""
	}
	return config.Version
}

func (s *Source) IsExt() bool {
	config, err := s.GetConfig()
	if err != nil {
		return false
	}
	return config.Ext
}

func (s *Source) GetProperties() []Dict {
	config, err := s.GetConfig()
	if err != nil {
		return nil
	}
	return config.Properties
}
