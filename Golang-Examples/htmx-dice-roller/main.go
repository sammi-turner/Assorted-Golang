package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func rollDice(n int, sides int) int {
	rand.Seed(time.Now().UnixNano())

	sum := 0
	for i := 0; i < n; i++ {
		sum += rand.Intn(sides) + 1
	}

	return sum
}

func diceRollHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(r.Form.Get("n"))
	if err != nil {
		http.Error(w, "Invalid number of dice", http.StatusBadRequest)
		return
	}

	sides, err := strconv.Atoi(r.Form.Get("sides"))
	if err != nil {
		http.Error(w, "Invalid number of sides", http.StatusBadRequest)
		return
	}

	roll := rollDice(n, sides)

	fmt.Fprintf(w, "<p>You rolled %d</p>", roll)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/roll", diceRollHandler)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
