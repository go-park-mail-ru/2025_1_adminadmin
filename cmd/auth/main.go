package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
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
	metricsmw "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/middleware/metrics"
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
	logFile, err := os.OpenFile(os.Getenv("AUTH_LOG_FILE"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error opening log file: " + err.Error())
		return
	}
	defer logFile.Close()
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), &slog.HandlerOptions{Level: slog.LevelInfo}))

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

	grpcMetrics, err := metrics.NewGrpcMetrics("auth")
	if err != nil {
		logger.Error("can't create metrics")
	}
	metricsMw := metricsmw.NewGrpcMw(*grpcMetrics)
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(metricsMw.ServerMetricsInterceptor))
	generatedAuth.RegisterAuthServiceServer(gRPCServer, AuthDelivery)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.PathPrefix("/metrics").Handler(promhttp.Handler())
	http.Handle("/", r)
	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "5459"))
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
