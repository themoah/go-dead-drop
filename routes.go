package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

// IndexHandler TODO: Should load html/webui
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/index")
	fmt.Fprintf(w, "hello, world !")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("/version")
	fmt.Fprintf(w, APIVersion)
}

// HealthCheckHandler returns json ok
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := statusResponse{
		Status: StatusOk,
	}
	w.Header().Add("Content-Type", "application/json")
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

	var s secret
	s.key = key
	s.data = encryptedData

	status := s.StoreDrop()

	if status != StatusOk {
		log.Println("failed to store the drop")
		key = ""
		password = ""
	}

	// log.Println(" key: " + key + " ; body: " + password)
	response := storedSecretResponse{key, password}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RetrieveSecretHandler is exported for tests
func RetrieveSecretHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/retrieve")

	requestParams := mux.Vars(r)
	key := requestParams["key"]
	password := requestParams["password"]
	decryptedData := ""
	var status, encryptedDrop string
	keyTest := verifyKey(key)
	if keyTest == StatusOk {
		status, encryptedDrop = RetrieveDrop(key)
	} else {
		//debug log
		log.Println("very bad key")
	}
	if status == StatusOk {
		decryptedData = Decrypt(encryptedDrop, password)
	}
	response := retrievedDropResponse{
		Status: status,
		Data:   decryptedData,
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func verifyKey(key string) (status string) {
	// IF FOUND == MEANS BAD
	match, err := regexp.MatchString("[^a-z0-9]+", key)
	if err != nil {
		log.Println("bad key " + key)
		return StatusError
	}
	if match {
		log.Println("key is bad")
		return StatusError
	}
	return StatusOk

}
