package notificationdailyquestion

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/go-co-op/gocron/v2"
	"github.com/rarimo/geo-points-svc/internal/config"
	"github.com/rarimo/geo-points-svc/internal/data/pg"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cleanquestiondeadlines"
	"github.com/rarimo/geo-points-svc/internal/service/workers/cron"
	"google.golang.org/api/option"
)

func Run(ctx context.Context, cfg config.Config, sig chan struct{}) {
	cron.Init(cfg.Log())
	logger := cfg.Log().WithField("notification", "daily-questions-notification")

	questionsQ := pg.NewDailyQuestionsQ(cfg.DB().Clone())

	offset := cfg.DailyQuestions().LocalTime(cleanquestiondeadlines.AtDayStart(time.Now().UTC())).Hour()
	_, err := cron.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(uint(offset), 0, 0))),
		gocron.NewTask(func() {
			curQuestion, err := questionsQ.FilterDayQuestions(time.Now().UTC()).Get()
			if err != nil {
				logger.Fatalf("error getting daily question: %v", err)
				return
			}
			if curQuestion == nil {
				logger.Infof("There's no daily question today")
				return
			}
			sendingNotifications(cfg.Creds().Path, cfg.DailyQuestionsNotification())
		}),
		gocron.WithName("daily-questions-notification"),
	)
	if err != nil {
		panic(fmt.Errorf("failed to initialize daily job: %w", err))
	}
	sig <- struct{}{}

	cron.Start(ctx)
}

func sendingNotifications(toCreds string, notification *config.DailyQuestionsNotification) {
	credFile := toCreds

	msg := &messaging.Message{
		Topic: "mokalake-rewardable-stage",
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					MutableContent: true,
					Alert: &messaging.ApsAlert{
						Title: notification.Title,
						Body:  notification.Body,
					},
				},
				CustomData: map[string]interface{}{
					"type": "daily_question",
				},
			},
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				Title: notification.Title,
				Body:  notification.Body,
			},
			Data: map[string]string{
				"type":        "daily_question",
				"content":     "{\"event_name\": \"early_test\"}", // custim metadata
				"title":       notification.Title,
				"description": notification.Body,
			},
		},
	}

	fmt.Printf("%+v\n", msg)

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credFile))
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

	log.Printf("Success: %s\n", response)
}
