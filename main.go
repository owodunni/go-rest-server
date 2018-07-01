package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", Test)
	router.HandleFunc("/hola/{name}", Hola)
	router.HandleFunc("/validate/{CD-KEY}", Validate)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Test(w http.ResponseWriter, r *http.Request) {
	response := HipChatResponse{Color: "yellow", Message: "This is a Test", Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

func Hola(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := HipChatResponse{Color: "yellow", Message: "Hola " + name, Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cdKey := vars["CD-KEY"]
	response := ValidateResponse{CdKey: "Valid cdKey " + cdKey}
	json.NewEncoder(w).Encode(response)
}

type HipChatResponse struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        string `json:"notify"`
	MessageFormat string `json:"message_format"`
}

type ValidateResponse struct {
	CdKey string `json:"CdKey"`
}
