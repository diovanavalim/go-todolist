package routes

import (
	"net/http"
	"todolist/cmd/api/model"
)

var TaskRoutes = []model.Route{
	{
		URI:    "/task",
		Method: http.MethodPost,
		Func:   func(w http.ResponseWriter, r *http.Request) {},
	},
	{
		URI:    "/task",
		Method: http.MethodGet,
		Func:   func(w http.ResponseWriter, r *http.Request) {},
	},
	{
		URI:    "/task/{id}",
		Method: http.MethodGet,
		Func:   func(w http.ResponseWriter, r *http.Request) {},
	},
	{
		URI:    "/task/{id}",
		Method: http.MethodPut,
		Func:   func(w http.ResponseWriter, r *http.Request) {},
	},
	{
		URI:    "/task/{id}",
		Method: http.MethodDelete,
		Func:   func(w http.ResponseWriter, r *http.Request) {},
	},
}
