FROM golang:1.19 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./secretConfig.yaml ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-weather-api
FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine:3.18 AS build-release-stage
WORKDIR /
COPY --from=build-stage /go-weather-api /go-weather-api
EXPOSE 8089
USER root:root
ENTRYPOINT ["/go-weather-api"]
