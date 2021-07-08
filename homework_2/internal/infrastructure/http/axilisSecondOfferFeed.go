package http

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_2/internal/domain/models"
)

const secondFeedURL = "http://18.193.121.232/axilis-feed-2"

type AxilisSecondOfferFeed struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewAxilisSecondOfferFeed(
	httpClient *http.Client,
) *AxilisSecondOfferFeed {
	return &AxilisSecondOfferFeed{
		updates:    make(chan models.Odd),
		httpClient: httpClient,
	}
}

func (a *AxilisSecondOfferFeed) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			response, err := a.httpClient.Get(secondFeedURL)
			if err != nil {
				log.Println("axilis offer feed, http get", err)
				continue
			}
			a.processSecondResponse(ctx, response)
		}
	}
}

func (a *AxilisSecondOfferFeed) processSecondResponse(ctx context.Context, response *http.Response) {
	defer response.Body.Close()

	bodyContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading feed 2 API response"),
		)
	}

	for _, odd := range strings.Split(string(bodyContent), "\n") {
		axilisOdd := strings.Split(odd, ",")
		coeff, err := strconv.ParseFloat(axilisOdd[3], 64)
		if err != nil {
			log.Println("parsing coefficient", err)
			return
		}

		odd := models.Odd{
			Id:          axilisOdd[0],
			Name:        axilisOdd[1],
			Match:       axilisOdd[2],
			Coefficient: coeff,
			Timestamp:   time.Now(),
		}

		select {
		case <-ctx.Done():
			return
		case a.updates <- odd:
			// do nothing
		}
	}
}

func (a *AxilisSecondOfferFeed) String() string {
	return "axilis second offer feed"
}

func (a *AxilisSecondOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}
