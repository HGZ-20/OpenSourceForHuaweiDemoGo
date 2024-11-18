FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
COPY user/etc /app/etc
COPY openGauss-connector-go-pq /app/openGauss-connector-go-pq
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/user user/user.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/openGauss-connector-go-pq /app/openGauss-connector-go-pq

CMD ["./user", "-f", "etc/user-api.yaml"]
