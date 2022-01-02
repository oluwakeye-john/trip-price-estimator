dev:
	nodemon --exec go run server.go --signal SIGTERM

generate:
	go run github.com/99designs/gqlgen