package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
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
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/satori/uuid"
	"google.golang.org/grpc"
)

var categoryWeights = map[string]float64{
	"Соусы":                0.05,
	"Закуски":              0.10,
	"Десерты":              0.08,
	"Выпечка/Хлеб":         0.05,
	"Горячие блюда":        0.15,
	"Паста":                0.10,
	"Блюда с рисом":        0.10,
	"Салаты":               0.07,
	"Напитки":              0.15,
	"Вегетарианские блюда": 0.05,
	"Мясные блюда":         0.10,
	"Карри":                0.05,
	"Рыбные блюда":         0.05,
	"Аперитивы/Миксы":      0.03,
	"Другое":               0.02,
	"Прочее":               0.02,
}

type CartItem struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	ImageURL string    `json:"image_url"`
	Weight   int       `json:"weight"`
	Amount   int       `json:"amount"`
}

// easyjson:json
type Cart struct {
	Id        uuid.UUID  `json:"restaurant_id"`
	Name      string     `json:"restaurant_name"`
	CartItems []CartItem `json:"products"`
	TotalSum  float64    `json:"total_sum"`
}

func chooseCategoryByWeight(categoryWeights map[string]float64, productGroups map[string][]Product) (string, bool) {
	var total float64
	categories := make([]string, 0)
	weights := make([]float64, 0)

	for cat, weight := range categoryWeights {
		if len(productGroups[cat]) > 0 { // только если есть товары
			total += weight
			categories = append(categories, cat)
			weights = append(weights, weight)
		}
	}

	if total == 0 {
		return "", false
	}

	randVal := rand.Float64() * total
	sum := 0.0

	for i, w := range weights {
		sum += w
		if sum >= randVal {
			return categories[i], true
		}
	}

	return categories[len(categories)-1], true // fallback
}

func groupProductsByCategory(products []Product) map[string][]Product {
	grouped := make(map[string][]Product)
	for _, p := range products {
		if _, exists := grouped[p.Category]; !exists {
			grouped[p.Category] = []Product{}
		}
		grouped[p.Category] = append(grouped[p.Category], p)
	}
	return grouped
}

func generateOrder(db *sql.DB, userID, addressID string, restaurant Restaurant, products []Product) error {
	productGroups := groupProductsByCategory(products)

	numItems := rand.Intn(5) + 2 // от 1 до 5 товаров
	var selectedProducts []CartItem
	var totalPrice float64

	for i := 0; i < numItems; i++ {
		category, ok := chooseCategoryByWeight(categoryWeights, productGroups)
		if !ok || len(productGroups[category]) == 0 {
			continue
		}

		items := productGroups[category]
		selected := items[rand.Intn(len(items))] // случайный товар из категории
		selectedProducts = append(selectedProducts, CartItem{
			Id:       selected.ID,
			Name:     selected.Name,
			Price:    selected.Price,
			ImageURL: selected.ImageURL,
			Weight:   selected.Weight,
			Amount:   1, // можно сделать >1 позже
		})
		totalPrice += selected.Price
	}

	if len(selectedProducts) == 0 {
		return nil // пропустить пустые заказы
	}

	cart := Cart{
		Id:        restaurant.ID,
		Name:      restaurant.Name,
		CartItems: selectedProducts,
		TotalSum:  totalPrice,
	}

	cartJSON, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	orderItems := make([]uuid.UUID, len(selectedProducts))
	for i, item := range selectedProducts {
		orderItems[i] = item.Id
	}

	_, err = db.Exec(`
        INSERT INTO orders (
            user_id, status, address_id, order_products, 
            apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door, final_price, order_items, created_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW())`,
		userID, "delivered", addressID, cartJSON,
		"123", "no", "A", "3", "leave at door", true, totalPrice, pq.Array(orderItems),
	)

	return err
}

type Product struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	ImageURL     string
	Weight       int
	Category     string
	RestaurantID uuid.UUID
}

type Restaurant struct {
	ID   uuid.UUID
	Name string
}

func getRestaurants(db *sql.DB) ([]Restaurant, error) {
	rows, err := db.Query("SELECT id, name FROM restaurants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []Restaurant
	for rows.Next() {
		var r Restaurant
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, r)
	}
	return restaurants, nil
}

func getProductsByRestaurant(db *sql.DB, restaurantID uuid.UUID) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price, image_url, weight, category FROM products WHERE restaurant_id = $1", restaurantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		p.RestaurantID = restaurantID
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.ImageURL, &p.Weight, &p.Category); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func getFirstUserID(db *sql.DB) (string, error) {
	var userID string
	err := db.QueryRow("SELECT id FROM users LIMIT 1").Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

// @title AdminAdmin API
// @version 1.0
// @description API для проекта DoorDashers.
// @host localhost:5458
// @BasePath /api
func main() {

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Получаем ID первого пользователя из БД
	userID, err := getFirstUserID(db)
	if err != nil {
		log.Fatal("Failed to get first user from DB:", err)
	}

	addressID := "existing-address-id-here" // или тоже можно взять из БД, если хочешь

	restaurants1, err := getRestaurants(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, restaurant := range restaurants1 {
		products, err := getProductsByRestaurant(db, restaurant.ID)
		if err != nil {
			log.Printf("failed to load products for %s: %v", restaurant.Name, err)
			continue
		}

		for i := 0; i < 1000; i++ {
			err = generateOrder(db, userID, addressID, restaurant, products)
			if err != nil {
				log.Printf("failed to insert order for %s: %v", restaurant.Name, err)
			}
		}

		fmt.Printf("✅ Generated 1000 orders for %s\n", restaurant.Name)
	}
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
