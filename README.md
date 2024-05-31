# Rate Limiter

This project implements a rate limiter in Go, which limits the number of requests per second based on IP address or access token. Uses Redis to store information.

## Running the project

1. Clone the repository.

2. Build and run the project using Docker Compose to build the application and Run Automated Tests:
```sh
docker-compose up --build
```
You should see
![image](https://github.com/renatafborges/rate-limiter-go/assets/63026376/eb92b024-0052-467f-bcce-e4fdd46272fa)


4. In cmd/app/.env
Its possible to configure values when running locally:
```sh
REDIS_ADDR=localhost:6379 (Redis local port)
RATE_LIMIT_IP=1 (limit by ip address 1 request/second)
RATE_LIMIT_TOKEN=2 (limit by token 2 request/second)
BLOCK_DURATION=5 # em segundos (time for blocking when the limit is exceed)
```

4. In api
Its possible to make requests locally, just open the file and after running redis and cmd/app/main.go click in Send Request:

```sh
docker-compose up redis
go run main.go
You should see: Starting server on: :8080
```

![image](https://github.com/renatafborges/rate-limiter-go/assets/63026376/74dbe26c-8989-4254-9ac4-5c4fb438d88a)

5. When the request is succeed:

![image](https://github.com/renatafborges/rate-limiter-go/assets/63026376/0faa03f5-e03b-49fc-8d63-3acf7b76c5a4)

6. When the request failed:

![image](https://github.com/renatafborges/rate-limiter-go/assets/63026376/31000d5b-ee84-447d-b67e-89d17c3ce94a)
