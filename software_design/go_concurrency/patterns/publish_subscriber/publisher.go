package publish_subscriber

import (
	"errors"
)

type publisher struct {
	/// Slice for storing subscribers
	subscribers []Subscriber
}

func (s *publisher) Notify(msg interface{}) error {
	return errors.New("Not implemented yet")

}

func (s *publisher) start() {

}

func (s *publisher) AddSubscriberCh() chan<- Subscriber {
	return nil
}

func (s *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return nil
}

func (s *publisher) PublishingCh() chan<- interface{} {
	return nil
}

func (s *publisher) Stop() {

}

/// NewWriterSubscriber accepts an ID and an io.Writer as the destination for writing the
func NewPublisher() Publisher {
	return &publisher{}
}
