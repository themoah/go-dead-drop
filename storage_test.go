package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testKey  = "123"
	testData = "meow"
)

// Fails now because config is loaded in main.
// does IO
func TestStore(t *testing.T) {

	s := secret{
		key:  testKey,
		data: testData,
	}
	status := s.StoreDrop()

	assert.Equal(t, status, StatusOk, "should be equal")
}

// TODO: also check if file was deleted afterwards
func TestRetrieveDropFromDisk(t *testing.T) {
	status, dat := RetrieveDrop(testKey)

	assert.Equal(t, status, StatusOk, "should be equal")
	assert.Equal(t, testData, string(dat), "data should be equal")
}
