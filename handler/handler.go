package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"projekat/data"
	"projekat/db"
	"projekat/dto"
	"projekat/enum"
	"projekat/serverErr"
	"projekat/util"
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
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, DELETE, PATCH, OPTIONS")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	dbRunner := db.CreateRunner(db.Handle)
	ctx = context.WithValue(ctx, db.RunnerKey, dbRunner)

	user, err := data.GetSession(ctx, authorization)
	if err != nil {
		http.Error(w, "Internal", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
	}
	ctx = context.WithValue(ctx, util.UserKey, user)

	r.URL.Path = r.URL.Path[5:]

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

	w.WriteHeader(httpResponseStatus)
	json.NewEncoder(w).Encode(response)
}

func handle(ctx context.Context, r *http.Request) (response interface{}, err error) {
	if strings.HasPrefix(r.URL.Path, "/admin") {
		user := ctx.Value(util.UserKey).(*dto.Authorization)
		if user.Role != enum.RoleAdmin {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[6:]
		response, err = handleAdmin(ctx, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/doctor") {
		user := ctx.Value(util.UserKey).(*dto.Authorization)
		if user.Role != enum.RoleDoctor {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[7:]
		response, err = handleDoctor(ctx, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/research") {
		user := ctx.Value(util.UserKey).(*dto.Authorization)
		if user.Role != enum.RoleResearch {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[9:]
		response, err = handleResearcher(ctx, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/nurse") {
		user := ctx.Value(util.UserKey).(*dto.Authorization)
		if user.Role != enum.RoleNurse {
			err = serverErr.ErrNotAuthenticated
			return
		}
		r.URL.Path = r.URL.Path[6:]
		response, err = handleNurse(ctx, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/pass") {
		r.URL.Path = r.URL.Path[5:]
		user := ctx.Value(util.UserKey).(*dto.Authorization)
		response, err = handlePassChange(ctx, r, user.UserUID)
		return
	}
	err = serverErr.ErrInvalidAPICall
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	dbRunner := db.CreateRunner(db.Handle)
	ctx = context.WithValue(ctx, db.RunnerKey, dbRunner)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, DELETE, PATCH, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	type loginRequest struct {
		Username string
		Password string
	}

	type loginResponse struct {
		Username      string
		Authenticated bool
		Authorization string
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

	token, err := util.CreateSession(ctx, user.UserUID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Authenticated = true
	response.Role = user.Role
	response.Username = request.Username
	response.Authorization = token
	buf := make([]byte, 0, 1000)
	responsew := bytes.NewBuffer(buf)
	json.NewEncoder(responsew).Encode(response)
	w.Write(responsew.Bytes())

}

func logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, DELETE, PATCH, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
	}
	ctx := context.Background()
	dbRunner := db.CreateRunner(db.Handle)
	ctx = context.WithValue(ctx, db.RunnerKey, dbRunner)
	err := data.RemoveSession(ctx, authorization)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
