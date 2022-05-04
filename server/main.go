package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/ekkinox/grpc-demo/proto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.TextToolsServiceServer
}

func main() {

	fmt.Println("Starting gRPC server on :50051 ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTextToolsServiceServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) TransformText(ctx context.Context, in *pb.TransformTextRequest) (*pb.TransformTextResponse, error) {

	log.Printf("Received a TransformText with %v\n", in)

	result := performTransformation(in)

	return &pb.TransformTextResponse{
		Result: result,
	}, nil
}

func (*server) TransformAndSplitText(stream pb.TextToolsService_TransformAndSplitTextServer) error {

	log.Println("Received a TransformAndSplitText")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("End of rpc")
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
		}

		split := strings.Split(performTransformation(req), " ")

		for _, word := range split {

			err = stream.Send(&pb.TransformTextResponse{
				Result: word,
			})

			time.Sleep(500 * time.Millisecond)

			if err != nil {
				log.Fatalf("Error while sending: %v", err)
				return err
			}
		}
	}
}

func performTransformation(t *pb.TransformTextRequest) string {
	switch t.Transformer {
	case pb.Transformer_TRANSFORMER_UPPERCASE:
		return strings.ToUpper(t.Text)
	case pb.Transformer_TRANSFORMER_LOWERCASE:
		return strings.ToLower(t.Text)
	default:
		return t.Text
	}
}
