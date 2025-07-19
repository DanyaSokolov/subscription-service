package repository

import (
	"context"
	"database/sql"

	"github.com/DanyaSokolov/subscription-service/internal/model"
	"github.com/google/uuid"
)

type SubscriptionRepository struct {
	DB *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{DB: db}
}

func (r *SubscriptionRepository) Create(ctx context.Context, s *model.Subscription) error {
	query := `
		INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	return r.DB.QueryRowContext(ctx, query, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate).Scan(&s.ID)
}

func (r *SubscriptionRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error) {
	query := `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscriptions WHERE id=$1
	`
	row := r.DB.QueryRowContext(ctx, query, id)
	var s model.Subscription
	var endDate *string
	err := row.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &endDate)
	if err != nil {
		return nil, err
	}
	if endDate != nil {
		s.EndDate = endDate 
	}
	return &s, nil
}

func (r *SubscriptionRepository) List(ctx context.Context) ([]*model.Subscription, error) {
	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []*model.Subscription
	for rows.Next() {
		var s model.Subscription
		var endDate *string
		if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &endDate); err != nil {
			return nil, err
		}
		if endDate != nil {
			s.EndDate = endDate
		}
		subs = append(subs, &s)
	}
	return subs, nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, s *model.Subscription) error {
	query := `
		UPDATE subscriptions
		SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5
		WHERE id=$6
	`
	_, err := r.DB.ExecContext(ctx, query, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, s.ID)
	return err
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM subscriptions WHERE id=$1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

//userID *uuid.UUID, serviceName *string, fromDate, toDate string

func (r *SubscriptionRepository) TotalCost(ctx context.Context, s *model.Subscription) (int, error) {

	parsedUUID, err := uuid.Parse(s.UserID)
	if err != nil {
		return 0, err
	}
	
	query := `
		SELECT SUM(price) FROM subscriptions
	WHERE start_date >= $1 AND (end_date IS NULL OR end_date <= $2)
	AND user_id = $3 AND service_name = $4
	`
	args := []interface{}{s.StartDate, *s.EndDate, parsedUUID, s.ServiceName}

	var total sql.NullInt64
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	if total.Valid {
		return int(total.Int64), nil
	}
	return 0, nil
}
