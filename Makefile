APIFOX_URL := http://127.0.0.1:4523/export/openapi?projectId=3257993&version=3.0

APP_NAME ?= my-auth-service
IMAGE_NAME ?= $(APP_NAME)
IMAGE_TAG := tmp-$(shell git describe --always --abbrev=8 --dirty)

.PHONY: all
all:

.PHONY: test
test:
	go install github.com/onsi/ginkgo/v2/ginkgo@latest
	ginkgo -gcflags=all=-l -r -p --randomize-all --output-dir=test/report/ -covermode count --coverprofile=coverage.txt --junit-report=report.xml --coverpkg=./internal/repository,./internal/handler

.PHONY: local-run
local-run: test
	@go mod tidy
	mkdir -p dist && go build -o dist/$(APP_NAME) ./cmd
	if [ $$? -eq 0 ]; then ./dist/$(APP_NAME) serve; fi;

.PHONY: go-build
go-build: test
	@go mod tidy
	mkdir -p dist && CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o dist/$(APP_NAME) ./cmd

.PHONY: docker-build
docker-build:
	docker build -f build/package/formal-deploy/Dockerfile -t $(IMAGE_NAME):$(IMAGE_TAG)  --build-arg APP_NAME=$(APP_NAME) .

.PHONY: docker-run
docker-run: docker-build
	docker run --rm -it -p 4201:4201 $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: apifox-gen
apifox-gen:
	curl '${APIFOX_URL}' -o ./api/openapi.json
	rm -rf internal/handler/gen internal/handler/temp
	mkdir -p internal/handler/gen internal/handler/temp
	docker run --rm -v ${PWD}:/src \
		openapitools/openapi-generator-cli generate \
		-t /src/pkg/api-template \
		-i /src/api/openapi.json \
		-g go-gin-server --package-name gen \
		-o /src/internal/handler/temp
	mv internal/handler/temp/go/*.go internal/handler/gen/
	mv internal/handler/temp/api/openapi.yaml api/
	gofmt -w internal/handler/gen
	rm -rf internal/handler/temp
