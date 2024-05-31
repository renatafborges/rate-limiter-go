package configs

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	RateLimitIP    int
	RateLimitToken int
	BlockDuration  time.Duration
	err            error
)

func LoadEnvConfig() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("unable to load .env configs", "error:", err)
		return
	}
}

func LoadRateLimitConfig() {

	RateLimitIP, err = strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	if err != nil {
		slog.Error("unable to parse RATE_LIMIT_IP", "error:", err)
		return
	}

	RateLimitToken, err = strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	if err != nil {
		slog.Error("unable to parse RATE_LIMIT_TOKEN", "error:", err)
		return
	}

	blockDurationSeconds, err := strconv.Atoi(os.Getenv("BLOCK_DURATION"))
	if err != nil {
		slog.Error("unable to parse BLOCK_DURATION", "error:", err)
		return
	}

	BlockDuration = time.Duration(blockDurationSeconds) * time.Second
}
