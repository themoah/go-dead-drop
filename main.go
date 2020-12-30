package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	memkv "github.com/kelseyhightower/memkv"
)

//StatusOk for sharing result and not boolean/err
const (
	defaultPort = "8080"
	StatusOk    = "ok"
	StatusError = "error"
)

// Rdb global value, needs refactor
var Rdb *redis.Client

// MemoryStore default storage engine
var MemoryStore memkv.Store

// TODO: maybe use https://github.com/google/go-cloud for cloud and db operations.
func main() {
	parseConfig()

	if config.StorageEngine == "redis" {
		Rdb = redis.NewClient(&redis.Options{
			Addr: config.Redis.Addr + ":" + config.Redis.Port,
		})
	} else if config.StorageEngine == "memory" {
		MemoryStore = memkv.New()
	}

	log.Println("starting go-dead-drop, listening on port 0.0.0.0:" + config.Port + "\n")

	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/store", StoreSecretHandler).Methods("POST")
	// TODO: maybe also with only 1 param - base64 key and password
	r.HandleFunc("/retrieve/{key}/{password}", RetrieveSecretHandler).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:"+config.Port, r))

}
