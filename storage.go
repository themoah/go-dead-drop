package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

//StoreDrop is used instead of interface
func StoreDrop(key, data string) (status string) {
	switch config.StorageEngine {
	case "localfile":
		return StoreDropOnDisk(key, data)
	case "redis":
		return StoreDropInRedis(key, data, Rdb)
	default:
		return StatusError
	}
}

//RetrieveDrop isn't interface
func RetrieveDrop(key string) (status, encryptedDropFromDisk string) {
	switch config.StorageEngine {
	case "localfile":
		return RetrieveDropFromDisk(key)
	case "redis":
		return RetrieveFromRedis(key, Rdb)
	default:
		return StatusError, ""
	}

}

var ctx = context.Background()

//StoreDropInRedis stores and sets TTL
func StoreDropInRedis(key, data string, rdb *redis.Client) (status string) {
	log.Println("storing in redis")

	// fmt.Printf("%T", rdb)

	ttl := time.Duration(config.DropExpiration) * time.Minute
	err := rdb.Set(ctx, key, data, ttl).Err()
	if err != nil {
		panic(err)
	}

	return StatusOk

}

//RetrieveFromRedis pulls encrypted drop and deletes it
func RetrieveFromRedis(key string, rdb *redis.Client) (status, enryptedDrop string) {
	log.Println("pulling from redis")
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return StatusError, ""
	}
	rdb.Del(ctx, key)

	return StatusOk, val

}

// StoreDropOnDisk writes encrypted data to the storage backend
// DONT USE IN PROD
// expects data to be non-null, as checked on previous stage
func StoreDropOnDisk(key, data string) (status string) {
	if data == "" {
		return StatusError
	}
	d1 := []byte(data)
	filepath := config.Localfile.BasePath + "/" + key

	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	e := ioutil.WriteFile(filepath, d1, 0644)
	if err != nil {
		panic(e)
	}

	defer f.Close()

	return StatusOk

}

// RetrieveDropFromDisk returns encrypted data and status.
// DONT USE IN PROD
// status - ok or error (in case doesn't exists)
// encrypted data will be empty string if it failed.
// don't use in prod.
func RetrieveDropFromDisk(key string) (status, encryptedDropFromDisk string) {
	//TODO: mutex or other solution to the race condition.
	filepath := config.Localfile.BasePath + "/" + key
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("failed to read the file")
		return StatusError, ""
	}

	stringData := string(dat)
	if stringData == "" {
		log.Println("oh, empty file")
		return StatusError, ""
	}
	deleteFile(filepath)
	return StatusOk, stringData
}

// an internal function to delete drop.
func deleteFile(filepath string) (status string) {
	e := os.Remove(filepath)
	if e != nil {
		log.Fatal(e)
		return StatusError
	}
	return StatusOk
}
