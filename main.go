package main

import (
	r "Replace/router"
	"fmt"
	"log"
	"net/http"
	"time"

	"Replace/handler"
	repo "Replace/repository"
	"Replace/service"

	"github.com/gorilla/mux"
)

const defaultPort = ":4200"

func main() {
	log.Println("[INFO] init: main()")
	fmt.Println("[INFO] ftm init: main()")

	//if err := repo.LoadSQLConnection(); err != nil {
	//	log.Fatal(err)
	//}

	router := mux.NewRouter().StrictSlash(true) // make the paths different from each other with slash /

	server := &http.Server{
		Addr:           defaultPort,      // port
		Handler:        router,           //
		ReadTimeout:    10 * time.Second, // reading time
		WriteTimeout:   10 * time.Second, // writing time
		MaxHeaderBytes: 1 << 20,          // 1mega in bits
	}
	log.Println("[INFO] init: listen server... Port http://localhost:4200")
	r.SetRoutes(router, initializedHandler())

	log.Fatal(server.ListenAndServe())
}

func initializedHandler() handler.Handler {
	return handler.Handler{
		HandlerManager: handler.NewHandlerManager(service.NewServiceManager(repo.PostgresSQL)),
	}
}
