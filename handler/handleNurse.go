package handler

import (
	"context"
	"masterRad/core"
	"masterRad/serverErr"
	"net/http"
	"strings"
)

func handleNurse(ctx context.Context, r *http.Request) (response interface{}, err error) {
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
			return core.GetPersons(ctx)
		case http.MethodDelete:
			return nil, core.RemovePerson(ctx, r.URL.Path[1:])
		}
	}
	if strings.HasPrefix(r.URL.Path, "/patient") {
		r.URL.Path = r.URL.Path[8:]
		switch r.Method {
		case http.MethodPost:
			return core.CreatePatient(ctx, r.Body)
		case http.MethodPatch:
			return nil, core.UpdatePatient(ctx, r.URL.Path[1:], r.Body)
		case http.MethodGet:
			if strings.HasPrefix(r.URL.Path, "/") {
				return core.GetPatient(ctx, r.URL.Path[1:])
			}
			return core.GetPatients(ctx)
		case http.MethodDelete:
			return nil, core.RemovePatient(ctx, r.URL.Path[1:])
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
