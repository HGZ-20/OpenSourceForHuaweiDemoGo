# proto
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero
# api
goctl api go -api user/api/user.api -dir user/api -style gozero
# swagger
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user/api/user.api -dir .
# 指定host
goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.2 -basepath /api -schemes https,wss" -api user/api/user.api -dir .
# swagger-ui
 docker run --rm -p 8083:8080 -e SWAGGER_JSON=/apps/user.json -v $PWD:/foo swaggerapi/swagger-ui