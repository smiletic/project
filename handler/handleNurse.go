package handler

import (
	"context"
	"masterRad/serverErr"
	"net/http"
	"strings"
)

func handleNurse(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/person") {
		switch r.Method {
		case http.MethodPost:
		case http.MethodPatch:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}
	if strings.HasPrefix(r.URL.Path, "/patient") {
		switch r.Method {
		case http.MethodPost:
		case http.MethodPatch:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}
	if strings.HasPrefix(r.URL.Path, "/examination") {
		switch r.Method {
		case http.MethodPost:
		case http.MethodPatch:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}

	return nil, serverErr.ErrInvalidAPICall
}
