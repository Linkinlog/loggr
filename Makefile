bin := loggr
main := cmd/loggr/main.go

build:
	@go build -o $(bin) $(main)

gen:
	@go generate ./...

test:
	@go test -v ./...
test.cov:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
test.dead:
	@deadcode -test ./...

clean:
	@rm -f $(bin)

lint:
	@make gen
	@templ fmt .
	@gofumpt -d -w .
	@golangci-lint run

dev:
	@make gen
	@make build
	@./$(bin)

done:
	@make clean || exit 1
	@make test  || exit 1
	@make lint  || exit 1
	@make build || exit 1

.PHONY: build gen dev test test.cov test.dead lint clean
