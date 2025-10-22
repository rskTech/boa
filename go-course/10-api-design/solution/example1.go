package main
import (
    "encoding/json"
    "log"
    "net/http"
)
type Message struct{ Text string `json:"text"` }
func hello(w http.ResponseWriter, r *http.Request){ json.NewEncoder(w).Encode(Message{"hello"}) }
func main(){ http.HandleFunc("/hello", hello); log.Println("Listening :8080"); log.Fatal(http.ListenAndServe(":8080", nil)) }
