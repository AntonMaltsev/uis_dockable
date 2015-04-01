package main

import (
	"fmt"
	"github.com/AntonMaltsev/uis_dockable/client"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestGetLdapUser(t *testing.T) {

	// given
	client := client.LdapClient{Host: "http://localhost:8080"}
	todo, _ := client.CreateLdapUser("amaltsev", "12345", "ldap://myserver:389")
	id := todo.Id

	// when
	todo, err := client.GetLdapUser(id)

	// then
	if err != nil {
		t.Error(err)
	}

	if todo.User != "amaltsev" && todo.Password != "12345" && todo.LdapServer != "ldap://myserver:389" {
		t.Error("returned Ldap user not right")
	}

	// cleanup
	_ = client.DeleteLdapUser(todo.Id)
}

func TestGetNotFoundLdapUser(t *testing.T) {

	// given
	client := client.LdapClient{Host: "http://localhost:8080"}
	id := int32(3)

	// when
	_, err := client.GetLdapUser(id)

	// then
	if err == nil {
		t.Error(err)
	}
}

func TestGetAllLdapUsers(t *testing.T) {

	// given
	client := client.LdapClient{Host: "http://localhost:8080"}
	client.CreateLdapUser("amaltsev", "12345", "ldap://myserver:389")
	client.CreateLdapUser("dpietropaolo", "12345", "ldap://myserver:389")

	// when
	todos, err := client.GetAllLdapUsers()

	// then
	if err != nil {
		t.Error(err)
	}

	if len(todos) != 2 {
		t.Errorf("wrong number of Ldap users: %d", len(todos))
	}
	if todos[0].User != "amaltsev" && todos[0].Password != "12345"  && todos[0].LdapServer != "ldap://myserver:389" {
		t.Error("returned Ldap user not right")
	}
	if todos[0].User != "dpietropaolo" && todos[0].Password != "12345"  && todos[0].LdapServer != "ldap://myserver:389" {
		t.Error("returned Ldap user not right")
	}

	// cleanup
	_ = client.DeleteLdapUser(todos[0].Id)
	_ = client.DeleteLdapUser(todos[1].Id)
}
