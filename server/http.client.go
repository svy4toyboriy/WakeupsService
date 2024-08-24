package server

import (
	"net/http"
	"time"
)

var HTTPClient = &http.Client{
	Timeout: time.Second * 10,
}
