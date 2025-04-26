package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcSurvey "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/delivery/grpc"
	generatedSurvey "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/delivery/grpc/gen"
	surveyRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/repo"
	surveyUsecase "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {

	db, err := pgxpool.Connect(context.Background(), "postgres://admin_test:Admin_test2025$@postgres:5432/adminadmin_v1?sslmode=disable")
	if err != nil {
		return
	}
	defer db.Close()

	//tlsCredentials, err := loadtls.LoadTLSCredentials(cfg.Grpc.NoteIP)
	//if err != nil {
	//	logger.Error(err.Error())
	//	return
	//}

	SurveyRepo := surveyRepo.CreateSurveyRepo(db)
	SurveyUsecase := surveyUsecase.CreateSurveyUsecase(SurveyRepo)
	SurveyDelivery := grpcSurvey.NewGrpcSurveyHandler(SurveyUsecase)

	gRPCServer := grpc.NewServer()
	generatedSurvey.RegisterStatServer(gRPCServer, SurveyDelivery)

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
