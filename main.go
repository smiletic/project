package main

import (
	"encoding/gob"
	"fmt"
	"masterRad/config"
	"masterRad/dto"
	"masterRad/handler"
	"masterRad/server"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// server main method

var router = mux.NewRouter()

func main() {

	var err error
	gob.Register(dto.Autorizacija{})

	// Seed function is part of rand initialization.
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Initializing config")
	_, err = config.InitConfig("server", nil)
	if err != nil {
		fmt.Printf("Could not initialize config: %v\n", err)
		return
	}

	fmt.Println("Initializing database")
	err = server.InitializeDb()
	if err != nil {
		fmt.Printf("Could not access database: %v\n", err)
		return
	}

	http.HandleFunc("/auth/", handler.HandleAuthorized)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
