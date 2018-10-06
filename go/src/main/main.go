package main

import(
	"quote"
    "fmt"
    "net/http"
    "log"
)
func getquote(w http.ResponseWriter, r *http.Request) {
	randomQuote := quote.GetRandomQuote()
    fmt.Fprintf(w,"\"%s\" (%s)", randomQuote.Quote, randomQuote.Author)
}

func main() {
	http.HandleFunc("/random", getquote)
    log.Fatal(http.ListenAndServe(":8080", nil))
}