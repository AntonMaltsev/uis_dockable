package main

import (
	"fmt"
	"github.com/AntonMaltsev/uis_dockable/client"
	"github.com/codegangsta/cli"
	"log"
	"os"
	// "strconv"
)

func main() {

	app := cli.NewApp()
	app.Name = "UIS cli"
	app.Usage = "cli to work with the UIS microservice"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"host", "http://localhost:8080", "Todo service host", "APP_HOST"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "(title description) Create an ldap user",
			Action: func(c *cli.Context) {
				user_id := c.Args().Get(0)				//title
				user_password := c.Args().Get(1)		//description
				ldap_host := c.Args().Get(2)

				host := c.GlobalString("host")

				client := client.LdapClient{Host: host}

				todo, err := client.CreateLdapUser(user_id, user_password, ldap_host)
				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Printf("%+v\n", todo)
			},
		},
		{
			Name:  "ls",
			Usage: "list all ldap users",
			Action: func(c *cli.Context) {

				host := c.GlobalString("host")

				client := client.LdapClient{Host: host}

				todos, err := client.GetAllLdapUsers()
				if err != nil {
					log.Fatal(err)
					return
				}
				for _, todo := range todos {
					fmt.Printf("%+v\n", todo)
				}
			},
		},
	}
	app.Run(os.Args)

}
