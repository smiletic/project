package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"masterRad/db"
	"masterRad/util"
	"net/http"

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

	handle(w, r)

}

func handle(w http.ResponseWriter, r *http.Request) {
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
