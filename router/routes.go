package router

import (
	"Replace/handler"
	"github.com/gorilla/mux"
	"log"
)

func SetRoutes(router *mux.Router, handler handler.Handler) {
	log.Println("[INFO] init: SetRoutes()")
	router.HandleFunc("/hello/{User}", handler.HandlerManager.HelloWorld).Methods("GET")
}
