MIGRATE_TAG = v4.14.1
DATABASE_URL ?= "postgres://dusanmalusev:password@localhost:5432/dusanmalusev?sslmode=disable"
RACE ?= 0


build:

test:
ifeq ($(RACE), 1)
	@go test ./... -race
else
	@go test ./...
endif

config.yml:
	@cp config.example.yml config.yml

setup: config.yml
	@docker-compose up -d

migrate: migrate-install
	@${GOPATH}/bin/migrate -source file://${shell pwd}/database/migrations -database ${DATABASE_URL} up

migrate-down: migrate-install
	@${GOPATH}/bin/migrate -source file://${shell pwd}/database/migrations -database ${DATABASE_URL} down

migration-create: migrate-install
	@${GOPATH}/bin/migrate -database ${DATABASE_URL} create -dir ./database/migrations -seq -ext sql ${M_NAME}


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

clean:
	@rm -rf ./bin

.PHONY: migrate-install migrate build clean test
