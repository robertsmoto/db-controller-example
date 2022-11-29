package main

import (
	"fmt"
	"net/http"

	"github.com/robertsmoto/db-controller-examples/api/handlers"
	"github.com/robertsmoto/db-controller-examples/repo"
	"github.com/robertsmoto/db-controller-examples/repo/sqldb"
)

func main() {
	db := sqldb.ConnectDB()

	// Create repos
	userRepo := repo.NewUserRepo(db)

	h := handlers.NewBaseHandler(userRepo)

	http.HandleFunc("/", h.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()

}
