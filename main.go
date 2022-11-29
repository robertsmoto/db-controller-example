package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/robertsmoto/db_controller_example/api/handlers"
	"github.com/robertsmoto/db_controller_example/api/middleware"
	"github.com/robertsmoto/db_controller_example/config"
	"github.com/robertsmoto/db_controller_example/repo/redisdb"
)

const (
	defaultPort = "8080"
)

var (
	port   string
	router *mux.Router
)

func main() {
	conf := config.Conf
	port := defaultPort
	// can change port by using env.yaml
	if conf.ApiPort != "" {
		port = conf.ApiPort
	}
	// can change port by using $ PORT=10000 go run main.go
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	// make db connections for each db Instance
	apiDb := redisdb.ConnectDB(redisdb.Api)
	defer apiDb.Close()
	acctDb := redisdb.ConnectDB(redisdb.Account)
	defer acctDb.Close()
	dataDb := redisdb.ConnectDB(redisdb.Data)
	defer dataDb.Close()
	// create middleware connections structs, for those that need it
	apiConn := middleware.NewMiddlewareConnector(&apiDb, conf)
	acctConn := middleware.NewMiddlewareConnector(&acctDb, conf)
	// connect handler(s)
	memberDAL := redisdb.NewMemberDAL(redisdb.NewBaseDAL(dataDb, conf))
	memberHandler := handlers.NewBaseHandler(memberDAL)
	// router
	r := mux.NewRouter().StrictSlash(true)
	//middlewares
	r.Use(middleware.TimerMiddleware) //<-- no connection
	// context middleware should run before others
	r.Use(middleware.ContextMiddleware) //<-- no connection
	r.Use(apiConn.RateLimiterMiddleware)
	r.Use(acctConn.AccountAuthMiddleware)
	r.Use(middleware.ContentAuthMiddleware) //<-- no connection
	// routes
	r.HandleFunc("/member", memberHandler.GetMemberHandler).
		Methods("GET")
	r.HandleFunc("/member", memberHandler.PutMemberHandler).
		Methods("PUT")
		// more here

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
