go-run:
	go run cmd/example.go

go-build:
	go build -o build/example *.go

go-run-build:
	./build/example

go-download:
	export GO111MODULE=on
	go mod download