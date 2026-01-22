package subscription

import "time"

type Status string

const (
	Trialing Status = "trialing"
	Active   Status = "active"
	PastDue  Status = "past_due"
	Canceled Status = "canceled"
)

type Subscription struct {
	ID         string
	CustomerID string
	PlanID     string
	Status     Status
	StartedAt  time.Time
}

func NewSubscription(customerID, planID string) *Subscription {
	return &Subscription{
		CustomerID: customerID,
		PlanID:     planID,
		Status:     Trialing,
	}
}
