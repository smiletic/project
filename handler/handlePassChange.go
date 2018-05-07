package handler

import (
	"context"
	"masterRad/serverErr"
	"net/http"
)

func handlePassChange(ctx context.Context, r *http.Request, userUID string) (response interface{}, err error) {

	if r.Method == http.MethodPost {
		return nil, core.ChangePass(ctx, r, userUID)
	}

	return nil, serverErr.ErrInvalidAPICall
}
