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

func (tc *LdapClient) CreateLdapUser(title string, description string) (api.Todo, error) {
	var respTodo api.Todo
	todo := api.Todo{Title: title, Description: description}

	url := tc.Host + "/ldap"
	r, err := makeRequest("POST", url, todo)
	if err != nil {
		return respTodo, err
	}
	err = processResponseEntity(r, &respTodo, 201)
	return respTodo, err
}

func (tc *LdapClient) GetAllLdapUsers() ([]api.Todo, error) {
	var respTodos []api.Todo

	url := tc.Host + "/ldap"
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return respTodos, err
	}
	err = processResponseEntity(r, &respTodos, 200)
	return respTodos, err
}

func (tc *LdapClient) GetLdapUSer(id int32) (api.Todo, error) {
	var respTodo api.Todo

	url := tc.Host + "/ldap/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return respTodo, err
	}
	err = processResponseEntity(r, &respTodo, 200)
	return respTodo, err
}

func (tc *LdapClient) DeleteLdapUser(id int32) error {
	url := tc.Host + "/ldap/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return processResponse(r, 204)
}
