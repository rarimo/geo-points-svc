package notificationdailyquestion

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func getMessage(msgFile string) (*messaging.Message, error) {
	file, err := os.Open(msgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data := &struct {
		Msg messaging.Message `json:"message"`
	}{}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode message: %w", err)
	}

	return &data.Msg, nil
}

func main() {
	credFile := flag.String("file", "creds.json", "Path to file with Firebase service account credentionals in json")
	msgFile := flag.String("msg", "msg.json", "Path to file with message. https://firebase.google.com/docs/cloud-messaging/concept-options")
	flag.Parse()

	msg, err := getMessage(*msgFile)
	if err != nil {
		log.Fatalf("failed to get message: %v\n", err)
	}
	msg.APNS = &messaging.APNSConfig{
		Headers: map[string]string{
			"apns-priority": "10",
		},
		Payload: &messaging.APNSPayload{
			Aps: &messaging.Aps{
				MutableContent: true,
			},
		},
	}

	fmt.Printf("%+v\n", msg)

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(*credFile))
	if err != nil {
		log.Fatalf("failed to initialize app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("failed to get Messaging client: %v\n", err)
	}

	response, err := client.Send(ctx, msg)
	if err != nil {
		log.Fatalf("failed to send message: %v\n", err)
	}

	log.Printf("Sucess: %s\n", response)
}
