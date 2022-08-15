package webserver

import (
	"context"
	"genxoft.dev/internal/container"
	"genxoft.dev/internal/webserver/handler"
	"genxoft.dev/internal/webserver/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Webserver struct {
	mode      string
	services  container.Services
	webFiles  string
	version   string
	releaseId string
}

type ctxKey int8

const (
	ctxKeyRequestId ctxKey = iota
)

func New(m string, s container.Services, w string, v string, r string) *Webserver {
	return &Webserver{
		m,
		s,
		w,
		v,
		r,
	}
}

func (s *Webserver) Start(listen string) error {

	r := mux.NewRouter()

	r.Use(s.setRequestID)
	r.Use(s.logRequest)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.CORS(s.mode, []string{"GET", "POST"}))

	api.Handle("/health", &handler.HealthHandler{
		Version:   s.version,
		ReleaseId: s.releaseId,
	})

	api.Handle("/fcm-token", handler.FcmTokenHandler(s.services.GetStore().FcmToken()))

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(s.webFiles))))

	srv := &http.Server{
		Handler:      r,
		Addr:         listen,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	return srv.ListenAndServe()
}

func (s *Webserver) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestId, id)))
	})
}

func (s *Webserver) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.services.GetLogger().WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestId),
			"headers":     r.Header,
		})
		logger.Infof("started %s %s \n", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}
