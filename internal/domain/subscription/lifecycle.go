package subscription

import "errors"

var ErrInvalidTransition = errors.New("invalid subscription state transition")

func (s *Subscription) Activate() error {
	if s.Status != Trialing {
		return ErrInvalidTransition
	}
	s.Status = Active
	return nil
}
