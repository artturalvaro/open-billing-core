package subscription

import "open-billing-core/internal/domain/subscription"

type CreateSubscriptionInput struct {
	CustomerID string
	PlanID     string
}

type CreateSubscription struct {
	repo subscription.Repository
}

func NewCreateSubscription(repo subscription.Repository) *CreateSubscription {
	return &CreateSubscription{repo: repo}
}

func (uc *CreateSubscription) Execute(input CreateSubscriptionInput) (*subscription.Subscription, error) {
	sub := subscription.NewSubscription(input.CustomerID, input.PlanID)
	return uc.repo.Save(sub)
}
