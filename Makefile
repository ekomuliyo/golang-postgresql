run: bin/golang-postgresql
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/golang-postgresql: main.go
	go build -o bin/golang-postgresql main.go