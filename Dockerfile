FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
COPY apps/user/api/etc /app/api/etc
COPY apps/user/rpc/etc /app/rpc/etc
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/user/rpc apps/user/rpc/user.go
RUN go build -ldflags="-s -w" -o /app/user/api apps/user/api/user.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user /app/user/rpc
COPY --from=builder /app/etc /app/rpc/etc

CMD ["./user", "-f", "etc/user-rpc.yaml"]

WORKDIR /app
COPY --from=builder /app/user /app/user/api
COPY --from=builder /app/etc /app/api/etc

CMD ["./user", "-f", "etc/user-api.yaml"]
