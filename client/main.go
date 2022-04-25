package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/ekkinox/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Starting gRPC client on :50051 ...")

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	c := pb.NewTextToolsClient(conn)

	//doTransformText(c)
	doTransformTextAndSplit(c)
}

func doTransformText(c pb.TextToolsClient) {

	fmt.Println("Calling TransformText ...")

	res, err := c.TransformText(context.Background(), &pb.Transformation{
		Transformer: pb.Transformer_UPPERCASE,
		Text:        "abcdef",
	})

	if err != nil {
		log.Fatalf("Could not call: %v", err)
	}

	fmt.Printf("result : %v", res.Result)
}

func doTransformTextAndSplit(c pb.TextToolsClient) {

	fmt.Println("Calling TransformAndSplitText ...")

	stream, err := c.TransformAndSplitText(context.Background())
	if err != nil {
		log.Fatalf("Could not call: %v", err)
	}

	reqs := []*pb.Transformation{
		{
			Transformer: pb.Transformer_LOWERCASE,
			Text:        "abc DEF hko",
		},
		{
			Transformer: pb.Transformer_LOWERCASE,
			Text:        "KLNLKNLK DSDSD",
		},
		{
			Transformer: pb.Transformer_UPPERCASE,
			Text:        "mlkjlk fdsfsdfds",
		},
		{
			Transformer: pb.Transformer_LOWERCASE,
			Text:        "AAAAAA",
		},
		{
			Transformer: pb.Transformer_UPPERCASE,
			Text:        "bbbbb",
		},
		{
			Transformer: pb.Transformer_LOWERCASE,
			Text:        "AbAbAb",
		},
	}

	wait := make(chan struct{})

	// sending
	go func() {
		for _, req := range reqs {
			fmt.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()

	}()

	// receiving
	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				close(wait)
				log.Fatalf("Error: %v", err)
			}

			fmt.Printf("Received response: %v\n", resp.Result)
		}
		close(wait)
	}()

	<-wait
}
