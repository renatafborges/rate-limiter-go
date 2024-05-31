package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime"

	"golang.org/x/sync/errgroup"
)

var (
	Host string
	URL  string
)

func init() {
	Host = os.Getenv("URL_RATE")
	if Host == "" {
		Host = "localhost"
	}

	URL = "http://" + Host + ":8080"
}

func main() {

	eg := errgroup.Group{}

	eg.Go(func() error {
		if !TestLimitWithTokenAndIp() {
			return errors.New("failed to test limitWithTokenAndIp")
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		_, file, line, _ := runtime.Caller(0)
		slog.Error("failed to execute automated tests", "caller", fmt.Sprintf("%s :%d", file, line), "error", err)
		os.Exit(1)
		return
	}

	fmt.Println("Automated tests executed with success")
	os.Exit(0)
}

func TestLimitWithTokenAndIp() bool {

	client := &http.Client{}

	rateLimitToken := 2

	token := "123456789"

	makeRequest := func(client *http.Client, token string) *http.Response {
		req, err := http.NewRequest(http.MethodGet, URL, nil)
		if err != nil {
			slog.Error("failed to make new request", err)
		}
		if token != "" {
			req.Header.Set("API_KEY", token)
		}

		resp, err := client.Do(req)
		if err != nil {
			slog.Error("Error doing req:", err)
			panic(err)
		}

		return resp
	}

	var testok bool

	for i := 0; i <= rateLimitToken; i++ {
		resp := makeRequest(client, token)
		if i < rateLimitToken {
			if resp.StatusCode == http.StatusOK {
				testok = false
				continue
			}
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			testok = true
			continue
		}

	}

	return testok
}
