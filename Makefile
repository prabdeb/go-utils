all:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
json:
	go test -coverprofile=coverage.out -json ./... > test-result.json
	go tool cover -func=coverage.out