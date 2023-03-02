package client

import (
	"github.com/feelgood-inc/flgd-gommon/events"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/jaevor/go-nanoid"
	"github.com/spf13/viper"
)

type TransactionalEventsClient struct {
	httpClient *resty.Client
}

func NewTransactionalEventsClient(config events.TransactionalEventClientConfig) *TransactionalEventsClient {
	var baseURL string
	if config.Host == "" {
		baseURL = "http://localhost:8080"
	} else {
		baseURL = config.Host
	}

	var debug bool
	if config.Debug == true {
		debug = true
	} else {
		debug = false
	}
	client := resty.New().
		SetBaseURL(baseURL).
		SetDebug(debug).
		SetRetryCount(3).
		SetRetryWaitTime(time.Duration(1) * time.Second).
		SetHeaders(map[string]string{
			"Content-Type":     "application/json",
			"User-Agent":       "flgd-resty-client",
			"X-Application-ID": viper.GetString("SERVICE_NAME"),
		})

	return &TransactionalEventsClient{
		httpClient: client,
	}
}

func (c *TransactionalEventsClient) SendEvent(payload events.TransactionalEventPayload) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	go func() {
		nanoID, err := nanoid.Standard(21)
		if err != nil {
			log.Printf("error generating nanoid: %v", err)
		}

		event := events.TransactionEvent{
			Event: payload.Event,
			Metadata: events.TransactionalEventMetadata{
				TransactionID: nanoID(),
				SentAt:        time.Now(),
				SentBy:        payload.SentBy,
			},
			Resource: events.TransactionalEventResource{
				OriginalResource:  payload.Resource.OriginalResource,
				ResultingResource: payload.Resource.ResultingResource,
				ActionTaken:       payload.Resource.ActionTaken,
			},
		}
		resp, err := c.httpClient.
			R().
			SetBody(event).
			Post("/events")
		if err != nil {
			log.Printf("error sending event: %v", err)
		}
		if resp.Error() != 200 {
			log.Printf("error sending event: %v", resp.Error())
		}
	}()

	return nil
}
