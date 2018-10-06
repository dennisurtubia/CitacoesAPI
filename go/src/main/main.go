package main

import(
	"quote"
    "net/http"
    "log"
	"encoding/json"
)
func getquote(w http.ResponseWriter, r *http.Request) {
	randomQuote := quote.GetRandomQuote()
	json.NewEncoder(w).Encode(randomQuote)
}

func main() {
	http.HandleFunc("/random", getquote)
    log.Fatal(http.ListenAndServe(":8080", nil))
}