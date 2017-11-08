package main

import (
	"google.golang.org/grpc"
	"log"
	"respublica.one/task/api/grpc/gen/message_sender"
	"golang.org/x/net/context"
)

func main() {

	conn, err := grpc.Dial(":5555", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := message_sender.NewMessengerClient(conn)

	send_response, err := client.SendMessage(context.Background(), &message_sender.SendMessageRequest{MessageText:"First Message"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sent message %s: %s\n", send_response.MessageUid, "First Message")

	get_response, err := client.GetMessage(context.Background(), &message_sender.GetMessageRequest{MessageUid:send_response.MessageUid})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got message %s: %s\n", send_response.MessageUid, get_response.MessageText)


}
