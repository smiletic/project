package handler

import (
	"context"
	"net/http"
	"projekat/core"
	"projekat/serverErr"
	"strings"
)

func handleResearcher(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/filled") {
		switch r.Method {
		case http.MethodGet:
			return core.GetFilledTest(ctx, r.URL.Path[1:])
		}
	}
	return nil, serverErr.ErrInvalidAPICall
}
