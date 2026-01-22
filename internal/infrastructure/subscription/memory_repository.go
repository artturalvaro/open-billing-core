package subscription

import (
	"errors"
	"fmt"

	domain "open-billing-core/internal/domain/subscription"
)

type MemoryRepository struct {
	data map[string]*domain.Subscription
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[string]*domain.Subscription),
	}
}

func (r *MemoryRepository) Save(sub *domain.Subscription) (*domain.Subscription, error) {
	if sub.ID == "" {
		sub.ID = generateID()
	}

	r.data[sub.ID] = sub
	return sub, nil
}

var idCounter = 1

func generateID() string {
	id := idCounter
	idCounter++
	return "sub_" + fmt.Sprint(id)
}

func (r *MemoryRepository) FindByID(id string) (*domain.Subscription, error) {
	sub, ok := r.data[id]
	if !ok {
		return nil, errors.New("subscription not found")
	}
	return sub, nil
}
