package handler

import (
	"context"
	"net/http"
	"strings"
)

func handleAdmin(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/person") {
		switch r.Method {
		case http.MethodPost:

			return
		case http.MethodPatch:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}
	if strings.HasPrefix(r.URL.Path, "/employee") {

	}
	if strings.HasPrefix(r.URL.Path, "/user") {

	}
	//Person CRUD
	//Employee CRUD
	//User CRUD
	return
}
