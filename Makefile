APIFOX_URL := http://127.0.0.1:4523/export/openapi?projectId=3257993&version=3.0

APP_NAME ?= my-auth-service
IMAGE_TAG := tmp-$(shell git describe --always --abbrev=8 --dirty)

.PHONY: all
all:

.PHONY: apifox
apifox:
	curl '${APIFOX_URL}' -o ./api/openapi.json
	rm -rf internal/handler/api
	mkdir internal/handler/api internal/handler/temp
	docker run --rm -v ${PWD}:/src \
		openapitools/openapi-generator-cli generate \
		-t /src/internal/handler/api-template \
		-i /src/api/openapi.json \
		-g go-gin-server --package-name api \
		-o /src/internal/handler/temp
	mv internal/handler/temp/go/*.go internal/handler/api/
	gofmt -w internal/handler/api
	rm -rf internal/handler/temp

.PHONY: local-run
local-run:
	@go mod tidy
	mkdir -p dist && go build -o dist/$(APP_NAME) ./cmd
	if [ $$? -eq 0 ]; then ./dist/$(APP_NAME) serve; fi;

.PHONY: test
test:
	ginkgo -gcflags=all=-l -r -p --randomize-all --output-dir=test/report/ -covermode count --coverprofile=coverage.txt --junit-report=report.xml --coverpkg=./internal/repository