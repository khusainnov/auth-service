FROM golang:1.18.3

RUN mkdir /auth-service
WORKDIR /auth-service

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY /cmd/main.go cmd/
COPY /config/.env config/
#COPY /doc/weather.swagger.json doc/
COPY /driver/*.go driver/
COPY /gen/pb/*.go gen/pb/
COPY /internal/entity/*.go internal/entity/
COPY /pkg/endpoint/*.go pkg/endpoint/
COPY /pkg/repository/*.go pkg/repository/
COPY /pkg/service/*.go pkg/service/
#COPY /proto/google/api/*.proto proto/google/api/
#COPY /proto/protoc-gen-openapiv2/options/*.proto proto/protoc-gen-openapiv2/options/
COPY /proto/ proto/
COPY Makefile .
COPY server.go .

RUN go build -o auth-service cmd/main.go

EXPOSE 80

CMD ["./auth-service"]