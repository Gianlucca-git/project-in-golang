package router

import (
	"IMPORTS/handler"
	"github.com/gorilla/mux"
	"log"
)

func SetRoutes(router *mux.Router, handler handler.Handler) {
	log.Print("[INFO] init: SetRoutes()")
	router.HandleFunc("/hello/{User}", handler.HandlerManager.HelloWorld).Methods("GET")
	router.HandleFunc("/classified", handler.HandlerManager.ClassifiedList).Methods("POST")
	router.HandleFunc("/balance/{filterMes}", handler.HandlerManager.Balance).Methods("POST")

}
