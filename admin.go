package gomb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type AdminServer struct {
	port   int
	router *chi.Mux
	store  *Store
}

func NewAdminServer(port int, s *Store) *AdminServer {
	r := chi.NewRouter()

	// Setup middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	a := &AdminServer{router: r, port: port, store: s}

	// Setup routes
	r.Get("/topics", http.HandlerFunc(a.handleGetTopics))
	r.Post("/topics", http.HandlerFunc(a.handleCreateTopic))
	r.Delete("/topics", http.HandlerFunc(a.handleDeleteTopic))
	return a
}

func (a *AdminServer) Start() {
	log.Printf("Starting admin server on port %d", a.port)
	addr := fmt.Sprintf(":%d", a.port)
	http.ListenAndServe(addr, a.router)
}

func (a *AdminServer) handleGetTopics(w http.ResponseWriter, r *http.Request) {
	topics := a.store.GetTopics()
	resp, err := json.Marshal(topics)
	if err != nil {
		w.Write([]byte("An error occured"))
	} else {
		w.Write(resp)
	}
}

func (a *AdminServer) handleCreateTopic(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name string `json:"name"`
	}{}
	json.NewDecoder(r.Body).Decode(&body)
	err := a.store.CreateTopic(body.Name)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(fmt.Sprintf("Topic %s created", body.Name)))
	}
}

func (a *AdminServer) handleDeleteTopic(w http.ResponseWriter, r *http.Request) {

}
