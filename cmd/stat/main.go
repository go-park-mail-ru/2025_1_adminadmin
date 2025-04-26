package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generatedSurvey "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/delivery/grpc/gen"
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

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.Header.Get("Content-Type") != "" && r.Header.Get("Content-Type")[:4] == "grpc" {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func run() error {
	gr := grpc.NewServer()
	generatedSurvey.RegisterStatServer(gr, &stubStatServer{})
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "5459"))
	if err != nil {
		return err
	}
	if err := gr.Serve(listener); err != nil {
		return err
	}

	return nil
}

func grpcClient() (generatedSurvey.StatClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:5459", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return generatedSurvey.NewStatClient(conn), conn, nil
}
