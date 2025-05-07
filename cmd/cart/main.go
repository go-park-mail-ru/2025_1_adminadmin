package main

import (
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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil { //log
		os.Exit(1)
	}
}

func run() (err error) {

	CartRepoPg, err := cartPgRepo.NewRestaurantRepository()
	if err != nil {
		return
	}
	cartRepoRedis, err := cartRedisRepo.NewCartRepository()
	if err != nil {
		return
	}
	CartUsecase := cartUsecase.NewCartUsecase(cartRepoRedis, CartRepoPg)
	CartDelivery := grpcCart.CreateCartHandler(CartUsecase)

	grpcMetrics, err := metrics.NewGrpcMetrics("cart")
	if err != nil {
		return
	}
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
