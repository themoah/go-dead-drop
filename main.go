package main

import (
	"fmt"
)

// TODO: add https://github.com/google/go-cloud for cloud and db operations.
func main() {

	fmt.Println("go-dead-drop")

	for x := 0; x < 10; x++ {
		fmt.Println("new key : " + GenerateEncryptionKey())
	}

}
