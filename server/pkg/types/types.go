package types

import (
	"net/http"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Users struct {
	Name string `json:"username"`
	Ping int    `json:"ping"`
}
