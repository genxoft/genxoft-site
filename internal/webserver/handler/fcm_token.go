package handler

import (
	"encoding/json"
	"genxoft.dev/internal/model"
	"genxoft.dev/internal/store"
	"io/ioutil"
	"net/http"
)

type TokenHandler struct {
	r store.FcmRepository
}

func FcmTokenHandler(r store.FcmRepository) http.Handler {
	return &TokenHandler{
		r,
	}
}

func (h *TokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if r.Body == nil {
		http.Error(w, "Body parse error", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Body parse error", 400)
		return
	}
	var fc model.FcmSettings
	if err = json.Unmarshal(body, &fc); err != nil {
		http.Error(w, "Body parse error", 400)
		return
	}

	fc.IP = r.RemoteAddr

	if err := fc.Validate(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	f, err := h.r.Find(fc.Token)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	if f != nil {
		if err := h.r.Update(&fc); err != nil {
			http.Error(w, "Internal server error", 500)
		}
	} else {
		if err := h.r.Create(&fc); err != nil {
			http.Error(w, "Internal server error", 500)
		}
	}

	w.WriteHeader(200)
}
