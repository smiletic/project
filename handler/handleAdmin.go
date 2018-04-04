package handler

import (
	"context"
	"masterRad/core"
	"masterRad/serverErr"
	"net/http"
	"strings"
)

func handleAdmin(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/person") {
		r.URL.Path = r.URL.Path[7:]
		switch r.Method {
		case http.MethodPost:
			return core.CreatePerson(ctx, r.Body)
		case http.MethodPatch:
			return nil, core.UpdatePerson(ctx, r.URL.Path[1:], r.Body)
		case http.MethodGet:
			if strings.HasPrefix(r.URL.Path, "/") {
				return core.GetPerson(ctx, r.URL.Path[1:])
			}
			return core.GetPersons(ctx, r.URL.Query())
		case http.MethodDelete:
			return nil, core.RemovePerson(ctx, r.URL.Path[1:])
		}
	}
	if strings.HasPrefix(r.URL.Path, "/employee") {
		r.URL.Path = r.URL.Path[9:]
		switch r.Method {
		case http.MethodPost:
			return core.CreateEmployee(ctx, r.Body)
		case http.MethodPatch:
			return nil, core.UpdateEmployee(ctx, r.URL.Path[1:], r.Body)
		case http.MethodGet:
			if strings.HasPrefix(r.URL.Path, "/") {
				return core.GetEmployee(ctx, r.URL.Path[1:])
			}
			return core.GetEmployees(ctx)
		case http.MethodDelete:
			return nil, core.RemoveEmployee(ctx, r.URL.Path[1:])
		}
	}
	if strings.HasPrefix(r.URL.Path, "/user") {
		switch r.Method {
		case http.MethodPost:
		case http.MethodPatch:
		case http.MethodGet:
		case http.MethodDelete:
		}
	}

	return nil, serverErr.ErrInvalidAPICall
}
