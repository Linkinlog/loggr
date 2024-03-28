bin := loggr
main := cmd/loggr/main.go
gcp_url := us-east5-docker.pkg.dev/loggr-418603/loggr/$(bin)

build:
	@go build -o $(bin) $(main)
build.docker:
	@docker build -t $(bin) -f ./build/Dockerfile .
build.tag:
	@docker tag $(bin) $(gcp_url)
build.push:
	@docker push $(gcp_url)
release:
	@make lint || exit 1
	@make test || exit 1
	@make build.docker
	@make build.tag
	@make build.push

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
