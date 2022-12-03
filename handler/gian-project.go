package handler

import (
	"Replace/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HandlerManager interface {
	HelloWorld(w http.ResponseWriter, r *http.Request)
}

func NewHandlerManager(manager service.ServiceManager) HandlerManager {
	return &handlerManager{
		ServiceManager: manager,
	}
}

type handlerManager struct {
	service.ServiceManager
}

func (hm *handlerManager) HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] init: HelloWorld()")
	defer func() {
		err := r.Body.Close()
		if err != nil {
			Response(err.Error(), http.StatusInternalServerError, w)
		}
	}()

	// GET URL PARAMS
	document := mux.Vars(r)
	if document["User"] == "" {
		Response("the id is required in url", http.StatusBadRequest, w)
		return
	}

	Response(fmt.Sprintf("Hello, nice to meet you %s!!!", document["User"]), http.StatusOK, w)
}

func Response(resp interface{}, statusCode int, w http.ResponseWriter) {
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic("error")
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}
