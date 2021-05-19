package http

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed(
	httpClient http.Client,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	// repeatedly:
	for {
		// - get odds from HTTP server
		httpResponse, err := http.Get(axilisFeedURL)
		if err != nil {
			return err
		}

		bodyContent, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return err
		}

		var decodedContent []axilisOfferOdd
		err = json.Unmarshal(bodyContent, &decodedContent)
		if err != nil {
			return nil
		}

		for _, odd := range decodedContent {
			select {
			// - if context is finished, exit and close updates channel
			case <-ctx.Done():
				close(a.updates)
				fmt.Println("finished")
				return nil
			// - write them to updates channel
			case <-time.After(time.Second):
				a.updates <- models.Odd{
					Id:          odd.Id,
					Name:        odd.Name,
					Match:       odd.Match,
					Coefficient: 1,
					Timestamp:   time.Time{},
				}
			}
		}
	}
	// (test your program from cmd/main.go)
	return nil
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}

//jedan queue, 2 httpa , merging u jedan string
// sredisnji dio cita 1 updates channel -> 2 feeda mergat u jedan updates
// ili u contructoru primi 2 interfacea za 2 feeda
