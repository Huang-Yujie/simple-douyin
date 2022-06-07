FROM golang:alpine as builder

ENV GO111MODULE=on \
	GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o api ./cmd/api

FROM alpine as prod

WORKDIR /app

COPY --from=builder /app/config.yaml .
COPY --from=builder /app/api .

# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["./api"]