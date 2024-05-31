FROM golang:1.22.3 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rate-limiter-tests

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/exec_automated_test.sh /app/rate-limiter-tests ./
ENTRYPOINT [ "sh", "./exec_automated_test.sh" ]