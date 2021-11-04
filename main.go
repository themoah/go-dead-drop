package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	bigcache "github.com/allegro/bigcache/v3"

	"github.com/gorilla/mux"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//StatusOk for sharing result and not boolean/err
const (
	defaultPort = "8080"
	StatusOk    = "ok"
	StatusError = "error"
	APIVersion  = "0.1"
)

var MemoryStore *bigcache.BigCache

func main() {
	parseConfig()

	log := zerolog.New(os.Stdout)
	// zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if checkIfDevEnv() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	err := setupStorageEngine()
	if err != nil {
		panic(err)
	}

	log.Info().Msg("starting go-dead-drop, listening on port 0.0.0.0:" + config.Port)
	log.Debug().Msg(fmt.Sprintf("Maximum capacity of memory is %v values ", MemoryStore.Capacity()))

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/version", versionHandler).Methods("GET")
	r.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/store", storeSecretHandler).Methods("POST")
	// TODO: maybe also with only 1 param - base64 key and password
	r.HandleFunc("/retrieve/{key}/{password}", RetrieveSecretHandler).Methods("POST")

	http.ListenAndServe("0.0.0.0:"+config.Port, r)

}

func checkIfDevEnv() bool {

	val, present := os.LookupEnv("DEBUG")
	if !present {
		log.Info().Msg("not dev environment")
		return false
	} else if val == "true" {
		log.Info().Msg("it's a dev environment")
		return true
	} else {
		log.Info().Msg("!DEBUG=true")
		return false
	}

}

func setupStorageEngine() error {

	bcConfig := bigcache.Config{
		Shards:      16,
		LifeWindow:  time.Minute * time.Duration(config.DropExpiration),
		CleanWindow: 15 * time.Minute,
		//used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// im bytes
		MaxEntrySize:       9999,
		Verbose:            true,
		StatsEnabled:       true,
		HardMaxCacheSize:   0,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	}

	cache, err := bigcache.NewBigCache(bcConfig)
	if err != nil {
		log.Error().Err(err).Msg("failed to create big cache")
		return err
	}

	MemoryStore = cache

	return nil

}
