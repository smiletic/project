package handler

import (
	"context"
	"net/http"
	"projekat/core"
	"projekat/serverErr"
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
			if strings.HasPrefix(r.URL.Path, "/") {
				return nil, core.RemoveTest(ctx, r.URL.Path[1:])
			}
			return nil, serverErr.ErrBadRequest
		}
	}
	if strings.HasPrefix(r.URL.Path, "/filled") {
		r.URL.Path = r.URL.Path[7:]
		switch r.Method {
		case http.MethodPost:
			return nil, core.CreateFilledTest(ctx, r.Body)
		case http.MethodGet:
			if strings.HasPrefix(r.URL.Path, "/") {
				return core.GetFilledTest(ctx, r.URL.Path[1:])
			}
			return core.GetFilledTests(ctx)
		case http.MethodDelete:
			if strings.HasPrefix(r.URL.Path, "/") {
				return nil, core.RemoveFilledTest(ctx, r.URL.Path[1:])
			}
			return nil, serverErr.ErrBadRequest
		}
	}
	return nil, serverErr.ErrInvalidAPICall
}
