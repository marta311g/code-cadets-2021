package handler

import (
	"context"
	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	"log"
	"strings"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Handler struct {
	calcBetRepository CalcBetRepository
}

func New(calcBetRepository CalcBetRepository) *Handler {
	return &Handler{
		calcBetRepository: calcBetRepository,
	}
}

func (h *Handler) HandleEventUpdates(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.EventUpdate,
) <-chan rabbitmqmodels.BetCalculated {
	betsCalculated := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(betsCalculated)

		for event := range eventUpdates {
			log.Println("Processing event: ", event.Id)
			//get bet
			bets, exists, err := h.calcBetRepository.GetBetsByEventID(ctx, event.Id)
			if err != nil {
				log.Println("Failed to fetch bets which should be updated, error: ", err)
				continue
			}
			if !exists {
				log.Println("Bets which should be updated do not exist, eventId: ", event.Id)
				continue
			}

			//update bets
			for _, betToBeUpdated := range bets {
				payout := 0.0

				if strings.EqualFold(event.Outcome, "won") {
					payout = float64(betToBeUpdated.Payment * betToBeUpdated.SelectionCoefficient)
				}

				betToBePublished := rabbitmqmodels.BetCalculated{
					Id:                   betToBeUpdated.Id,
					Status:               event.Outcome,
					SelectionId:          betToBeUpdated.SelectionId,
					SelectionCoefficient: betToBeUpdated.SelectionCoefficient,
					Payment:              betToBeUpdated.Payment,
					Payout:               payout,
				}

				select {
				case betsCalculated <- betToBePublished:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return betsCalculated
}

func (h *Handler) HandleBets(ctx context.Context, bets <-chan rabbitmqmodels.Bet) error {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for bet := range bets {
			log.Println("Processing bet: ", bet.Id)
			//if bet not in db

			newBet := domainmodels.Bet{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment,
			}

			exists, err := h.calcBetRepository.CalcBetWithIDExists(ctx, bet.Id)
			if err != nil {
				log.Println("Failed to fetch a bet which should be inserted, error: ", err)
				continue
			}
			if exists {
				log.Println("Updating bet:", bet.Id)
				err = h.calcBetRepository.UpdateCalcBet(ctx, newBet)
				if err != nil {
					log.Println("Failed to update bet, error: ", err)
				}
				continue
			}

			//insert bet
			log.Println("Inserting new bet with id ", bet.Id)

			err = h.calcBetRepository.InsertCalcBet(ctx, newBet)
			if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}
		}
	}()

	return nil
}
