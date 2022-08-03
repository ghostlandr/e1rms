package e1rm_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"e1rms/internal/e1rm"
	e1rm_calc "e1rms/internal/e1rm/calc"
)

type e1rmResponse struct {
	E1RM float64 `json:"e1rm"`
}

type e1rmHandler struct {
	s e1rm.E1RMService
}

/*
  weight: number;
  rpe: number;
  reps: number;
  lift: string;
  e1rm: string;
  created?: Date;
*/
type ApiE1RMCalculation struct {
	TotalWeight  float64 `json:"weight"`
	RPE          float64 `json:"rpe"`
	Reps         int16   `json:"reps"`
	Lift         string  `json:"lift"`
	Estimated1RM float64 `json:"e1rm"`
	CreatedAt    string  `json:"created"`
}

func New(s e1rm.E1RMService) e1rm.E1RMHandler {
	return &e1rmHandler{s}
}

func (e *e1rmHandler) ServeE1rmRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	q := r.URL.Query()
	totalWeight := q.Get("totalWeight")
	reps := q.Get("reps")
	rpe := q.Get("rpe")
	lift := q.Get("lift")
	if totalWeight == "" || reps == "" || rpe == "" {
		http.Error(w, "bad request: totalWeight, reps, and rpe must be provided as query parameters", http.StatusBadRequest)
		return
	}

	result, err := e.s.CalculateE1RM(ctx, totalWeight, rpe, reps, lift)

	if err != nil {
		http.Error(w, fmt.Sprintf("error calculating E1RM: %s", err), http.StatusBadRequest)
		return
	}

	resp := e1rmResponse{
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

func (e *e1rmHandler) ServeListE1rmRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e1rms, err := e.s.ListE1RMs(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing e1rms: %v", err), http.StatusInternalServerError)
		return
	}

	apiE1rms := e1rmsToApiE1rms(e1rms)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err = json.NewEncoder(w).Encode(apiE1rms)
	if err != nil {
		http.Error(w, fmt.Sprintf("not sure what happened: %s", err), 400)
	}
}

func e1rmsToApiE1rms(e1rms []*e1rm_calc.E1RMCalculation) []ApiE1RMCalculation {
	apiE1rms := make([]ApiE1RMCalculation, len(e1rms))
	for idx, e1rm := range e1rms {
		lift := e1rm.Lift
		if lift == "" {
			lift = "Unspecified"
		}
		apiE1rms[idx] = ApiE1RMCalculation{
			TotalWeight:  e1rm.TotalWeight,
			RPE:          e1rm.RPE,
			Reps:         e1rm.Reps,
			Estimated1RM: e1rm.E1RM,
			CreatedAt:    e1rm.CreatedAt,
			Lift:         lift,
		}
	}
	return apiE1rms
}
