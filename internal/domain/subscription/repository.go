package subscription

type Repository interface {
	Save(subscription *Subscription) (*Subscription, error)
	FindByID(id string) (*Subscription, error)
}
