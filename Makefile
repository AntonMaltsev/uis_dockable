default: build

clean:
	rm -f cmd/server/server
	rm -f cmd/ldap/ldap

build:
	cd cmd/server; \
	go build
	cd cmd/ldap; \
	go build

deps:
	cd cmd/server; \
	go get

migrate:
	./cmd/server/server --config config.yaml migratedb

test:
	./cmd/server/server --config config.yaml server & \
	pid=$$!; \
	go test; \
	kill $$pid
