package apiclient

import (
	"time"
)

type OneCloudAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Tenant   string `json:"tenant"`
}

type OneCloudAuthResponse struct {
	Expires time.Time `json:"expires"`
	ID      string    `json:"id"`
	Tenant  string    `json:"tenant"`
}

type APIError struct {
	Errors []struct {
		Code          int    `json:"code"`
		Message       string `json:"message"`
		SystemMessage string `json:"systemMessage"`
	} `json:"errors"`
}
