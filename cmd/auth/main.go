package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcAuth "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc"
	generatedAuth "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	authRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/metrics"
	mw "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/metrics"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

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

	AuthRepo := authRepo.CreateAuthRepo(db)
	AuthUsecase := authUsecase.CreateAuthUsecase(AuthRepo)
	AuthDelivery := grpcAuth.CreateAuthHandler(AuthUsecase)

	grpcMetrics, _ := metrics.NewGrpcMetrics("auth")
	grpcMiddleware := mw.NewGrpcMw(grpcMetrics)

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(grpcMiddleware.UnaryServerInterceptor()))
	generatedAuth.RegisterAuthServiceServer(gRPCServer, AuthDelivery)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "5459"))
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
	httpSrv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "5462")}
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
