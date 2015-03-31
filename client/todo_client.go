package client

import (
	"github.com/AntonMaltsev/uis_dockable/api"
	"log"
	"strconv"
)

var _ = log.Print

type LdapClient struct {
	Host string
}

// Ldap CRUD 
// Update - todo

//	CREATE
func (tc *LdapClient) CreateLdapUser(user_id string, user_password string, ldap_server string) (api.LdapUser, error) {
	var respLdap api.LdapUser
	todo := api.LdapUser{User: user_id, Password: user_password, LdapServer: ldap_server}

	url := tc.Host + "/ldap"
	r, err := makeRequest("POST", url, todo)
	if err != nil {
		return respLdap, err
	}
	err = processResponseEntity(r, &respLdap, 201)
	return respLdap, err
}

//	READ
func (tc *LdapClient) GetAllLdapUsers() ([]api.LdapUser, error) {
	var respLdaps []api.LdapUser

	url := tc.Host + "/ldap"
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return respLdaps, err
	}
	err = processResponseEntity(r, &respLdaps, 200)
	return respLdaps, err
}

//	READ
func (tc *LdapClient) GetLdapUser(id int32) (api.LdapUser, error) {
	var respLdap api.LdapUser

	url := tc.Host + "/ldap/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return respLdap, err
	}
	err = processResponseEntity(r, &respLdap, 200)
	return respLdap, err
}

// DELETE
func (tc *LdapClient) DeleteLdapUser(id int32) error {
	url := tc.Host + "/ldap/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return processResponse(r, 204)
}
