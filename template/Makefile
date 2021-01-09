.PHONY: all
all: build-prod copy

.PHONY: dev
dev: install-fiber
	@fiber dev

.PHONY: build-prod
build-prod:
	@npm install
	@npm run build:prod
	@cp -r ./img/ ./dist/img/
	@cp -r ./views ./views_copy
	@node assets.js

.PHONY: copy
copy:
	@rm -rf ../server/views
	@rm -rf ../server/public
	@cp -r ./views_copy ../server/views
	@cp -r ./dist ../server/public

install-fiber:
ifeq ($(findstring fiber,$(shell ls ${GOPATH}/bin)),fiber)
	@go get -u github.com/gofiber/cli/fiber
endif

clean:
	@rm -rf ./views_copy
	@rm -rf ./dist

include ../Makefile
