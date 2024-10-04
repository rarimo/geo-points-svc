package config

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"google.golang.org/api/option"
)

type NotificationConfig struct {
	Title string
	Body  string
	// offset from location time
	SendAt int
	Topic  string

	creds []byte
}

type DailyQuestions struct {
	Location    *time.Location
	RawLocation int

	Notifications *NotificationConfig

	deadlines    map[string]Deadline
	lastDeadline *time.Time
	disabled     bool

	mu sync.Mutex
}

type Deadline struct {
	ID int
	At time.Time
}

func (c *config) DailyQuestions() *DailyQuestions {
	return c.DailyQuestion.Do(func() interface{} {
		var cfg struct {
			Timezone      int `fig:"timezone"`
			Notifications struct {
				Title  string `fig:"title,required"`
				Body   string `fig:"body"`
				SendAt int    `fig:"send_at"`
				Creds  string `fig:"creds_file,required"`
				Topic  string `fig:"topic,required"`
			} `fig:"notifications,required"`
		}

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "daily_questions")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out daily questions config: %w", err))
		}

		if cfg.Timezone < -12 || cfg.Timezone > 12 {
			panic(fmt.Errorf("timezone must be between -12 and 12"))
		}

		location := time.FixedZone(fmt.Sprint(cfg.Timezone), cfg.Timezone*3600)

		var notificationsConfig NotificationConfig
		notificationsConfig.Title = cfg.Notifications.Title
		notificationsConfig.Body = cfg.Notifications.Body
		notificationsConfig.SendAt = cfg.Notifications.SendAt
		notificationsConfig.Topic = cfg.Notifications.Topic

		notificationsConfig.creds, err = os.ReadFile(cfg.Notifications.Creds)
		if err != nil {
			panic(fmt.Errorf("failed to read firebase creds: %w", err))
		}

		return &DailyQuestions{
			Location:      location,
			RawLocation:   cfg.Timezone,
			Notifications: &notificationsConfig,
			deadlines:     make(map[string]Deadline),
			mu:            sync.Mutex{},
		}
	}).(*DailyQuestions)
}

func (q *DailyQuestions) SendNotification() error {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(q.Notifications.creds))
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}

	msg := q.Notification()

	client, err := app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("failed to get Messaging client: %w", err)
	}

	if _, err = client.Send(ctx, msg); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (q *DailyQuestions) Notification() *messaging.Message {
	return &messaging.Message{
		Topic: q.Notifications.Topic,
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					MutableContent: true,
					Alert: &messaging.ApsAlert{
						Title: q.Notifications.Title,
						Body:  q.Notifications.Body,
					},
				},
				CustomData: map[string]interface{}{
					"type": "daily_question",
				},
			},
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Data: map[string]string{
				"type":        "daily_question",
				"title":       q.Notifications.Title,
				"description": q.Notifications.Body,
			},
		},
	}
}

func (q *DailyQuestions) LocalTime(date time.Time) time.Time {
	return date.In(q.Location)
}

func (q *DailyQuestions) GetDeadline(nullifier string) *Deadline {
	q.mu.Lock()
	defer q.mu.Unlock()

	deadline, ok := q.deadlines[nullifier]
	if !ok {
		return nil
	}
	return &deadline
}

func (q *DailyQuestions) SetDeadline(nullifier string, id int, duration time.Duration) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.disabled {
		return false
	}

	date := time.Now().UTC().Add(duration)
	q.deadlines[nullifier] = Deadline{ID: id, At: date}
	q.lastDeadline = &date
	return true
}

func (q *DailyQuestions) ZeroDeadline(nullifier string) *Deadline {
	q.mu.Lock()
	defer q.mu.Unlock()

	deadline, ok := q.deadlines[nullifier]
	if !ok {
		return nil
	}
	deadline.At = time.Unix(0, 0)
	q.deadlines[nullifier] = deadline

	return &deadline
}

func (q *DailyQuestions) ClearDeadlines() map[int]int {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.disabled = true
	if q.lastDeadline != nil && q.lastDeadline.After(time.Now().UTC()) {
		q.mu.Unlock()
		timer := time.NewTimer(q.lastDeadline.Sub(time.Now().UTC()))
		<-timer.C
		q.mu.Lock()
	}
	q.disabled = false

	counts := make(map[int]int)
	for _, k := range q.deadlines {
		if _, ok := counts[k.ID]; !ok {
			counts[k.ID] = 0
		}
		counts[k.ID]++
	}

	q.deadlines = make(map[string]Deadline)
	return counts
}
