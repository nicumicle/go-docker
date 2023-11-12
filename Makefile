build:
	go build -o out/go-docker main.go

run:
	go run main.go

lint:
	golangci-lint run

test:
	go test -race $$(go list ./... | grep -v /*_test.go/ | grep -v /vendor/) -cover
