.PHONY: clean test security build run

APP_NAME = web
BUILD_DIR = $(PWD)/build
DATABASE_CONFIG_FILE = $(PWD)/infrastructure/database/database.yml
MIGRATIONS_FOLDER = $(PWD)/infrastructure/database/migrations
MODELS_FOLDER = $(PWD)/domain

configure:
	asdf install
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.9.1
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.3
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/go-critic/go-critic/cmd/gocritic@latest
	golangci-lint linters --enable gocritic
	asdf reshim golang

dependencies:
	go mod download
	go mod tidy

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

linter:
	golangci-lint --enable gocritic run ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

# https://go.dev/blog/cover
# https://blog.seriesci.com/how-to-measure-code-coverage-in-go/
# https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d
# https://blog.dharnitski.com/2019/04/29/skip-tests-in-go/
cover:
	go test -coverprofile=cover.out -cover ./...
	go tool cover -func cover.out -o cover_percentages.out
	go tool cover -html=cover.out -o coverage.html
	TEST_COVERAGE_CHECK=true go test ./pkg/test_utils -run TestCheck

swag.init:
	swag init -g ./application/web/web.go  -o ./application/web/infrastructure/docs

web.run:
	buffalo dev

docker.up:
	mkdir -p .docker/redisinsight
	sudo chown -R 1001 .docker/redisinsight
	docker-compose up -d

docker.down:
	docker-compose down

docker.build:
	docker build -t fiber .

run: docker.up swag.init migrate.up web.run.air

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/apiserver/main.go

build.run: swag.init build
	$(BUILD_DIR)/$(APP_NAME)

soda.create:
	soda -c $(DATABASE_CONFIG_FILE) -p $(MIGRATIONS_FOLDER) create -a

soda.drop:
	soda -c $(DATABASE_CONFIG_FILE) -p $(MIGRATIONS_FOLDER) drop -a

soda.migrate.status:
	soda -c $(DATABASE_CONFIG_FILE) -p $(MIGRATIONS_FOLDER) migrate status

soda.migrate.up:
	soda -c $(DATABASE_CONFIG_FILE) -p $(MIGRATIONS_FOLDER) migrate up

soda.migrate.down:
	soda -c $(DATABASE_CONFIG_FILE) -p $(MIGRATIONS_FOLDER) migrate down

soda.migrate.generate:
	soda -c $(DATABASE_CONFIG_FILE) -p $(MIGRATIONS_FOLDER) generate fizz $@

soda.generate.model:
	soda -c $(DATABASE_CONFIG_FILE) -p $(MIGRATIONS_FOLDER)/user --models-path $(MODELS_FOLDER)/user generate model user
