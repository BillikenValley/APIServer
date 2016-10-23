package main

import (
	"net/http"

	"github.com/BillikenValley/APIServer/model"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ShelterIndex",
		"GET",
		"/api/v1/shelter",
		model.ShelterIndex,
	},
	Route{
		"ShelterShow",
		"GET",
		"/api/v1/shelter/{ShelterID}",
		model.ShelterShow,
	},
	Route{
		"ShelterUpload",
		"POST",
		"/api/v1/shelter",
		model.ShelterUpload,
	},
}
