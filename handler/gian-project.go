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
	ClassifiedList(w http.ResponseWriter, r *http.Request)
	Balance(w http.ResponseWriter, r *http.Request)
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
			Response(struct {
				Message string `json:"message"`
			}{err.Error()}, http.StatusInternalServerError, w)
		}
	}()

	user := mux.Vars(r)

	Response(struct {
		Message string `json:"message"`
	}{fmt.Sprintf("Hello, nice to meet you %s!!!", user["User"])}, http.StatusOK, w)
}

// ClassifiedList Gets the request, decomposes it and sends it to the Service. Then reply to server
func (hm *handlerManager) ClassifiedList(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] init: ClassifiedList()")
	defer func() {
		err := r.Body.Close()
		if err != nil {
			Response(struct {
				Message string `json:"message"`
			}{err.Error()}, http.StatusInternalServerError, w)
		}
	}()

	var listResponse dto.ClassifiedList
	err := json.NewDecoder(r.Body).Decode(&listResponse)
	if err != nil {
		log.Printf("[Error] init: error decoding object in request (%s)", err.Error())
		Response(struct {
			Message string `json:"message"`
		}{http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest, w)
		return
	}

	err = hm.ServiceManager.OrderList(&listResponse)
	if err != nil {
		log.Printf("[INFO] init: proccess OrderList (%s)", err.Error())
		Response(struct {
			Message string `json:"message"`
		}{err.Error()}, http.StatusBadRequest, w)
		return
	}

	Response(listResponse, http.StatusOK, w)
}

// Balance Gets the request, decomposes it and sends it to the Service. Then reply to server
func (hm *handlerManager) Balance(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] init: Balance()")
	defer func() {
		err := r.Body.Close()
		if err != nil {
			Response(struct {
				Message string `json:"message"`
			}{err.Error()}, http.StatusInternalServerError, w)
		}
	}()

	var request dto.BalanceRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("[ERROR] init: error decoding object in request or bad request (%s)", err.Error())
		Response(struct {
			Message string `json:"message"`
		}{http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest, w)
		return
	}

	err = hm.ServiceManager.ValidatedRequestBalance(&request)
	if err != nil {
		log.Printf("[INFO] init: bad request (%s)", err.Error())
		Response(struct {
			Message string `json:"message"`
		}{err.Error()}, http.StatusBadRequest, w)
		return
	}

	urlVars := mux.Vars(r)
	response := hm.ServiceManager.GeneralBalance(&request, urlVars["filterMes"])

	Response(response, http.StatusOK, w)
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
