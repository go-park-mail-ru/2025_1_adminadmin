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

	authGen "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	cartPgRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"

	authHandler "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/http"
	cartGen "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	cartHandler "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/http"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/hub"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/metrics"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/cors"
	logs "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/log"
	metricsmw "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/metrics"
	promocodeDelivery "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/promocode/delivery/http"
	promocodeRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/promocode/repo"
	promocodeUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/promocode/usecase"
	restaurantDelivery "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/delivery/http"
	restaurantRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/repo"
	restaurantUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/usecase"
	searchDelivery "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search/delivery/http"
	searchRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search/repo"
	searchUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search/usecase"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

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

	cartConn, err := grpc.Dial("cart:5460", grpc.WithInsecure())
	if err != nil {
		logger.Error("Ошибка подключения к gRPC Cart-сервису: " + err.Error())
		return
	}
	defer cartConn.Close()

	//CartRepoPg, err := cartPgRepo.NewRestaurantRepository()
	//if err != nil {
	//	return
	//}
	CartRepoPg, err := cartPgRepo.NewRestaurantRepository()
	if err != nil {
		return
	}
	hubNew := &hub.Hub{Repo: CartRepoPg}
	go hubNew.Run(context.Background())

	cartGRPCClient := cartGen.NewCartServiceClient(cartConn)
	cartHandler := cartHandler.NewCartHandler(cartGRPCClient, hubNew)

	Metrics, err := metrics.NewHttpMetrics("main")
	if err != nil {
		logger.Error("can't create metrics")
	}
	MetricsMiddleware := metricsmw.CreateHttpMetricsMiddleware(Metrics, logger)
	logMW := logs.CreateLoggerMiddleware(logger)

	conn, err := grpc.Dial("auth:5459", grpc.WithInsecure())
	if err != nil {
		logger.Error("Ошибка подключения к gRPC Auth-сервису: " + err.Error())
		return
	}
	defer conn.Close()

	authGRPCClient := authGen.NewAuthServiceClient(conn)

	authHandler := authHandler.CreateAuthHandler(authGRPCClient)

	restaurantRepo, err := restaurantRepo.NewRestaurantRepository()
	if err != nil {
		return
	}
	restaurantUsecase := restaurantUsecase.NewRestaurantsUsecase(restaurantRepo)
	restaurantDelivery := restaurantDelivery.NewRestaurantHandler(restaurantUsecase)

	promocodeRepo, err := promocodeRepo.NewPromocodeRepository()
	if err != nil {
		return
	}
	promocodeUsecase := promocodeUsecase.NewPromocodeUsecase(promocodeRepo)
	promocodeDelivery := promocodeDelivery.NewPromocodeHandler(promocodeUsecase)

	searchRep, err := searchRepo.NewSearchRepo()
	if err != nil {
		return
	}
	searchUsecase := searchUsecase.NewSearchUsecase(searchRep)
	searchDelivery := searchDelivery.NewSearchHandler(searchUsecase)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Не найдено", http.StatusTeapot)
	})

	r.Use(
		logMW,
		MetricsMiddleware,
		cors.CorsMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	{

		auth.HandleFunc("/signin", authHandler.SignIn).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/qr", authHandler.GetQRCode).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/check2fa", authHandler.CheckCode).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/disable2fa", authHandler.Disable2fa).Methods(http.MethodDelete, http.MethodOptions)
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
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
		restaurants.HandleFunc("/{id}/reviews", restaurantDelivery.ReviewsList).Methods(http.MethodGet, http.MethodOptions)
		restaurants.HandleFunc("/{id}/reviews", restaurantDelivery.CreateReview).Methods(http.MethodPost, http.MethodOptions)
		restaurants.HandleFunc("/{id}/check", restaurantDelivery.CheckReviews).Methods(http.MethodGet, http.MethodOptions)
		restaurants.HandleFunc("/{id}/search", searchDelivery.SearchProductsInRestaurant).Methods(http.MethodGet)
	}
	cart := r.PathPrefix("/cart").Subrouter()
	{
		cart.HandleFunc("", cartHandler.GetCart).Methods(http.MethodGet, http.MethodOptions)
		cart.HandleFunc("/ws", cartHandler.Subscribe)
		cart.HandleFunc("/update/{productID}", cartHandler.UpdateQuantityInCart).Methods(http.MethodPost, http.MethodOptions)
		cart.HandleFunc("/clear", cartHandler.ClearCart).Methods(http.MethodPost, http.MethodOptions)
	}

	order := r.PathPrefix("/order").Subrouter()
	{
		order.HandleFunc("", cartHandler.GetOrders).Methods(http.MethodGet)
		order.HandleFunc("/{orderID}", cartHandler.GetOrderById).Methods(http.MethodGet)
		order.HandleFunc("/{orderID}/update", cartHandler.UpdateOrderStatus).Methods(http.MethodPost)
		order.HandleFunc("/create", cartHandler.CreateOrder).Methods(http.MethodPost, http.MethodOptions)
	}

	promocodes := r.PathPrefix("/promocodes").Subrouter()
	{
		promocodes.HandleFunc("", promocodeDelivery.GetPromocodes).Methods(http.MethodGet)
		promocodes.HandleFunc("/check", promocodeDelivery.CheckPromocode).Methods(http.MethodPost)
	}

	search := r.PathPrefix("/search").Subrouter()
	{
		search.HandleFunc("", searchDelivery.SearchRestaurantWithProducts).Methods(http.MethodGet)
	}

	r.HandleFunc("/payment", cartHandler.UpdateOrderStatus).Methods(http.MethodPost)
	r.PathPrefix("/metrics").Handler(promhttp.Handler())
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