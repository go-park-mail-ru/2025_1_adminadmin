package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcAuth "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc"
	generatedAuth "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	authRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
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

	gRPCServer := grpc.NewServer()
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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	gRPCServer.GracefulStop()
	return nil
}
