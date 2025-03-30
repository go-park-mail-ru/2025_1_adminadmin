package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/middleware"
	authHandler "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/http"
	authRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	restaurantDelivery "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/delivery/http"
	restaurantRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/repo"
	restaurantUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/usecase"
	cartRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo"
	cartHandler "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/http"
	cartUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/usecase"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
)

func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	return client
}

func initDB() (*pgxpool.Pool, error) {
	connStr := os.Getenv("POSTGRES_CONN")

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

	redisClient := initRedis()
	cartRepo := cartRepo.NewCartRepository(redisClient)
	cartUsecase := cartUsecase.NewCartUsecase(cartRepo)
	cartHandler := cartHandler.NewCartHandler(cartUsecase)

	authRepo := authRepo.CreateAuthRepo(pool)
	authUsecase := authUsecase.CreateAuthUsecase(authRepo)
	authHandler := authHandler.CreateAuthHandler(authUsecase)

	restaurantRepo := restaurantRepo.NewRestaurantRepository(pool)
	restaurantUsecase := restaurantUsecase.NewRestaurantsUsecase(restaurantRepo)
	restaurantDelivery := restaurantDelivery.NewRestaurantHandler(restaurantUsecase)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Не найдено", http.StatusTeapot)
	})

	r.Use(middleware.CorsMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signin", authHandler.SignIn).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/check", authHandler.Check).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/logout", authHandler.LogOut).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/update_user", authHandler.UpdateUser).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/update_userpic", authHandler.UpdateUserPic).Methods(http.MethodPost, http.MethodOptions)
	}
	restaurants := r.PathPrefix("/restaurants").Subrouter()
	{
		restaurants.HandleFunc("/list", restaurantDelivery.RestaurantList).Methods(http.MethodGet, http.MethodOptions)
		restaurants.HandleFunc("/{id}", restaurantDelivery.RestaurantById).Methods(http.MethodGet, http.MethodOptions)
		restaurants.HandleFunc("/{id}/products", restaurantDelivery.GetProductsByRestaurant).Methods(http.MethodGet, http.MethodOptions)
	}
	cart := r.PathPrefix("/cart").Subrouter()
	{
		cart.HandleFunc("", cartHandler.GetCart).Methods(http.MethodGet, http.MethodOptions)
		cart.HandleFunc("/add/{productID}", cartHandler.AddToCart).Methods(http.MethodGet, http.MethodOptions)
		cart.HandleFunc("/remove/{productID}", cartHandler.RemoveFromCart).Methods(http.MethodGet, http.MethodOptions)
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
