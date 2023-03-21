package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	entity "github.com/cassiusbessa/create-service/entity"
	errors "github.com/cassiusbessa/create-service/errors"
	repository "github.com/cassiusbessa/create-service/repository"
	"github.com/gorilla/mux"
)

type Response struct {
	Message interface{} `json:"message"`
}

var _, cancel = repository.MongoConnection()

func CreateService(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	var service entity.Service
	err := json.NewDecoder(req.Body).Decode(&service)
	if err != nil {
		errors.SendError(res, http.StatusBadRequest, Response{"Invalid request payload"})
		return
	}

	validateError := service.Validate()
	if validateError != nil {
		errors.SendError(res, http.StatusBadRequest, validateError)
		return
	}

	db := mux.Vars(req)["company"]
	err = repository.CreateService(db, service)
	if err != nil {
		repository.SaveError(db, *errors.NewError(http.StatusInternalServerError, "Error Mongo creating service", "CreateService", err))
		errors.SendError(res, http.StatusInternalServerError, err.Error())
		return
	}
	res.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(res).Encode(Response{Message: "Service created successfully"})
	if err != nil {
		repository.SaveError(db, *errors.NewError(http.StatusInternalServerError, "Error encoding response", "CreateService", err))
		errors.SendError(res, http.StatusInternalServerError, err.Error())
		return
	}
	defer cancel()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/service/{company}", CreateService).Methods("POST")
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)

}
