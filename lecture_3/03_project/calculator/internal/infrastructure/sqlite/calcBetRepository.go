package sqlite

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

type CalcBetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetCalculatedMapper
}

func NewCalcBetRepository(dbExecutor DatabaseExecutor, betCalculatedMapper BetCalculatedMapper) *CalcBetRepository {
	return &CalcBetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betCalculatedMapper,
	}
}

// InsertCalcBet inserts the provided bet into the database. An error is returned if the operation
// has failed.
func (r *CalcBetRepository) InsertCalcBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *CalcBetRepository) queryInsertBet(ctx context.Context, bet storagemodels.Bet) error {
	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id,	bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// UpdateCalcBet updates the provided bet in the database. An error is returned if the operation
// has failed.
func (r *CalcBetRepository) UpdateCalcBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to update a bet with id "+bet.Id)
	}


	return nil
}

func (r *CalcBetRepository) queryUpdateBet(ctx context.Context, bet storagemodels.Bet) error {
	updateBetSQL := "UPDATE bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Id)
	return err
}

// GetBetsByEventID fetches bets from the database where the selection_id is event id. The second returned value indicates
// whether a bet exists in DB. If a bet does not exist, an error will not be returned.
func (r *CalcBetRepository) GetBetsByEventID(ctx context.Context, id string) ([]domainmodels.Bet, bool, error) {
	storageBets, err := r.queryGetBetsByEventID(ctx, id)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	var domainBets []domainmodels.Bet
	for _, storageBet := range storageBets {
		domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, true, nil
}

func (r *CalcBetRepository) queryGetBetsByEventID(ctx context.Context, eventId string) ([]storagemodels.Bet, error) {
	rows, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selection_id='"+eventId+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer rows.Close()

	var allBets []storagemodels.Bet

	// This will move to the "next" row for every row.
	// At least I hope that's what it does.
	for rows.Next() {
		var id string
		var selectionId string
		var selectionCoefficient int
		var payment int

		err = rows.Scan(&id, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		allBets = append(allBets, storagemodels.Bet{
			Id:                   id,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		})
	}

	return allBets, nil
}

// CalcBetWithIDExists returns True if a bet with the provided id exists.
// If the bet does not exist, an error will not be returned. <- WRONG and I don't know how to fix it
func (r *CalcBetRepository) CalcBetWithIDExists(ctx context.Context, id string) (bool, error) {
	_, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	return true, nil
}

func (r *CalcBetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.Bet{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	row.Next()

	var selectionId string
	var selectionCoefficient int
	var payment int

	err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	return storagemodels.Bet{
		Id:                   id,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}, nil
}
