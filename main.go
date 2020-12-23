package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//StatusOk for sharing result and not boolean/err
const (
	defaultPort = "8080"
	StatusOk    = "ok"
	StatusError = "error"
)

// TODO: add https://github.com/google/go-cloud for cloud and db operations.
func main() {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
		fmt.Printf("Defaulting to port %s \n", serverPort)
	}

	fmt.Println("starting server, listening on port 0.0.0.0:" + serverPort + "\n")

	r := mux.NewRouter()
	fmt.Println("go-dead-drop")

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/healthz", healthCheckHandler).Methods("GET")
	r.HandleFunc("/store", storeSecretHandler).Methods("POST")
	// TODO: maybe also with only 1 param - base64 key and password
	r.HandleFunc("/retrieve/{key}/{password}", retrieveSecretHandler).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:"+serverPort, r))

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := statusResponse{
		Status: StatusOk,
	}
	json.NewEncoder(w).Encode(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("index 200")
	fmt.Fprintf(w, "hello, world !")
}

func storeSecretHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("store secret")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(requestBody))
	key, password := StoreSecret(string(requestBody))
	log.Println(" key: " + key + " ; body: " + password)
	response := storedSecretResponse{
		Key:      key,
		Password: password,
	}
	json.NewEncoder(w).Encode(response)
}

func retrieveSecretHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getting the secret")

	requestParams := mux.Vars(r)
	key := requestParams["key"]
	password := requestParams["password"]
	status, encryptedDrop := RetrieveDropFromDisk(key)
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
