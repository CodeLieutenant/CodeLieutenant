ENVIRONMENT ?= development

.PHONY: all
all: build copy

.PHONY: dev
dev: install-fiber-cli
	fiber dev

.PHONY: template
template:
	go run main.go -prerender=true

.PHONY: build
build: clean template
	npm install
	mkdir -p dist/img
	mkdir -p views_copy
	cp -r img/* dist/img/
	cp -r views/* ./views_copy/
ifeq ($(ENVIRONMENT),production)
	npm run production
	node assets.js
else ifeq ($(ENVIRONMENT),development)
	npm run development
else
	echo "Target ${ENVIRONMENT} is not supported"
endif

.PHONY: copy
copy:
	rm -rf ../server/views
	rm -rf ../server/public
	mkdir -p ../server/public
	mkdir -p ../server/views
	cp -r views_copy/* ../server/views/
	cp -r dist/* ../server/public

clean:
	rm -rf ./views_copy
	rm -rf ./dist
	rm -rf ./rendered

include ../Makefile
