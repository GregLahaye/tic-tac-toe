package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	switch r.Method {
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var clientBoard ClientBoard

	err := json.NewDecoder(r.Body).Decode(&clientBoard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	board := NewBoard(clientBoard)

	_, x, y := board.MaxAlphaBeta(-2, 2)
	board[x][y] = O

	clientBoard = NewClientBoard(board)
	json.NewEncoder(w).Encode(clientBoard)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
