package handler

import (
	"context"
	"masterRad/serverErr"
	"net/http"
	"strings"
)

func handleResearcher(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/filled") {
		switch r.Method {
		case http.MethodGet:
		}
	}
	return nil, serverErr.ErrInvalidAPICall
}
