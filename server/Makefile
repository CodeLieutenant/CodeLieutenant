MIGRATE_TAG = v4.14.1
DATABASE_URL ?= "postgres://postgres:postgres@localhost:5432/dusanmalusev?sslmode=disable"
RACE ?= 0
ENVIRONMENT ?= development
VERSION ?= dev
GOPATH ?= ${HOME}/go

.PHONY: build
build: copy-files
ifeq ($(ENVIRONMENT),production)
	CGO_ENABLED=0 go build -tags "binary_log" -ldflags="-s -w -X 'main.Version=${VERSION}'" -o ../bin/dusanmalusev main.go
else ifeq ($(ENVIRONMENT),development)
	go build -o ../bin/dusanmalusev main.go
else
	echo "Target ${ENVIRONMENT} is not supported"
endif

.PHONY: copy-files
copy-files:
ifeq ($(ENVIRONMENT),production)
	mkdir -p ../bin/migrations
	mkdir -p ../bin/public
	mkdir -p ../bin/view
	cp -r ./database/migrations/ ../bin/migrations/
	cp -r ./public ../bin/public
	cp -r ./views ../bin/views
else ifeq ($(ENVIRONMENT),development)
	mkdir -p ./bin/migrations
	mkdir -p ./bin/public
	mkdir -p ./bin/view
	cp -r ./database/migrations ./bin/migrations
	cp -r ./public ./bin/public
	cp -r ./views ./bin/views
else
	echo "Target ${ENVIRONMENT} is not supported"
endif

.PHONY: test
test:
ifeq ($(RACE), 1)
	DB_CONN=${DATABASE_URL} go test ./... -race -covermode=atomic -coverprofile=coverage.txt
else
	DB_CONN=${DATABASE_URL} go test ./... -covermode=atomic -coverprofile=coverage.txt
endif

config.yml:
	cp config.example.yml config.yml

.PHONY: setup
setup: config.yml
	docker-compose up -d

.PHONY: migrate
migrate: install-migrate-cli
	${GOPATH}/bin/migrate -source file://${shell pwd}/database/migrations -database ${DATABASE_URL} up

.PHONY: migrate-down
migrate-down: install-migrate-cli
	${GOPATH}/bin/migrate -source file://${shell pwd}/database/migrations -database ${DATABASE_URL} down

.PHONY: migrate-create
migration-create: install-migrate-cli
	${GOPATH}/bin/migrate -database ${DATABASE_URL} create -dir ./database/migrations -seq -ext sql ${M_NAME}

.PHONY: tidy
tidy:
	rm -f go.sum
	go mod tidy

.PHONY: dev
dev: install-fiber-cli install-migrate-cli config.yml
	docker-compose up -d
	make migrate
	fiber dev

.PHONY: clean
clean:
ifeq ($(ENVIRONMENT),production)
	rm -rf ../bin
else ifeq ($(ENVIRONMENT),development)
	rm -rf ./bin
else
	echo "Target ${ENVIRONMENT} is not supported"
endif

include ../Makefile
