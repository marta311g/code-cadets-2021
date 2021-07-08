package main

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const activeBetsURL = "/bets?status=active"
const eventUpdateURL = "/event/update"

func getActiveEvents() ([]string, error) {
	httpResponse, err := http.Get(activeBetsURL)
	if err != nil {
		return nil, err
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var decodedBets []bet
	err = json.Unmarshal(bodyContent, &decodedBets)
	if err != nil {
		return nil, err
	}

	var eventIds []string
	for _, bet := range decodedBets {
		eventIds = append(eventIds, bet.selectionId)
	}

	return eventIds, nil
}

func postEventUpdates(eventIds []string) error {
	rand.Seed(time.Now().UTC().UnixNano())
	status := []string{"won", "lost"}
	for _, eventId := range eventIds {

		eventJson, err := json.Marshal(eventUpdate{
			id: eventId,
			outcome: status[rand.Intn(2)],
		})

		_, err = http.Post(eventUpdateURL, "application/json", bytes.NewBuffer(eventJson))
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	eventIds, err := getActiveEvents()
	if err != nil {
		log.Fatal(errors.WithMessage(err, "failed to get active bets"))
	}

	err = postEventUpdates(eventIds)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "failed to post event updates"))
	}
}

type bet struct {
	id                   string  `json:"id"`
	customerId           string  `json:"customer_id"`
	selectionId          string  `json:"selection_id"`
	selectionCoefficient float64 `json:"selection_coefficient"`
	payment              float64 `json:"payment"`
	payout               float64 `json:"payout"`
}

type eventUpdate struct {
	id      string `json:"id"`
	outcome string `json:"outcome"`
}
