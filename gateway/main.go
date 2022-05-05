package main

import (
	"context"
	"net/http"
	"os"

	pb "github.com/ekkinox/grpc-demo/gateway/proto/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Str("component", "text-tool-gateway").Logger()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := "server:50051"
	port := ":8888"

	// Register TextTool Service
	logger.Info().Msgf("Registering TextTool Service %s", endpoint)
	err := pb.RegisterTextToolsServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		logger.Fatal().Err(err)
	}

	logger.Info().Msgf("TextTool endpoint created, listening on %s", port)
	handler := http.Handler(mux)
	err = http.ListenAndServe(port, handler)
	if err != nil {
		logrus.Fatal(err)
	}
}
