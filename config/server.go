package config

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Abeldlp/architectured-go/service"
)

type Server struct {
	svc service.Service
}

func NewApiServer(svc service.Service) *Server {
	return &Server{
		svc: svc,
	}
}

func (s *Server) Start(listedAdr string) error {
	http.HandleFunc("/", s.handleGetFact)
	return http.ListenAndServe(listedAdr, nil)
}

func (s *Server) handleGetFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
