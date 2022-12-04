package handler

import (
	"IMPORTS/model/dto"
	"IMPORTS/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HandlerManager interface {
	// HelloWorld this method is in charge of greeting the user
	HelloWorld(w http.ResponseWriter, r *http.Request)
	// ClassifiedList Gets the request, decomposes it and sends it to the Service. Then reply to server
	ClassifiedList(w http.ResponseWriter, r *http.Request)
}

func NewHandlerManager(manager service.ServiceManager) HandlerManager {
	return &handlerManager{
		ServiceManager: manager,
	}
}

type handlerManager struct {
	service.ServiceManager
}

// HelloWorld this method is in charge of greeting the user
func (hm *handlerManager) HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] init: HelloWorld()")
	defer func() {
		err := r.Body.Close()
		if err != nil {
			Response(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w)
		}
	}()

	user := mux.Vars(r)

	Response(fmt.Sprintf("Hello, nice to meet you %s!!!", user["User"]), http.StatusOK, w)
}

// ClassifiedList Gets the request, decomposes it and sends it to the Service. Then reply to server
func (hm *handlerManager) ClassifiedList(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] init: ClassifiedList()")
	defer func() {
		err := r.Body.Close()
		if err != nil {
			Response(err.Error(), http.StatusInternalServerError, w)
		}
	}()

	var listResponse dto.ClassifiedList
	err := json.NewDecoder(r.Body).Decode(&listResponse)
	if err != nil {
		log.Printf("[Error] init: error decoding object in request (%s)", err.Error())
		Response(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
		return
	}

	err = hm.ServiceManager.OrderList(&listResponse)
	if err != nil {
		log.Printf("[INFO] init: proccess OrderList (%s)", err.Error())
		Response(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w)
		return
	}

	Response(listResponse, http.StatusOK, w)
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
