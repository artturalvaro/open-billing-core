package subscription

import (
	"testing"

	domain "open-billing-core/internal/domain/subscription"
	"open-billing-core/internal/infrastructure/repository/memory"
)

func TestCreateSubscription(t *testing.T) {
	repo := memory.NewSubscriptionRepository()
	uc := NewCreateSubscription(repo)

	input := CreateSubscriptionInput{
		CustomerID: "cust_123",
		PlanID:     "plan_basic",
	}

	sub, err := uc.Execute(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sub.ID == "" {
		t.Errorf("expected subscription ID to be set")
	}

	if sub.Status != domain.Trialing {
		t.Errorf("expected status trialing, got %s", sub.Status)
	}

	if sub.CustomerID != input.CustomerID {
		t.Errorf("customerID mismatch")
	}
}
