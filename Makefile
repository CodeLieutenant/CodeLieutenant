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

.PHONY: install-migrate-cli
install-migrate-cli:/
ifneq ($(findstring migrate,$(shell ls ${GOPATH}/bin)),migrate)
	cd ${GOPATH} && \
	rm -rf ${GOPATH}/src/github.com/golang-migrate/migrate && \
	go get -u -d github.com/golang-migrate/migrate/cmd/migrate && \
	cd ${GOPATH}/src/github.com/golang-migrate/migrate && \
	git checkout ${MIGRATE_TAG} && \
	cd cmd/migrate && \
	go build -tags 'postgres github' -ldflags="-X main.Version=${MIGRATE_TAG}" -o ${GOPATH}/bin/migrate ${GOPATH}/src/github.com/golang-migrate/migrate/cmd/migrate
endif
