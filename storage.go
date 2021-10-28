package main

import "log"

//StoreDrop is used instead of interface
func (s secret) StoreDrop() (status string) {

	err := MemoryStore.Set(s.key, []byte(s.data))
	if err != nil {
		log.Printf("wow")
		return StatusError
	}
	return StatusOk
}

//RetrieveDrop gets and deletes a secret
func RetrieveDrop(key string) (status, encryptedDrop string) {

	byteData, err := MemoryStore.Get(key)
	if err != nil {
		return StatusError, ""
	}
	return StatusOk, string(byteData)

}
