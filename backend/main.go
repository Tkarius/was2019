package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	fmt.Println("For those once loyal: ")
	inputReader := bufio.NewReader(os.Stdin)
	cfgSecret, _ := inputReader.ReadString('\n')
	cfgSecret = strings.Replace(cfgSecret, "\n", "", -1)
	env.selectUserSecret, env.insertUserSecret = decryptCfgs(cfgSecret)

	selectAnnouncements()

	router := InitRouter()
	fmt.Println("Http server starting up.")
	log.Fatal(http.ListenAndServe(":8000", router))
}

type Environment struct {
	selectUserSecret string
	insertUserSecret string
}

var env Environment
