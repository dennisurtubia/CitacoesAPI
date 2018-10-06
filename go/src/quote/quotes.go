package quote

import (
    "io/ioutil"
    "log"
	"encoding/json"
	"math/rand"
	"time"
	"path/filepath"
	"os"
    "regexp"
)

func getAllQuotes() Book {
	re := regexp.MustCompile("go$")
	goPath := os.Getenv("GOPATH")
    path := re.ReplaceAllString(goPath, "citacoes.json")
	b, errFile := ioutil.ReadFile(filepath.FromSlash(path)) // just pass the file name
    if errFile != nil {
		log.Fatal(errFile)
    }
   	var book Book
	errJson := json.Unmarshal(b, &book)
	if errJson != nil {
		log.Fatal(errJson)
	}
	return book
}

func GetRandomQuote() Quote {
	quotes := getAllQuotes()	
	rand.Seed(time.Now().UnixNano())
	randomPosition := rand.Intn(len(quotes.Citacao))
	return quotes.Citacao[randomPosition]
}

type Quote struct {
	Id		string
	Tema	string
	Author	string
	Quote	string
}

type Book struct {
	Citacao []Quote
}