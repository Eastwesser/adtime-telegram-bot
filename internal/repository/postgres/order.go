package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"

	"github.com/pkg/errors"

	"adtime-telegram-bot/internal/entity"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(ctx context.Context, order *entity.Order) error {
	query := `INSERT INTO orders 
        (user_id, service, date, contact, created_at) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id`

	err := r.db.QueryRow(ctx, query,
		order.UserID,
		order.Service,
		order.Date,
		order.Contact,
		time.Now(),
	).Scan(&order.ID)

	return errors.Wrap(err, "create order")
}

func (r *OrderRepo) GetByUserID(ctx context.Context, userID int64) ([]entity.Order, error) {
	query := `SELECT id, service, date, contact, created_at 
              FROM orders WHERE user_id = $1`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "query orders")
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var o entity.Order
		if err := rows.Scan(&o.ID, &o.Service, &o.Date, &o.Contact, &o.CreatedAt); err != nil {
			return nil, errors.Wrap(err, "scan order")
		}
		orders = append(orders, o)
	}

	return orders, nil
}
