package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/decrement", decrementHandler)
	http.HandleFunc("/", homeHandler)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	response := fmt.Sprintf(`<div id="counter" hx-swap="outerHTML">%s</div>`, strconv.Itoa(counter))
	w.Write([]byte(response))
}

func decrementHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter--
	mutex.Unlock()
	response := fmt.Sprintf(`<div id="counter" hx-swap="outerHTML">%s</div>`, strconv.Itoa(counter))
	w.Write([]byte(response))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
