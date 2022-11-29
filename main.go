package main

import (
	"fmt"
	"net/http"

	"github.com/robertsmoto/db_controller_example/api/handlers"
	"github.com/robertsmoto/db_controller_example/repo/sqldb"
)

func main() {
	// Create connection
	db := sqldb.ConnectDB()
	// Create repos (implementations)
	userRepo := sqldb.NewUserRepo(db)

	h := handlers.NewBaseHandler(userRepo)
	http.HandleFunc("/", h.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}
	s.ListenAndServe()

}
