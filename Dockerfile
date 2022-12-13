# ---- 编译环境 ----
FROM harbor.ickey.com.cn/common/golang:1.18-debian10 AS builder

ARG APP_ENV=test \
    APP=undefine \
    GIT_BRANCH= \
    GIT_COMMIT_ID=

ENV APP_ENV=$APP_ENV \
    APP=$APP \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on

# 下载依赖
WORKDIR /app_build
COPY go.* ./

RUN go env && go mod download

# 编译
COPY . .
RUN go build  -tags 'osusergo,netgo' \
    -ldflags "-X main.branch=$GIT_BRANCH  -X main.commit=$GIT_COMMIT_ID" \
    -v -o bin/${APP} *.go \
    && cp -rf etc bin/etc \
    && chown 999.999 -R bin

#
# ---- 运行环境 ----

FROM harbor.ickey.com.cn/common/runtime:golang-debian10 as running

ARG APP_ENV=test \
    APP=undefine \
    APP_TAG=
ENV APP_ENV=$APP_ENV \
    APP_TAG=$APP_TAG \
    APP=$APP

WORKDIR /app

COPY --from=builder /app_build/bin /app/

CMD ["bash", "-c", "exec /app/${APP} -t \"${APP_TAG}\" -f /app/etc/app_${APP_ENV}.yaml"]