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
func StoreDropOnDisk(key, data string) (storageSuccess bool) {
	if data == "" {
		return false
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

	return true

}

// RetrieveDropFromDisk returns encrypted data if it wasn't stored before.
// returns nothing if file doesn't exists (possibly was already retrieved)
// not every storage backend can be "acid compliant":
// e.g. eliminates race condition when same "drop" would be retrieved more than once.
func RetrieveDropFromDisk(key string) (encryptedDropFromDisk string) {
	//TODO: mutex or other solution to the race condition.
	filepath := "/tmp/" + key
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	stringData := string(dat)
	if stringData == "" {
		log.Println("oh, something is bad")
	}
	go deleteFile(filepath)
	return stringData
}

// an internal function to delete drop.
func deleteFile(filepath string) {
	e := os.Remove(filepath)
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
