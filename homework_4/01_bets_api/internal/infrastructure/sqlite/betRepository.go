package sqlite

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}


func (r *BetRepository) GetBetByID(ctx context.Context, id string) (domainmodels.BetResponseDto, bool, error) {
	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.BetResponseDto{}, false, nil
	}
	if err != nil {
		return domainmodels.BetResponseDto{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, true, nil
}

func (r *BetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.Bet{}, err
	}
	defer row.Close()

	row.Next()

	var customerId string
	var status string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payoutSql sql.NullInt64

	err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	var payout int
	if payoutSql.Valid {
		payout = int(payoutSql.Int64)
	}

	return storagemodels.Bet{
		Id:                   id,
		CustomerId:           customerId,
		Status:               status,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
		Payout:               payout,
	}, nil
}

func (r *BetRepository) GetBetsByUserId(ctx context.Context, userId string) ([]domainmodels.BetResponseDto, bool, error) {
	storageBets, err := r.queryGetBetsByCustomerID(ctx, userId)
	if err == sql.ErrNoRows {
		return []domainmodels.BetResponseDto{}, false, nil
	}
	if err != nil {
		return []domainmodels.BetResponseDto{}, false, errors.Wrap(err, "bet repository failed to get bets from user with id "+userId)
	}

	var domainBets []domainmodels.BetResponseDto
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, true, nil
}

func (r *BetRepository) queryGetBetsByCustomerID(ctx context.Context, userId string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE customer_id='"+userId+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var resultingBets []storagemodels.Bet
	for row.Next() {
		var id string
		var status string
		var selectionId string
		var selectionCoefficient int
		var payment int
		var payoutSql sql.NullInt64

		err = row.Scan(&id, &userId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int
		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}

		resultingBets = append(resultingBets, storagemodels.Bet{
			Id:                   id,
			CustomerId:           userId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		})
	}

	return resultingBets, nil
}

func (r *BetRepository) GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.BetResponseDto, bool, error) {
	storageBets, err := r.queryGetBetsByStatus(ctx, status)
	if err == sql.ErrNoRows {
		return []domainmodels.BetResponseDto{}, false, nil
	}
	if err != nil {
		return []domainmodels.BetResponseDto{}, false, errors.Wrap(err, "bet repository failed to get bets with status "+status)
	}

	var domainBets []domainmodels.BetResponseDto
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, true, nil
}

func (r *BetRepository) queryGetBetsByStatus(ctx context.Context, status string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE status='"+status+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var resultingBets []storagemodels.Bet
	for row.Next() {
		var id string
		var customerId string
		var selectionId string
		var selectionCoefficient int
		var payment int
		var payoutSql sql.NullInt64

		err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int
		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}

		resultingBets = append(resultingBets, storagemodels.Bet{
			Id:                   id,
			CustomerId:           customerId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		})
	}

	return resultingBets, nil
}
