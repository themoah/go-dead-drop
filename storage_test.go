package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	key  = "123"
	data = "meow"
)

// Fails now because config is loaded in main.
// does IO
func TestStoreDropOnDisk(t *testing.T) {
	status := StoreDropOnDisk(key, data)

	if status != StatusOk {
		t.Errorf("failed to write the file to the disk")
	}

	filepath := config.Localfile.BasePath + "/" + key
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	if string(dat) != data {
		t.Errorf("stored and retrieved data didn't match")
	}

	assert.FileExists(t, filepath)
}

// Don't run it separately from previous one
// TODO: also check if file was deleted afterwards
func TestRetrieveDropFromDisk(t *testing.T) {
	status, dat := RetrieveDropFromDisk(key)

	if status != StatusOk {
		t.Errorf("failed to read from dik")
	}

	if dat != data {
		t.Errorf("wrong data retrieved")
	}
	filepath := config.Localfile.BasePath + "/" + key
	assert.NoFileExists(t, filepath)
}
