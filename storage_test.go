package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testKey  = "123"
	testData = "meow"
)

// Fails now because config is loaded in main.
// does IO
func TestStoreDropOnDisk(t *testing.T) {
	status := StoreDropOnDisk(testKey, testData)

	if status != StatusOk {
		t.Errorf("failed to write the file to the disk")
	}

	filepath := config.Localfile.BasePath + "/" + testKey
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	if string(dat) != testData {
		t.Errorf("stored and retrieved data didn't match")
	}

	assert.FileExists(t, filepath)
}

// Don't run it separately from previous one
// TODO: also check if file was deleted afterwards
func TestRetrieveDropFromDisk(t *testing.T) {
	status, dat := RetrieveDropFromDisk(testKey)

	if status != StatusOk {
		t.Errorf("failed to read from dik")
	}

	if dat != testData {
		t.Errorf("wrong data retrieved")
	}
	filepath := config.Localfile.BasePath + "/" + testKey
	assert.NoFileExists(t, filepath)
}
