## UIS Docker REST Microservice for Gin

This microservices project is using the [Gin](http://gin-gonic.github.io/gin/) web framework.

See the [blog post](http://txt.fliglio.com/2014/07/restful-microservices-in-go-with-gin/) for a walk through.

### Hacking

#### Build Service
	
	cd cmd/server; 
	go build
	cd cmd/ldap; 
	go build

#### Build the Database

##Create the database by hand or	
	
	mysql -u root -p -e 'Create Database Todo;'

	./cmd/server/server --config config.yaml migratedb


#### Run the Service

	./cmd/server/server --config config.yaml server

#### Testing
The tests leverage a running instance of the server. 

	go test
