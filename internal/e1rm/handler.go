package e1rm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type E1rmHandler interface {
	ServeE1rmRequest(w http.ResponseWriter, r *http.Request)
}

type E1rmResponse struct {
	E1RM float64 `json:"e1rm"`
}

type e1rmHandler struct{}

func NewHandler() E1rmHandler {
	return &e1rmHandler{}
}

func (e *e1rmHandler) ServeE1rmRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	totalWeight := q.Get("totalWeight")
	reps := q.Get("reps")
	rpe := q.Get("rpe")
	if totalWeight == "" || reps == "" || rpe == "" {
		http.Error(w, "bad request: totalWeight, reps, and rpe must be provided as query parameters", http.StatusBadRequest)
		return
	}

	result, err := CalculateE1RMFromStrings(totalWeight, rpe, reps)

	if err != nil {
		http.Error(w, fmt.Sprintf("error calculating E1RM: %s", err), http.StatusBadRequest)
		return
	}

	resp := E1rmResponse{
		E1RM: result,
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("not sure what happened: %s", err), 400)
	}
}
