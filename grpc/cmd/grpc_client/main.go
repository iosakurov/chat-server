package main

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"

	desc "github.com/iosakurov/chat-server/grpc/pkg/chat_server_v1"
)

const address = "localhost:50051"

var usernames = []string{gofakeit.Name(), gofakeit.Name()}

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := desc.NewChatAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf(color.RedString("ChatAPI Create\n"))
	response, err := client.Create(ctx, &desc.CreateRequest{Usernames: usernames})
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}
	log.Printf(color.RedString("Create Response:\n"), color.GreenString("%+v", response.GetId()))

	log.Printf(color.RedString("ChatAPI SendMessage\n"))
	sendResponse, sendErr := client.SendMessage(ctx, &desc.SendMessageRequest{
		From:      gofakeit.Name(),
		Text:      gofakeit.Quote(),
		Timestamp: timestamppb.New(gofakeit.Date()),
	})
	if sendErr != nil {
		log.Fatalf("failed to send message: %v", sendErr)
	}
	log.Printf(color.RedString("SendMessage Response:\n"), color.GreenString("%+v", sendResponse))

	log.Printf(color.RedString("ChatAPI Delete\n"))
	deleteResponse, deleteErr := client.Delete(ctx, &desc.DeleteRequest{Id: 666})
	if deleteErr != nil {
		log.Fatalf("failed to delete: %v", deleteErr)
	}
	log.Printf(color.RedString("Delete Response:\n"), color.GreenString("%+v", deleteResponse))

}
