package api

import (
    "log"
    "net/http"

    "github.com/Aspandiyar933/Ilovedogs/store"
    "github.com/Aspandiyar933/Ilovedogs/tasks"
    "github.com/Aspandiyar933/Ilovedogs/users"
    "github.com/gorilla/mux"
)

type APIServer struct {
    addr  string
    store store.Store
}

func NewAPIServer(addr string, store store.Store) *APIServer {
    return &APIServer{
        addr:  addr,
        store: store,
    }
}

func (s *APIServer) Serve() {
    // Create a new router
    router := mux.NewRouter()

    subrouter := router.PathPrefix("/api/v1").Subrouter()

    userService := users.NewUserService(s.store)
    userService.RegisterRoutes(subrouter)

    postService := tasks.NewPostService(s.store)
    postService.RegisterRoutes(subrouter)

    log.Printf("API server listening on %s", s.addr)
    if err := http.ListenAndServe(s.addr, router); err != nil {
        log.Fatalf("Failed to start the API server: %v", err)
    }
}
