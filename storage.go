package main

// TODO: create interface to store drops at different ends
// start with file, mongodb and then s3.

// StoreDrop writes encrypted data to the storage backend
func StoreDrop(filename string) {}

// RetrieveDrop returns encrypted data if it wasn't stored before.
// returns nothing if file doesn't exists (possibly was already retrieved)
// not every storage backend can be "acid compliant":
// e.g. eliminates race condition when same "drop" would be retrieved more than once.
func RetrieveDrop(filename string) {}

// an internal function to delete drop.
func deleteFile(filename string) {}
