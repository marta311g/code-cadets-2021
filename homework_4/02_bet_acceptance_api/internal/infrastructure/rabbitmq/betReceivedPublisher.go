package rabbitmq

import (
	"encoding/json"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
	"log"

	"github.com/streadway/amqp"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/infrastructure/rabbitmq/models"
)

const contentTypeTextPlain = "text/plain"

// BetReceivedPublisher handles bet received queue publishing.
type BetReceivedPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewBetReceivedPublisher create a new instance of BetReceivedPublisher.
func NewBetReceivedPublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetReceivedPublisher {
	return &BetReceivedPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

func getRandomUUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		errors.WithMessage(err, "failed to generate id")
	}
	return id.String()
}

// Publish publishes an bet received to the queue.
func (p *BetReceivedPublisher) Publish(customerId string, selectionId string, selectionCoefficient float64, payment float64) error {
	betReceived := &models.BetReceivedDto{
		Id:	getRandomUUID(),
		CustomerId:           customerId,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}

	betReceivedJson, err := json.Marshal(betReceived)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        betReceivedJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", betReceivedJson)
	return nil
}
