package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

// StoreSecret runs the flow.
// Key is used to find and retrieve the secret.
// Password will be used to decrypt it.
// assumes that data isn't zero
// still doesn't store it anywhere.
func StoreSecret(data string) (key, password string) {
	password = GenerateEncryptionKey()
	encryptedData := Encrypt(data, password)

	// TODO: here it goes to storing data.
	// fmt.Printf("### Encrypted secret: \n " + encryptedData + " \n ### End of encrypted data \n")

	key = encryptedData[0:9]

	if !StoreDropOnDisk(key, encryptedData) {
		log.Println("failed to write the data")
	}

	return key, password

}

// RetrieveSecret finds secret by key and decrypts data.
// If key doesn't exist/was already retrieved/etc just returns empty string
// func RetrieveSecret(key, password string) (decryptedData string) {

// }

// based on https://gist.github.com/donvito/efb2c643b724cf6ff453da84985281f8
// TODO: check if it's better to use https://godoc.org/gocloud.dev/secrets#example-package
// or https://github.com/m1/go-generate-password

//GenerateEncryptionKey would be passed to the end user.
func GenerateEncryptionKey() (encryptionKey string) {
	// TODO: create stronger key (crypto).
	// TODO: configurable length of the key.
	bytes := make([]byte, 16) //generate a random 16 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes)
	// fmt.Printf("encryption key : %s\n", key)

	return key
}

// Encrypt with key (aes)
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

//Decrypt with user provided key.
func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
