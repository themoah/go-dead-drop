package main

import (
	"io/ioutil"
	"log"
	"os"
)

// TODO: create interface to store drops at different ends
// start with file, redis and then s3.

// StoreDropOnDisk writes encrypted data to the storage backend
// expects data to be non-null, as checked on previous stage
func StoreDropOnDisk(key, data string) (status string) {
	if data == "" {
		return StatusError
	}
	d1 := []byte(data)
	filepath := "/tmp/" + key

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
// status - ok or error (in case doesn't exists)
// encrypted data will be empty string if it failed.
// don't use in prod.
func RetrieveDropFromDisk(key string) (status, encryptedDropFromDisk string) {
	//TODO: mutex or other solution to the race condition.
	filepath := "/tmp/" + key
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
	go deleteFile(filepath)
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
