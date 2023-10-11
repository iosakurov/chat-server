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

	desc "github.com/iosakurov/chat-server/pkg/chat_server_v1"
)

const address = "localhost:50051"

var usernames = []string{gofakeit.Name(), gofakeit.Name()}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal("Произошла ошибка!")
		}
	}()

	client := desc.NewChatAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Print(color.RedString("ChatAPI Create\n"))
	response, err := client.Create(ctx, &desc.CreateRequest{Usernames: usernames})
	if err != nil {
		log.Fatalf("failed to create request: %v", err)
	}
	log.Printf(color.RedString("Create Response:\n"), color.GreenString("%+v", response.GetId()))

	log.Print(color.RedString("ChatAPI SendMessage\n"))
	sendResponse, err := client.SendMessage(ctx, &desc.SendMessageRequest{
		From:      gofakeit.Name(),
		Text:      gofakeit.Quote(),
		Timestamp: timestamppb.New(gofakeit.Date()),
	})
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}
	log.Printf(color.RedString("SendMessage Response:\n"), color.GreenString("%+v", sendResponse))

	log.Print(color.RedString("ChatAPI Delete\n"))
	deleteResponse, err := client.Delete(ctx, &desc.DeleteRequest{Id: 666})
	if err != nil {
		log.Fatalf("failed to delete: %v", err)
	}
	log.Printf(color.RedString("Delete Response:\n"), color.GreenString("%+v", deleteResponse))
}
