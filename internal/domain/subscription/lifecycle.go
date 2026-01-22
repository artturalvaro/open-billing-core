package subscription

import (
	"errors"
	"time"
)

var ErrInvalidTransition = errors.New("invalid subscription state transition")

func (s *Subscription) Activate() error {
	if s.Status != Trialing {
		return ErrInvalidTransition
	}
	s.Status = Active
	s.StartedAt = time.Now()
	return nil
}

func (s *Subscription) Cancel() error {
	if s.Status == Canceled {
		return ErrInvalidTransition
	}
	s.Status = Canceled
	return nil
}

func (s *Subscription) MarkPastDue() error {
	if s.Status != Active {
		return ErrInvalidTransition
	}
	s.Status = PastDue
	return nil
}

func (s *Subscription) Reactivate() error {
	if s.Status != PastDue {
		return ErrInvalidTransition
	}
	s.Status = Active
	return nil
}
