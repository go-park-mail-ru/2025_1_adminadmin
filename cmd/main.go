package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
)

func initDB() (*pgxpool.Pool, error) {
	connStr := "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") +
		"@" + os.Getenv("POSTGRES_HOST") + ":" + "5432" + "/" + os.Getenv("POSTGRES_DB") +
		"?sslmode=disable"

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	log.Println("Успешное подключение к PostgreSQL")
	return pool, nil
}

func main() {
	pool, err := initDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к PostgreSQL: %v", err)
	}
	defer pool.Close()

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusTeapot)
	})

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", handlers.SignUp).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signin", handlers.SignIn).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/check", handlers.Check).Methods(http.MethodGet, http.MethodOptions)
	}
	restaurants := r.PathPrefix("/restaurants").Subrouter()
	{
		restaurants.HandleFunc("/list", handlers.RestaurantList).Methods(http.MethodGet, http.MethodOptions)
		restaurants.HandleFunc("/{id}", handlers.RestaurantByID).Methods(http.MethodGet, http.MethodOptions)

	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	http.Handle("/", http.StripPrefix("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
		c.Handler(r).ServeHTTP(w, r)
	})))

	http.Handle("/", r)
	srv := http.Server{
		Handler:           c.Handler(r),
		Addr:              ":5458",
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Получен сигнал остановки")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Printf("Ошибка при остановке сервера: %v", err)
	} else {
		log.Println("Сервер успешно остановлен")
	}
}
