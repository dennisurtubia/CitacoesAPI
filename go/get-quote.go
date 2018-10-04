package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"encoding/json"
	"math/rand"
	"time"
)

type Citacao struct {
	Id		string
	Tema	string
	Author	string
	Quote	string
}

type CitacaoWrapper struct {
	Citacao []Citacao
}

func main() {
    b, errFile := ioutil.ReadFile("../citacoes.json") // just pass the file name
    if errFile != nil {
		log.Fatal(errFile)
    }
   	var citacoesWrapper CitacaoWrapper
	errJson := json.Unmarshal(b, &citacoesWrapper)
	if errJson != nil {
		log.Fatal(errJson)
	}
	rand.Seed(time.Now().UnixNano())
	randomPosition := rand.Intn(len(citacoesWrapper.Citacao))
	fmt.Println(citacoesWrapper.Citacao[randomPosition].Quote)

}