package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(plainText string, key string) []byte {
	toEncrypt := []byte(plainText)
	keyBytes := []byte(key)
	ciph, err := aes.NewCipher(keyBytes)

	if err != nil {
		fmt.Println("DEBUG: Error while encrypting: %s with: %s", plainText, key)
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(ciph)

	if err != nil {
		fmt.Println("DEBUG: Error while entering Galois/Counter Mode.")
		fmt.Println(err)
	}

	// create a nonce and populate it with cryptographically randomized sequence.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("DEBUG: Error while populating nonce with randomized sequence.")
	}

	return gcm.Seal(nonce, nonce, toEncrypt, nil)

}

func decrypt() {

}
