package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	State          string      `json:"state"`
	Board          ClientBoard `json:"board"`
	Recommendation Coordinate  `json:"recommendation"`
}

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	str := values.Get("size")

	size, err := strconv.Atoi(str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	board := NewBoard(size)
	clientBoard := NewClientBoard(board)

	json.NewEncoder(w).Encode(clientBoard)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var clientBoard ClientBoard

	stateToStringMap := make(map[State]string)
	stateToStringMap[DRAW] = "DRAW"
	stateToStringMap[EMPTY] = " "
	stateToStringMap[X] = "X"
	stateToStringMap[O] = "O"

	err := json.NewDecoder(r.Body).Decode(&clientBoard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	board := NewBoardFromClient(clientBoard)

	hash := board.Hash()

	if response, exists := GetCache(hash); exists {
		json.NewEncoder(w).Encode(response)
		return
	}

	state := board.DetermineState()

	var recommendation Coordinate
	if state == EMPTY {
		_, x, y := board.MaxAlphaBeta(-2, 2)
		board[x][y] = O

		_, x, y = board.MinAlphaBeta(-2, 2)
		recommendation = Coordinate{x, y}
	}

	state = board.DetermineState()
	result := stateToStringMap[state]

	clientBoard = NewClientBoard(board)

	response := Response{result, clientBoard, recommendation}
	json.NewEncoder(w).Encode(response)

	SetCache(hash, response)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
