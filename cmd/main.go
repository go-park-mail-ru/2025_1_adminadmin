package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusTeapot)
	})
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", handlers.SignUp).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signin", handlers.SignIn).Methods(http.MethodPost, http.MethodOptions)
	}
	restaurants := r.PathPrefix("/restaurants").Subrouter()
	{
		restaurants.HandleFunc("/list", handlers.RestaurantList).Methods(http.MethodGet, http.MethodOptions)
	}
	http.Handle("/", r)
	srv := http.Server{
		Handler:           r,
		Addr:              ":5458",
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}
