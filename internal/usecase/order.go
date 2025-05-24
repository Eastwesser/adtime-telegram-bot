package usecase

import (
	"context"
	"go.uber.org/zap"
	"time"

	"github.com/pkg/errors"

	"adtime-telegram-bot/internal/entity"
)

type OrderUseCase struct {
	orderRepo OrderRepository
	userRepo  UserRepository
	logger    *zap.Logger
	cache     Cache
}

func NewOrderUseCase(or OrderRepository, ur UserRepository, l *zap.Logger, c Cache) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: or,
		userRepo:  ur,
		logger:    l,
		cache:     c,
	}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, userID int64, service, date, contact string) (*entity.Order, error) {
	// Проверяем согласие пользователя
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	if !user.Consent {
		return nil, errors.New("user consent required")
	}

	// Валидация даты
	if _, err := time.Parse("02.01.2006", date); err != nil {
		return nil, errors.Wrap(err, "invalid date format")
	}

	order := &entity.Order{
		UserID:  userID,
		Service: service,
		Date:    date,
		Contact: contact,
	}

	if err := uc.orderRepo.Create(ctx, order); err != nil {
		return nil, errors.Wrap(err, "create order")
	}

	// Инвалидируем кэш
	uc.cache.Delete(ctx, userOrdersCacheKey(userID))

	return order, nil
}
