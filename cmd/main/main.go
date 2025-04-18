package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	authHandler "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/http"
	authRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	cartHandler "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/http"
	cartPgRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
	cartRedisRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
	cartUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/usecase"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/cors"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/log"
	restaurantDelivery "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/delivery/http"
	restaurantRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/repo"
	restaurantUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/usecase"
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

func initDB(logger *slog.Logger) (*pgxpool.Pool, error) {
	connStr := os.Getenv("POSTGRES_CONN")

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	logger.Info("Успешное подключение к PostgreSQL")
	return pool, nil
}

// @title AdminAdmin API
// @version 1.0
// @description API для проекта DoorDashers.
// @host localhost:5458
// @BasePath /api
func main() {
	logFile, err := os.OpenFile(os.Getenv("MAIN_LOG_FILE"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error opening log file: " + err.Error())
		return
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), &slog.HandlerOptions{Level: slog.LevelInfo}))

	pool, err := initDB(logger)
	if err != nil {
		logger.Error("Ошибка при подключении к PostgreSQL: " + err.Error())
	}
	defer pool.Close()

	redisClient := initRedis()
	cartRepoPg := cartPgRepo.NewRestaurantRepository(pool)
	cartRepoRedis := cartRedisRepo.NewCartRepository(redisClient)
	cartUsecase := cartUsecase.NewCartUsecase(cartRepoRedis, cartRepoPg)
	cartHandler := cartHandler.NewCartHandler(cartUsecase)

	logMW := log.CreateLoggerMiddleware(logger)

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

	r.Use(
		logMW,
		cors.CorsMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signin", authHandler.SignIn).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/check", authHandler.Check).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/logout", authHandler.LogOut).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/update_user", authHandler.UpdateUser).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/update_userpic", authHandler.UpdateUserPic).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/address", authHandler.GetUserAddresses).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/address", authHandler.DeleteAddress).Methods(http.MethodDelete, http.MethodOptions)
		auth.HandleFunc("/address", authHandler.AddAddress).Methods(http.MethodPost, http.MethodOptions)
	}
	restaurants := r.PathPrefix("/restaurants").Subrouter()
	{
		restaurants.HandleFunc("/list", restaurantDelivery.RestaurantList).Methods(http.MethodGet, http.MethodOptions)
		restaurants.HandleFunc("/{id}", restaurantDelivery.GetProductsByRestaurant).Methods(http.MethodGet, http.MethodOptions)
	}
	cart := r.PathPrefix("/cart").Subrouter()
	{
		cart.HandleFunc("", cartHandler.GetCart).Methods(http.MethodGet, http.MethodOptions)
		cart.HandleFunc("/update/{productID}", cartHandler.UpdateQuantityInCart).Methods(http.MethodPost, http.MethodOptions)
		cart.HandleFunc("/clear", cartHandler.ClearCart).Methods(http.MethodPost, http.MethodOptions)
	}

	order := r.PathPrefix("/order").Subrouter()
	{
		order.HandleFunc("/create", cartHandler.CreateOrder).Methods(http.MethodPost, http.MethodOptions)
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
			logger.Error("Ошибка при запуске сервера: " + err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	logger.Info("Получен сигнал остановки")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		logger.Error("Ошибка при остановке сервера: " + err.Error())
	} else {
		logger.Info("Сервер успешно остановлен")
	}
}
