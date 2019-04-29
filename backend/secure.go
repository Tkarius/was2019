package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"

	_ "github.com/denisenkom/go-mssqldb"
)

func encrypt(plainText string, key string) {
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
	err = ioutil.WriteFile("select.cfg", gcm.Seal(nonce, nonce, toEncrypt, nil), 0444)
	//return gcm.Seal(nonce, nonce, toEncrypt, nil)

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

func decryptCfgs(cfgSecret string) (string, string) {
	selectSecret, err := ioutil.ReadFile("select.cfg")
	selectSecretString := decrypt(string(selectSecret), cfgSecret)
	if err != nil {
		fmt.Println("DEBUG: Error while reading cfg")
		panic(err)
	}
	insertSecret, err := ioutil.ReadFile("insert.cfg")
	insertSecretString := decrypt(string(insertSecret), cfgSecret)
	if err != nil {
		fmt.Println("DEBUG: Error while reading cfg")
		panic(err)
	}
	return selectSecretString, insertSecretString
}
