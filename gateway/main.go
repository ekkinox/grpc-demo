package main

import (
	"context"
	_ "embed"
	pb "github.com/ekkinox/grpc-demo/gateway/proto/go"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
	"strings"
)

//go:embed doc/text_tools.swagger.json
var content []byte

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(ctx context.Context, logger zerolog.Logger) (http.Handler, error) {
	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := "server:50051"
	// Register TextTool Service
	logger.Info().Msgf("Registering TextTool Service %s", endpoint)
	err := pb.RegisterTextToolsServiceHandlerFromEndpoint(ctx, mux, endpoint, dialOpts)
	if err != nil {
		logger.Fatal().Err(err)
	}
	return mux, nil
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

// Run starts a HTTP server and blocks forever if successful.
func Run(address string) error {
	logger := zerolog.New(os.Stderr).With().Timestamp().Str("component", "text-tool-gateway").Logger()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := http.NewServeMux()
	// Gateway server
	gw, err := newGateway(ctx, logger)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)
	// swagger doc
	mux.HandleFunc("/text_tools.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(content)
	})

	logger.Info().Msgf("TextTool endpoint created, listening on %s", address)
	return http.ListenAndServe(address, allowCORS(mux))
}
func main() {
	if err := Run(":8888"); err != nil {
		logrus.Fatal(err)
	}
}
