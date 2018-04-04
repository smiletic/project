package handler

import (
	"context"
	"masterRad/serverErr"
	"net/http"
	"strings"
)

func handleDoctor(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/test") {
		switch r.Method {
		case http.MethodPost:
		case http.MethodPatch:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}
	if strings.HasPrefix(r.URL.Path, "/filled") {
		switch r.Method {
		case http.MethodPost:
		case http.MethodPatch:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}
	return nil, serverErr.ErrInvalidAPICall
}
