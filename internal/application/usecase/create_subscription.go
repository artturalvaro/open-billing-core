package usecase

import (
	"time"

	"open-billing-core/internal/domain/subscription"
)

type CreateSubscriptionInput struct {
	CustomerID string
	PlanID     string
}

type CreateSubscription struct {
	Repo subscription.Repository
}

func (uc *CreateSubscription) Execute(input CreateSubscriptionInput) (*subscription.Subscription, error) {
	sub := &subscription.Subscription{
		CustomerID: input.CustomerID,
		PlanID:     input.PlanID,
		Status:     subscription.Trialing,
		StartedAt:  time.Now(),
	}

	return uc.Repo.Save(sub)
}
