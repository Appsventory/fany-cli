package utils

import (
	"net/http"
	"time"
)

func HasInternet() bool {
	c := &http.Client{Timeout: 2 * time.Second}
	_, err := c.Get("https://github.com")
	return err == nil
}
