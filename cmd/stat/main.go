package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generatedSurvey "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/delivery/grpc/gen/proto"
	"github.com/gorilla/mux"
)

type stubStatServer struct {
	generatedSurvey.UnimplementedStatServer
}

func (s *stubStatServer) GetSurvey(ctx context.Context, req *generatedSurvey.GetSurveyRequest) (*generatedSurvey.GetSurveyResponse, error) {
	return &generatedSurvey.GetSurveyResponse{Questions: []*generatedSurvey.Question{}}, nil
}

func (s *stubStatServer) Vote(ctx context.Context, req *generatedSurvey.VoteRequest) (*generatedSurvey.VoteResponse, error) {
	return &generatedSurvey.VoteResponse{}, nil
}

func (s *stubStatServer) CreateSurvey(ctx context.Context, req *generatedSurvey.CreateSurveyRequest) (*generatedSurvey.CreateSurveyResponse, error) {
	return &generatedSurvey.CreateSurveyResponse{}, nil
}

func (s *stubStatServer) GetStats(ctx context.Context, req *generatedSurvey.GetStatsRequest) (*generatedSurvey.GetStatsResponse, error) {
	return &generatedSurvey.GetStatsResponse{Stats: []*generatedSurvey.StatModel{}}, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("main exited with error:", err)
		os.Exit(1)
	}
}

func run() error {
	grpcPort := "9090"
	httpPort := "5459"

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	generatedSurvey.RegisterStatServer(grpcServer, &stubStatServer{})

	go func() {
		grpcServer.Serve(lis)
	}()

	httpServer := &http.Server{
		Addr:    ":" + httpPort,
		Handler: setupRouter(),
	}

	go func() {
		httpServer.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	grpcServer.GracefulStop()
	httpServer.Shutdown(context.Background())
	return nil
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/survey", GetSurveyHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/vote", VoteHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/survey", CreateSurveyHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/stats", GetStatsHandler).Methods(http.MethodGet)
	return r
}

func grpcClient() (generatedSurvey.StatClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return generatedSurvey.NewStatClient(conn), conn, nil
}

func GetSurveyHandler(w http.ResponseWriter, r *http.Request) {
	client, conn, err := grpcClient()
	if err != nil {
		http.Error(w, "gRPC connection failed", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	resp, err := client.GetSurvey(r.Context(), &generatedSurvey.GetSurveyRequest{})
	if err != nil {
		http.Error(w, "gRPC GetSurvey failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", resp)
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	client, conn, err := grpcClient()
	if err != nil {
		http.Error(w, "gRPC connection failed", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	_, err = client.Vote(r.Context(), &generatedSurvey.VoteRequest{
		QuestionId: "q1",
		Vote:       1,
	})
	if err != nil {
		http.Error(w, "gRPC Vote failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("vote submitted"))
}

func CreateSurveyHandler(w http.ResponseWriter, r *http.Request) {
	client, conn, err := grpcClient()
	if err != nil {
		http.Error(w, "gRPC connection failed", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	_, err = client.CreateSurvey(r.Context(), &generatedSurvey.CreateSurveyRequest{
		Questions: []*generatedSurvey.CreateQuestionRequest{
			{Title: "Fav color?", MinMark: 1, Skip: 0, QuestionType: "single_choice"},
		},
	})
	if err != nil {
		http.Error(w, "gRPC CreateSurvey failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("survey created"))
}

func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	client, conn, err := grpcClient()
	if err != nil {
		http.Error(w, "gRPC connection failed", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	resp, err := client.GetStats(r.Context(), &generatedSurvey.GetStatsRequest{})
	if err != nil {
		http.Error(w, "gRPC GetStats failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", resp)
}
