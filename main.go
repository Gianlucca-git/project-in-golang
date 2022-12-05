package main

import (
	r "IMPORTS/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"IMPORTS/handler"
	repo "IMPORTS/repository"
	"IMPORTS/service"

	"github.com/gorilla/mux"
)

const defaultPort = ":5200"

func main() {
	log.Print("[INFO] init: main()")

	if err := repo.LoadSQLConnection(); err != nil {
		log.Fatal(err)
	}

	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	if port == ":" {
		port = defaultPort
	}

	router := mux.NewRouter().StrictSlash(true) // make the paths different from each other with slash /

	server := &http.Server{
		Addr:           port,             // port
		Handler:        router,           //
		ReadTimeout:    10 * time.Second, // reading time
		WriteTimeout:   10 * time.Second, // writing time
		MaxHeaderBytes: 1 << 20,          // 1mega in bits
	}
	log.Printf("[INFO] init: listen server... Port http://localhost%s", port)
	r.SetRoutes(router, initializedHandler())

	log.Fatal(server.ListenAndServe())
}

func initializedHandler() handler.Handler {
	return handler.Handler{
		HandlerManager: handler.NewHandlerManager(service.NewServiceManager(repo.PostgresSQL)),
	}
}
