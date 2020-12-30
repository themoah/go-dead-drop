package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// IndexHandler TODO: Should load html/webui
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("index")
	fmt.Fprintf(w, "hello, world !")
}

// HealthCheckHandler returns json ok
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := statusResponse{
		Status: StatusOk,
	}
	json.NewEncoder(w).Encode(response)
}

// StoreSecretHandler is exported for tests
// Key is used to for storing the drop
// Password will be used to decrypt it.
// assumes that data isn't zero
func StoreSecretHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/store")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	password := GenerateEncryptionKey()
	encryptedData := Encrypt(string(requestBody), password)
	// TODO: configurable length
	key := encryptedData[0:9]

	status := StoreDrop(key, encryptedData)

	if status != StatusOk {
		log.Println("failed to write the data")
		key = ""
		password = ""
	}

	// log.Println(" key: " + key + " ; body: " + password)
	response := storedSecretResponse{key, password}
	json.NewEncoder(w).Encode(response)
}

// RetrieveSecretHandler is exported for tests
func RetrieveSecretHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/retrieve")

	requestParams := mux.Vars(r)
	key := requestParams["key"]
	password := requestParams["password"]
	status, encryptedDrop := RetrieveDrop(key)
	// var decryptedData string
	decryptedData := ""
	if status == StatusOk {
		decryptedData = Decrypt(encryptedDrop, password)
	}

	response := retrievedDropResponse{
		Status: status,
		Data:   decryptedData,
	}
	json.NewEncoder(w).Encode(response)

}
