package handler

import (
	"context"
	"masterRad/core"
	"masterRad/serverErr"
	"net/http"
	"strings"
)

func handleDoctor(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/test") {
		r.URL.Path = r.URL.Path[5:]
		switch r.Method {
		case http.MethodPost:
			return nil, core.CreateTest(ctx, r)
		case http.MethodGet:
			if strings.HasPrefix(r.URL.Path, "/") {
				return core.GetTest(ctx, r.URL.Path[1:])
			}
			return core.GetTests(ctx)
		case http.MethodDelete:
			return nil, core.RemoveTest(ctx, r.URL.Path[1:])
		}
	}
	if strings.HasPrefix(r.URL.Path, "/filled") {
		switch r.Method {
		case http.MethodPost:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}
	return nil, serverErr.ErrInvalidAPICall
}
