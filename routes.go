package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// IndexHandler TODO: Should load html/webui
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("/index")
	fmt.Fprintf(w, "hello, world !")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("/version")
	fmt.Fprintf(w, APIVersion)
}

// HealthCheckHandler returns json ok
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("/health")
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
func storeSecretHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("/store")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("can't read requestBody")
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
		log.Error().Err(nil).Msg("failed to store the drop")
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
	log.Debug().Msg("/retrieve")

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
		log.Debug().Msg("very bad key")
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
		log.Error().Err(nil).Msg("bad key " + key)
		return StatusError
	}
	if match {
		log.Error().Msg("key is bad")
		return StatusError
	}
	return StatusOk

}
