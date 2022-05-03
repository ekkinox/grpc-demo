package main

import (
	"context"
	"fmt"
	"github.com/ekkinox/grpc-demo/proto/github.com/ekkinox/grpc-demo/proto"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	proto.TextToolsServiceServer
}

func main() {

	fmt.Println("Starting gRPC server on :50051 ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterTextToolsServiceServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) TransformText(ctx context.Context, in *proto.TransformTextRequest) (*proto.TransformTextResponse, error) {

	log.Printf("Received a TransformText with %v\n", in)

	result := performTransformation(in)

	return &proto.TransformTextResponse{
		Result: result,
	}, nil
}

func (*server) TransformAndSplitText(stream proto.TextToolsService_TransformAndSplitTextServer) error {

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

			err = stream.Send(&proto.TransformTextResponse{
				Result: word,
			})

			if err != nil {
				log.Fatalf("Error while sending: %v", err)
				return err
			}
		}
	}
}

func performTransformation(t *proto.TransformTextRequest) string {
	switch t.Transformer {
	case proto.Transformer_TRANSFORMER_UPPERCASE_UNSPECIFIED:
		return strings.ToUpper(t.Text)
	case proto.Transformer_TRANSFORMER_LOWERCASE:
		return strings.ToLower(t.Text)
	default:
		return t.Text
	}
}
