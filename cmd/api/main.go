package main

import (
	"fmt"

	"open-billing-core/internal/application/usecase"
	infra "open-billing-core/internal/infrastructure/subscription"
)

func main() {
	repo := infra.NewMemoryRepository()

	createSub := usecase.CreateSubscription{
		Repo: repo,
	}

	sub, _ := createSub.Execute(usecase.CreateSubscriptionInput{
		CustomerID: "cust_1",
		PlanID:     "plan_basic",
	})

	fmt.Println("Subscription created:", sub.ID, sub.Status)

	_ = sub.Activate()

	fmt.Println("Subscription status after activation:", sub.Status)
}
