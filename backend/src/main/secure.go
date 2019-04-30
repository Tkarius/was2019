package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"

	_ "github.com/denisenkom/go-mssqldb"
)

func encrypt(plainText string, key string) {
	toEncrypt := []byte(plainText)
	keyBytes := []byte(key)
	ciph, err := aes.NewCipher(keyBytes)

	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(ciph)

	if err != nil {
		panic(err)
	}

	// create a nonce and populate it with cryptographically randomized sequence.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("non_existant_file.cfg", gcm.Seal(nonce, nonce, toEncrypt, nil), 0444)
}

func decrypt(cipherText string, key string) string {
	toDecrypt := []byte(cipherText)
	keyBytes := []byte(key)
	ciph, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, toDecrypt := toDecrypt[:nonceSize], toDecrypt[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, toDecrypt, nil)
	if err != nil {
		panic(err)
	}
	return string(decrypted)
}

func decryptCfgs(cfgSecret string) (string, string) {
	selectSecret, err := ioutil.ReadFile("select.cfg")
	selectSecretString := decrypt(string(selectSecret), cfgSecret)
	if err != nil {
		panic(err)
	}
	insertSecret, err := ioutil.ReadFile("insert.cfg")
	insertSecretString := decrypt(string(insertSecret), cfgSecret)
	if err != nil {
		panic(err)
	}
	return selectSecretString, insertSecretString
}
