package service

import (
	"github.com/AntonMaltsev/uis_dockable/api"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

type LdapService struct {
}

func (s *LdapService) getDb(cfg Config) (gorm.DB, error) {
	connectionString := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":3306)/" + cfg.DbName + "?charset=utf8&parseTime=True"

	return gorm.Open("mysql", connectionString)
}

func (s *LdapService) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&api.LdapUser{})
	return nil
}
func (s *LdapService) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	ldapResource := &TodoResource{db: db}

	// creating GIN Ldap router 
	r := gin.Default()	

	r.POST("/ldap", ldapResource.CreateUser)
	r.GET("/ldap", ldapResource.GetAllUsers)
	r.GET("/ldap/:id", ldapResource.GetUser)
	r.DELETE("/ldap/:id", ldapResource.DeleteUser)

	r.Run(cfg.SvcHost)

	return nil
}
