package server

import (
	"api/router"
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ApiServer struct {
	Server    *http.Server
	AssetPath string
}

func NewApiServer(port, assets string) ApiServer {
	return ApiServer{
		Server: &http.Server{
			Handler: NewRouteHandler(),
			Addr:    port,
		},
		AssetPath: assets,
	}

}

func NewRouteHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", router.WebApp)
	r.HandleFunc("/assets/{rest:.*}", router.WebAssets)
	r.HandleFunc("/api/v1/{rest:.*}", router.RouterV1)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)
	return handler
}

func Start(api ApiServer) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Listening on port " + api.Server.Addr[1:])
		if err := api.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Server error:", err)
		}
	}()
	return &wg
}

func Shutdown(api ApiServer) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := api.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed:", err)
	}

}
