package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/devbdaily/sigil-solver/pkg/solver"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./website/public")))
	http.HandleFunc("/solve", solveHandler)
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}

func solveHandler(w http.ResponseWriter, req *http.Request) {
	input := SolveParams{}
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&input)

	pieces := []solver.Piece{}
	for _, piece := range input.Pieces {
		pieces = append(pieces, solver.Piece{
			Definition: piece.Blocks,
		})
	}

	s := solver.NewSolver(&input.Board, pieces)

	result, err := s.Solve()

	encoder := json.NewEncoder(w)

	if err != nil {
		res := ErrorResponse{
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		encoder.Encode(res)
		return
	}

	res := SolveResponse{}

	for _, piece := range result.Placed {
		res.Pieces = append(res.Pieces, SimplePiece{
			Blocks: piece.Blocks(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(res)
}

// SolveParams are the parameters provided in the request for the API: POST "/solve"
type SolveParams struct {
	Board  solver.Board
	Pieces []SimplePiece `json:"pieces"`
}

// SolveResponse is the response given in the API: POST "/solve"
type SolveResponse struct {
	Pieces []SimplePiece `json:"pieces"`
}

// SimplePiece is just a simplified version of the Piece struct for use in the API.
type SimplePiece struct {
	Blocks []solver.Location `json:"blocks"`
}

// ErrorResponse is the response given by the API when there is an error.
type ErrorResponse struct {
	Message string `json:"error"`
}
