package handler

import (
	"encoding/json"
	"genxoft.dev/internal/model"
	"log"
	"net/http"
)

type HealthHandler struct {
	Version   string
	ReleaseId string
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	m := model.Health{
		Status:    "pass",
		Version:   h.Version,
		ReleaseId: h.ReleaseId,
	}

	body, err := json.Marshal(m)
	if err != nil {
		log.Panicln(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		log.Panicln(err.Error())
	}
}
