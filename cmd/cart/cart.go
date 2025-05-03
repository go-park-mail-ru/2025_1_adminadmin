package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcCart "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc"
	generatedCart "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	cartPgRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
	cartRedisRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
	cartUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
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

	gRPCServer := grpc.NewServer()
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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	gRPCServer.GracefulStop()
	return nil
}
