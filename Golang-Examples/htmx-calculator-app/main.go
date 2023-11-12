package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Couldn't parse form", http.StatusBadRequest)
		return
	}

	first, err1 := strconv.ParseFloat(r.FormValue("firstNumber"), 64)
	second, err2 := strconv.ParseFloat(r.FormValue("secondNumber"), 64)
	op := r.FormValue("operation")

	if err1 != nil || err2 != nil {
		http.Error(w, "Numbers must be valid", http.StatusBadRequest)
		return
	}

	var result float64
	switch op {
	case "add":
		result = first + second
	case "sub":
		result = first - second
	case "mul":
		result = first * second
	case "div":
		if second != 0 {
			result = first / second
		} else {
			http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%.2f", result)
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	http.Handle("/", http.FileServer(http.Dir("./")))
    fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
