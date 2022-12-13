PROJECT = $(shell basename `pwd`)
APP_ENV="local"
APP_TAG=""
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT_ID=$(shell git rev-parse --short HEAD)
APP_BUILD_TIME=$(shell date +'%Y%m%d%H%M')
APP_IMAGE=harbor.ickey.com.cn/dev/${PROJECT}:${GIT_BRANCH}-${GIT_COMMIT_ID}-${APP_BUILD_TIME}

all: clean build

.PHONY: build
build:
	go build -ldflags '-X main.branch=${GIT_BRANCH} -X main.commit=${GIT_COMMIT_ID}' -v -o bin/${PROJECT} *.go
	mkdir build
	cp bin/${PROJECT} build/
	cp -r etc build/

run:
	go run *.go -f ./etc/app_local.yaml -t "${tag}"

create:
	ickey-goctl -t rpc -o create
gen:
	ickey-goctl -t rpc -o gen
gen2:
	goctl rpc protoc database-service.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=go_zero --home=/home/infra
	rm -fv ./etc/databaseservice.yaml
sql:
	ickey-goctl -o sql
test:
	go test -timeout 1m --race ./...
clean:
	go clean -i ./...
	rm -rf bin
	rm -rf build
swagger:
	protoc --swagger_out=. ./*.proto
build-image:
	docker build -t ${APP_IMAGE} -f Dockerfile --build-arg APP=${PROJECT} --build-arg APP_ENV=${APP_ENV} --build-arg APP_TAG=${APP_TAG} --build-arg GIT_BRANCH=${GIT_BRANCH} --build-arg GIT_COMMIT_ID=${GIT_COMMIT_ID} .
	docker push ${APP_IMAGE}