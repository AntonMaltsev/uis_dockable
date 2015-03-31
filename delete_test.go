package main

import (
	"fmt"
	"github.com/AntonMaltsev/uis_dockable/client"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestDeleteLdap(t *testing.T) {

	// given
	client := client.LdapClient{Host: "http://localhost:8080"}
	todo, _ := client.CreateLdapUser("amaltsev", "12345", "ldap://myserver:389")
	id := todo.Id

	// when
	err := client.DeleteLdapUser(id)

	// then
	if err != nil {
		t.Error(err)
	}

	_, err = client.GetLdapUser(id)
	if err == nil {
		t.Error(err)
	}
}

func TestDeleteNotFoundLdap(t *testing.T) {

	// given
	client := client.LdapClient{Host: "http://localhost:8080"}
	id := int32(3)
	// when
	err := client.DeleteLdapUser(id)

	// then
	if err == nil {
		t.Error(err)
	}

}
