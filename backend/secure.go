package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func encrypt(plainText string, key string) []byte {
	toEncrypt := []byte(plainText)
	keyBytes := []byte(key)
	ciph, err := aes.NewCipher(keyBytes)

	if err != nil {
		fmt.Printf("DEBUG: Error while encrypting: %s with: %s\n", plainText, key)
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

func decrypt(cipherText string, key string) string {
	toDecrypt := []byte(cipherText)
	keyBytes := []byte(key)
	ciph, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Printf("DEBUG: Error while decrypting: %s with: %s\n", cipherText, key)
		panic(err)
	}

	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		fmt.Printf("DEBUG: Error while entering Galois/Counter Mode.\n")
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, toDecrypt := toDecrypt[:nonceSize], toDecrypt[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, toDecrypt, nil)
	if err != nil {
		fmt.Println("DEBUG: Error while decrypting.")
		panic(err)
	}
	fmt.Printf("DEBUG: decryption complete: %s\n", string(decrypted))
	return string(decrypted)
}

func decryptCfgs(cfgSecret string) string {
	decryptedCfg, err := ioutil.ReadFile("was-server.cfg")
	if err != nil {
		fmt.Println("DEBUG: Error while reading cfg")
		panic(err)
	}
	return decrypt(string(decryptedCfg), cfgSecret)
}
