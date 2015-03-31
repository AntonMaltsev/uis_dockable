package api

type LdapUser struct {
	Id          int32  	`json:"id"`
	User       string 	`json:"ldap_user" binding:"required"`
	Password   string 	`json:"ldap_password" binding:"required"`
	LdapServer string 	`json:"ldap_server"`
	Created 	int32 	`json:"created"`
	Status 		string 	`json:"status"`
}

// Ldap user status
const (
	ActiveStatus  string = "active"
	DisableStatus string = "disabled"
)
