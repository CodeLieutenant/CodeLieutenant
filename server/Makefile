MIGRATE_TAG = v4.14.1
RACE ?= 0
ENVIRONMENT ?= development
VERSION ?= dev
GOPATH ?= ${HOME}/go
DOCKER ?= 0

ifeq ($(DOCKER),0)
DATABASE_URL ?= "postgres://postgres:postgres@localhost:5432/dusanmalusev?sslmode=disable"
else
DATABASE_URL ?= "postgres://postgres:postgres@postgres:5432/dusanmalusev?sslmode=disable"
endif

CC = gcc
CXX = g++

.PHONY: build
build: copy-files
ifeq ($(ENVIRONMENT),production)
	CGO_ENABLED=0 go build -tags "binary_log" -ldflags="-s -w -X 'main.Version=${VERSION}'" -o ../bin/dusanmalusev main.go
else ifeq ($(ENVIRONMENT),development)
	go build -o ./bin/dusanmalusev main.go
else
	echo "Target ${ENVIRONMENT} is not supported"
endif

.PHONY: copy-files
copy-files: config.yml
	mkdir -p ./bin/migrations
	mkdir -p ./bin/public
ifeq ($(DOCKER),1)
	cp config.docker.yml ./bin/config.yml
else
	cp config.yml ./bin/config.yml
endif
	cp -r ./database/migrations ./bin
	cp -r ./public ./bin/

.PHONY: test
test:
ifeq ($(RACE), 1)
	DB_CONN=${DATABASE_URL} go test ./... -race -covermode=atomic -coverprofile=coverage.txt -timeout 5m
else
	DB_CONN=${DATABASE_URL} go test ./... -covermode=atomic -coverprofile=coverage.txt -timeout 1m
endif

config.yml:
	cp config.example.yml config.yml

.PHONY: setup
setup: config.yml
	docker-compose up -d

.PHONY: migrate
migrate:
	migrate -source file://$(shell pwd)/database/migrations -database ${DATABASE_URL} up

.PHONY: migrate-down
migrate-down:
	migrate -source file://$(shell pwd)/database/migrations -database ${DATABASE_URL} down

.PHONY: migrate-create
migration-create:
	migrate -database $(DATABASE_URL) create -dir ./database/migrations -seq -ext sql $(M_NAME)

.PHONY: tidy
tidy:
	rm -f go.sum
	go mod tidy

.PHONY: dev
dev: install-fiber-cli config.yml
	docker-compose up -d
	make migrate
	@make run

.PHONY: run
run:
	@go run main.go server -d true

.PHONY: generate
generate:
	@go run main.go generate

.PHONY: clean
clean:
	rm -rf ./bin

GOPATH ?= ${HOME}/go
MIGRATE_TAG ?= v4.14.1
.PHONY: git-setup
git-setup:
	git config user.name GitHub
	git config user.email noreply@github.com
	git remote set-url origin https://x-access-token:${GITHUB_TOKEN}@github.com/malusev998/dusanmalusev.git


.PHONY: commit
commit:
	git add .
ifneq ($(shell git status --porcelain),)
	git commit --author "github-actions[bot] <github-actions[bot]@users.noreply.github.com>" --message "${MESSAGE}"
	git push
endif

.PHONY: install-fiber-cli
install-fiber-cli:
ifneq ($(findstring fiber,$(shell ls ${GOPATH}/bin)),fiber)
	cd ${GOPATH} && go get -u github.com/gofiber/cli/fiber
endif

