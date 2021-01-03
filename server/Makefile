MIGRATE_TAG = v4.14.1
DATABASE_URL ?= "postgres://dusanmalusev:password@localhost:5432/dusanmalusev?sslmode=disable"
RACE ?= 0
ENVIRONMENT ?= development
VERSION ?= dev


.PHONY: build
build: copy-files
ifeq ($(ENVIRONMENT),production)
	@CGO_ENABLED=0 go build -ldflags="-s -w -X 'main.Version=${VERSION}'" -o ../bin/dusanmalusev main.go
else ifeq ($(ENVIRONMENT),development)
	@go build -o ../bin/dusanmalusev main.go
else
	@echo "Target ${ENVIRONMENT} is not supported"
endif

.PHONY: copy-files
copy-files:
	@cp -r ./database/migrations ../bin/migrations
	@cp -r ./public ../bin/public
	@cp -r ./views ../bin/views

.PHONY: test
test:
ifeq ($(RACE), 1)
	@go test ./... -race
else
	@go test ./...
endif

config.yml:
	@cp config.example.yml config.yml

.PHONY: setup
setup: config.yml
	@docker-compose up -d

.PHONY: migrate
migrate: migrate-install
	@${GOPATH}/bin/migrate -source file://${shell pwd}/database/migrations -database ${DATABASE_URL} up

.PHONY: migrate-down
migrate-down: migrate-install
	@${GOPATH}/bin/migrate -source file://${shell pwd}/database/migrations -database ${DATABASE_URL} down

.PHONY: migrate-create
migration-create: migrate-install
	@${GOPATH}/bin/migrate -database ${DATABASE_URL} create -dir ./database/migrations -seq -ext sql ${M_NAME}

.PHONY: migrate-install
migrate-install:
ifneq ($(findstring migrate,$(shell ls ${GOPATH}/bin)),migrate)
	@cd ${GOPATH} && \
	rm -rf ${GOPATH}/src/github.com/golang-migrate/migrate && \
	go get -u -d github.com/golang-migrate/migrate/cmd/migrate && \
	cd ${GOPATH}/src/github.com/golang-migrate/migrate && \
	git checkout ${MIGRATE_TAG} && \
	cd cmd/migrate && \
	go build -tags 'postgres github' -ldflags="-X main.Version=${MIGRATE_TAG}" -o ${GOPATH}/bin/migrate ${GOPATH}/src/github.com/golang-migrate/migrate/cmd/migrate
endif

.PHONY: clean
clean:
	@rm -rf ./bin

