package client

import (
	"net/http"
	"time"
)

func MyHttpClient() *http.Client {
	client := &http.Client{Timeout: 30 * time.Second}
	return client
}
