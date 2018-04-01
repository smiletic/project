package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"masterRad/db"
	"masterRad/dto"
	"masterRad/enum"
	"masterRad/serverErr"
	"masterRad/util"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

var (
	Login            = login
	Logout           = logout
	HandleAuthorized = handleAuthorized
)

var (
	store = sessions.NewCookieStore([]byte("something-very-secret"))
)

func handleAuthorized(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["user"]; !ok || auth == nil {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
	}
	r.URL.Path = r.URL.Path[5:]
	ctx := context.Background()
	dbRunner := db.CreateRunner(db.Handle)
	ctx = context.WithValue(ctx, db.RunnerKey, dbRunner)
	response, err := handle(ctx, r)
	var httpResponseStatus int
	if err == nil {
		httpResponseStatus = http.StatusOK
	} else {
		switch err {
		case serverErr.ErrBadRequest:
			httpResponseStatus = http.StatusBadRequest
		case serverErr.ErrNotAuthenticated:
			httpResponseStatus = http.StatusUnauthorized
		case serverErr.ErrForbidden:
			httpResponseStatus = http.StatusForbidden
		case serverErr.ErrInvalidAPICall, serverErr.ErrResourceNotFound:
			httpResponseStatus = http.StatusNotFound
		case serverErr.ErrMethodNotAllowed:
			httpResponseStatus = http.StatusMethodNotAllowed
		default:
			httpResponseStatus = http.StatusInternalServerError
		}
	}

	// Write response.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpResponseStatus)
	json.NewEncoder(w).Encode(response)
}

func handle(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/admin") {
		session, _ := store.Get(r, "cookie-name")
		if session.Values["user"].(dto.Authorization).Role != enum.RoleAdmin {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[6:]
		response, err = handleAdmin(ctx, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/doctor") {
		session, _ := store.Get(r, "cookie-name")
		if session.Values["user"].(dto.Authorization).Role != enum.RoleDoctor {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[7:]
		response, err = handleDoctor(ctx, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/research") {
		session, _ := store.Get(r, "cookie-name")
		if session.Values["user"].(dto.Authorization).Role != enum.RoleDoctor {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[9:]
		response, err = handleResearcher(ctx, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/nurse") {
		session, _ := store.Get(r, "cookie-name")
		if session.Values["user"].(dto.Authorization).Role != enum.RoleNurse {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[6:]
		response, err = handleNurse(ctx, r)
		return
	}
	err = serverErr.ErrInvalidAPICall
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	dbRunner := db.CreateRunner(db.Handle)
	ctx = context.WithValue(ctx, db.RunnerKey, dbRunner)
	session, _ := store.Get(r, "cookie-name")

	type loginRequest struct {
		Username string
		Password string
	}

	type loginResponse struct {
		Authenticated bool
		Role          enum.Role
	}
	request := &loginRequest{}
	response := &loginResponse{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := util.Login(ctx, request.Username, request.Password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session.Values["user"] = user
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}
	response.Authenticated = true
	response.Role = user.Role
	buf := make([]byte, 0, 1000)
	responsew := bytes.NewBuffer(buf)
	json.NewEncoder(responsew).Encode(response)
	w.Write(responsew.Bytes())

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["user"] = nil
	session.Save(r, w)
}
