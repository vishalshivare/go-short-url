run-local:
	go mod tidy
	go run main.go

build:
	go mod tidy
	go build -o go-short-url

test:
	go test ./... -v -coverprofile=coverage.out 

