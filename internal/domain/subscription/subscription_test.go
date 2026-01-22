package subscription

import "testing"

func TestNewSubscription(t *testing.T) {
	sub := NewSubscription("cust_1", "plan_basic")

	if sub.CustomerID != "cust_1" {
		t.Errorf("expected customerID cust_1, got %s", sub.CustomerID)
	}

	if sub.PlanID != "plan_basic" {
		t.Errorf("expected plan_basic, got %s", sub.PlanID)
	}

	if sub.Status != Trialing {
		t.Errorf("expected status pending, got %s", sub.Status)
	}

	if !sub.StartedAt.IsZero() {
		t.Errorf("startedAt should be zero on creation")
	}
}

func TestActivateSubscription(t *testing.T) {
	sub := NewSubscription("cust_1", "plan_basic")

	err := sub.Activate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sub.Status != Active {
		t.Errorf("expected status active, got %s", sub.Status)
	}

	if sub.StartedAt.IsZero() {
		t.Errorf("startedAt should be set on activation")
	}
}

func TestActivateAlreadyActiveSubscription(t *testing.T) {
	sub := NewSubscription("cust_1", "plan_basic")
	_ = sub.Activate()

	err := sub.Activate()
	if err == nil {
		t.Errorf("expected error when activating active subscription")
	}
}
