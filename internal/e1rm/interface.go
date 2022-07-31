package e1rm

import "net/http"

type E1RMService interface {
	CalculateE1RM(string, string, string) (float64, error)
}

type E1RMHandler interface {
	ServeE1rmRequest(w http.ResponseWriter, r *http.Request)
}
