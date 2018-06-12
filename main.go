package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"net/http"
	"projekat/config"
	"projekat/dto"
	"projekat/handler"
	"projekat/server"
	"time"

	"github.com/gorilla/context"
)

func main() {

	var err error
	gob.Register(dto.Authorization{})

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
