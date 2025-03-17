package utils

import (
	"net/http"
	"time"
)

var (
	Client *http.Client
)

func init() {
	Client = &http.Client{
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			MaxConnsPerHost:     120,
			MaxIdleConnsPerHost: 20,
			MaxIdleConns:        80,
		},
		Timeout: time.Second * 30,
	}
}
