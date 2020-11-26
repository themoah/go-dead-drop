package main

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key1 := GenerateEncryptionKey()
	key2 := GenerateEncryptionKey()

	if key1 == key2 {
		t.Errorf("generated two same keys %q and %q", key1, key2)
	}
}

func TestEncrypt(t *testing.T) {
	source := "Hello, World!"
	key := GenerateEncryptionKey()

	encryptedSource := Encrypt(source, key)

	if source == encryptedSource {
		t.Errorf("encrypted source is the same as unecrypted %q == %q", source, encryptedSource)
	}
}

func TestE2E(t *testing.T) {
	source := "Hello, World!"
	key := GenerateEncryptionKey()

	encryptedSource := Encrypt(source, key)
	decryptedSouce := Decrypt(encryptedSource, key)

	if decryptedSouce != source {
		t.Errorf("e2e encryption test failed %q %q", source, decryptedSouce)
	}
}
