package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"strings"

	pb "github.com/ekkinox/grpc-demo/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.TextToolsServer
}

func main() {

	fmt.Println("Starting gRPC server on :50051 ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTextToolsServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) TransformText(ctx context.Context, in *pb.Transformation) (*pb.TransformationResult, error) {

	log.Printf("Received a TransformText with %v\n", in)

	result := performTransformation(in)

	return &pb.TransformationResult{
		Result: result,
	}, nil
}

func (*server) TransformAndSplitText(stream pb.TextTools_TransformAndSplitTextServer) error {

	log.Println("Received a TransformAndSplitText")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
		}

		split := strings.Split(performTransformation(req), " ")

		for _, word := range split {

			err = stream.Send(&pb.TransformationResult{
				Result: word,
			})

			if err != nil {
				log.Fatalf("Error while sending: %v", err)
				return err
			}
		}
	}
}

func performTransformation(t *pb.Transformation) string {
	switch t.Transformer {
	case pb.Transformer_UPPERCASE:
		return strings.ToUpper(t.Text)
	case pb.Transformer_LOWERCASE:
		return strings.ToLower(t.Text)
	default:
		return t.Text
	}
}
