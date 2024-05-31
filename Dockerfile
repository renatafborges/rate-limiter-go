FROM golang:1.22.3 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd/app -o rate_limiter

FROM scratch
WORKDIR /app
COPY --from=build /app/cmd/app/.env /app/cmd/app/rate_limiter ./
ENTRYPOINT ["./rate_limiter"]