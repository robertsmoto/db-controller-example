package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/robertsmoto/db_controller_example/api/handlers"
)

const (
	defaultPort = "8080"
)

var (
	port   string
	router *mux.Router
)

func main() {
	// change port by using $ PORT=10000 go run main.go
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter().StrictSlash(true)
	//middlewares
	r.Use(middleware.TimerMiddleware)
	// context middleware should run before others
	r.Use(middleware.ContextMiddleware)
	r.Use(middleware.RateLimiterMiddleware)
	r.Use(middleware.AccountAuthMiddleware)
	r.Use(middleware.ContentAuthMiddleware)
	//rest
	r.HandleFunc("/member", handlers.GetMemberHandler).
		Methods("GET")
	r.HandleFunc("/member", handlers.PutMemberHandler).
		Methods("PUT")

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Serving on --> ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
