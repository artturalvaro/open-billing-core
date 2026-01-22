package memory

import (
	"errors"
	"sync"

	"open-billing-core/internal/domain/subscription"

	"github.com/google/uuid"
)

type SubscriptionRepository struct {
	mu    sync.Mutex
	store map[string]*subscription.Subscription
}

func NewSubscriptionRepository() *SubscriptionRepository {
	return &SubscriptionRepository{
		store: make(map[string]*subscription.Subscription),
	}
}

func (r *SubscriptionRepository) Save(sub *subscription.Subscription) (*subscription.Subscription, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if sub.ID == "" {
		sub.ID = uuid.NewString()
	}

	r.store[sub.ID] = sub
	return sub, nil
}

func (r *SubscriptionRepository) FindByID(id string) (*subscription.Subscription, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	sub, ok := r.store[id]
	if !ok {
		return nil, errors.New("subscription not found")
	}

	return sub, nil
}
