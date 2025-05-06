package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcCart "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc"
	generatedCart "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	cartPgRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
	cartRedisRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
	cartUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/usecase"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/metrics"
	mw "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/metrics"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	return client
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {

	db, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRES_CONN"))
	if err != nil {
		return
	}
	defer db.Close()

	//tlsCredentials, err := loadtls.LoadTLSCredentials(cfg.Grpc.NoteIP)
	//if err != nil {
	//	logger.Error(err.Error())
	//	return
	//}

	redisClient := initRedis()
	CartRepoPg := cartPgRepo.NewRestaurantRepository(db)
	cartRepoRedis := cartRedisRepo.NewCartRepository(redisClient)
	CartUsecase := cartUsecase.NewCartUsecase(cartRepoRedis, CartRepoPg)
	CartDelivery := grpcCart.CreateCartHandler(CartUsecase)

	grpcMetrics, _ := metrics.NewGrpcMetrics("cart")
	grpcMiddleware := mw.NewGrpcMw(grpcMetrics)

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(grpcMiddleware.UnaryServerInterceptor()))
	generatedCart.RegisterCartServiceServer(gRPCServer, CartDelivery)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "5460"))
		if err != nil {
			return
		}
		if err := gRPCServer.Serve(listener); err != nil {
			return
		}
	}()

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.PathPrefix("/metrics").Handler(promhttp.Handler())
	http.Handle("/", r)
	httpSrv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "5461")}
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	gRPCServer.GracefulStop()
	return nil
}
