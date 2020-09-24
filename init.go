package main

import (
	"time"
	"net/http"
)

type Config struct {
	KAFKA_CONNECT_REST   string
	FORMAT               string
}

 
var (
	config      Config
	httpClient  *http.Client
)

func init() {
		httpClient = &http.Client{Timeout: time.Duration(30 * time.Second)}
}