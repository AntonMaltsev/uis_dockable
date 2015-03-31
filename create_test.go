package main

import (
	"fmt"
	"github.com/AntonMaltsev/uis_dockable/client"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestCreateLdap(t *testing.T) {

	// given
	client := client.LdapClient{Host: "http://localhost:8080"}

	// when
	todo, err := client.CreateLdapUser("amaltsev", "12345", "ldap://myserver:389")

	//then
	if err != nil {
		t.Error(err)
	}

	if todo.User != "amaltsev" && todo.Password != "12345" && todo.LdapServer != "ldap://myserver:389" {
		t.Error("returned Ldap user is not right")
	}

	// cleanup
	_ = client.DeleteLdapUser(todo.Id)
}
