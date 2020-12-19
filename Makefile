run: bin/golang-postgresql

bin/golang-postgresql: main.go
	go build -o bin/golang-postgresql main.go