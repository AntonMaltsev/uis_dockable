package api

type LdapUser struct {
	Id          int32  	`json:"id"`
	User       string 	`json:"ldap_user" binding:"required"`
	Password   string 	`json:"ldap_password" binding:"required"`
	LdapServer string 	`json:"ldap_server"`
	Created 	int32 	`json:"created"`
}

const (
	TodoStatus  string = "todo"
	DoingStatus string = "doing"
	DoneStatus  string = "done"
)
