package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generatedSurvey "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/delivery/grpc/gen/proto"
	"github.com/gorilla/mux"
)

// Реализация заглушек сервиса Stat
type stubStatServer struct {
	generatedSurvey.UnimplementedStatServer
}

func (s *stubStatServer) GetSurvey(ctx context.Context, req *generatedSurvey.GetSurveyRequest) (*generatedSurvey.GetSurveyResponse, error) {
	// Заглушка: вернём пустой список вопросов
	return &generatedSurvey.GetSurveyResponse{
		Questions: []*generatedSurvey.Question{},
	}, nil
}

func (s *stubStatServer) Vote(ctx context.Context, req *generatedSurvey.VoteRequest) (*generatedSurvey.VoteResponse, error) {
	// Заглушка: просто принимаем голос без обработки
	return &generatedSurvey.VoteResponse{}, nil
}

func (s *stubStatServer) CreateSurvey(ctx context.Context, req *generatedSurvey.CreateSurveyRequest) (*generatedSurvey.CreateSurveyResponse, error) {
	// Заглушка: просто принимаем создание опроса
	return &generatedSurvey.CreateSurveyResponse{}, nil
}

func (s *stubStatServer) GetStats(ctx context.Context, req *generatedSurvey.GetStatsRequest) (*generatedSurvey.GetStatsResponse, error) {
	// Заглушка: вернём пустую статистику
	return &generatedSurvey.GetStatsResponse{
		Stats: []*generatedSurvey.StatModel{},
	}, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("main exited with error:", err)
		os.Exit(1)
	}
}

func run() error {
	// Настройка логирования
	logFile, err := os.OpenFile(os.Getenv("STAT_LOG_FILE"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error opening log file:", err)
		return err
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logFile, os.Stdout), &slog.HandlerOptions{Level: slog.LevelInfo}))

	// Настройка gRPC сервера
	port := os.Getenv("STAT_PORT")
	if port == "" {
		port = "9090" // дефолт
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Error("failed to listen", slog.String("error", err.Error()))
		return err
	}

	server := grpc.NewServer()
	generatedSurvey.RegisterStatServer(server, &stubStatServer{})

	// Запуск gRPC сервера в отдельной горутине
	go func() {
		logger.Info("gRPC server is listening", slog.String("port", port))
		if err := server.Serve(listener); err != nil {
			logger.Error("gRPC server error", slog.String("error", err.Error()))
		}
	}()

	// Запуск HTTP сервера
	httpServer := &http.Server{
		Addr:    ":5459",
		Handler: setupRouter(),
	}

	go func() {
		logger.Info("HTTP server is listening", slog.String("port", "5459"))
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Error("HTTP server error", slog.String("error", err.Error()))
		}
	}()

	// Ожидание сигнала остановки
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Info("Shutting down servers...")
	server.GracefulStop()
	httpServer.Shutdown(context.Background())
	return nil
}

// Функция для настройки маршрутов HTTP сервера
func setupRouter() *mux.Router {
	r := mux.NewRouter()

	// HTTP роуты
	r.HandleFunc("/api/survey", GetSurveyHandler).Methods("GET")
	r.HandleFunc("/api/vote", VoteHandler).Methods("POST")
	r.HandleFunc("/api/survey", CreateSurveyHandler).Methods("POST")
	r.HandleFunc("/api/stats", GetStatsHandler).Methods("GET")

	return r
}

// Реализация обработчиков HTTP запросов
func GetSurveyHandler(w http.ResponseWriter, r *http.Request) {
	// Создание контекста и вызов gRPC сервиса
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := generatedSurvey.NewStatClient(conn)
	resp, err := client.GetSurvey(context.Background(), &generatedSurvey.GetSurveyRequest{})
	if err != nil {
		http.Error(w, "failed to get survey", http.StatusInternalServerError)
		return
	}

	// Возврат ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", resp)))
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	// Пример обработки голосования
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := generatedSurvey.NewStatClient(conn)
	req := &generatedSurvey.VoteRequest{
		QuestionId: "question1",
		Vote:       1,
	}
	_, err = client.Vote(context.Background(), req)
	if err != nil {
		http.Error(w, "failed to submit vote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Vote submitted successfully"))
}

func CreateSurveyHandler(w http.ResponseWriter, r *http.Request) {
	// Пример создания опроса
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := generatedSurvey.NewStatClient(conn)
	req := &generatedSurvey.CreateSurveyRequest{
		Questions: []*generatedSurvey.CreateQuestionRequest{
			{
				Title:        "What is your favorite color?",
				MinMark:      1,
				Skip:         0,
				QuestionType: "single_choice",
			},
		},
	}
	_, err = client.CreateSurvey(context.Background(), req)
	if err != nil {
		http.Error(w, "failed to create survey", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Survey created successfully"))
}

func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	// Пример получения статистики
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := generatedSurvey.NewStatClient(conn)
	resp, err := client.GetStats(context.Background(), &generatedSurvey.GetStatsRequest{})
	if err != nil {
		http.Error(w, "failed to get stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", resp)))
}
