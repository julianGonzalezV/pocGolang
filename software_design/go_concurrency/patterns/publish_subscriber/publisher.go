package publish_subscriber

import (
	"errors"
)

type publisher struct {
	/// Slice for storing subscribers
	subscribers []Subscriber
	/// addSubCh cannel of subcribers
	addSubCh chan Subscriber
	///removeSubCh channel for removing subscribers
	removeSubCh chan Subscriber
	/// for incming message that will be broadcast to all the subscribers
	in chan interface{}
	/// channel for killing all goRoutines
	stop chan struct{}
}

func (p *publisher) Notify(msg interface{}) error {
	return errors.New("Not implemented yet")

}

/// start is a private method (lower case "s" :) )
func (p *publisher) start() {
	/// infinite loop that repeats "select" operation
	for {
		select {
		case msg := <-p.in: /// case for incomming messages
			for _, subCh := range p.subscribers {
				subCh.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			/// O(N) approach
			for i, candidate := range p.subscribers {
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i],
						p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}

		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)
			return

		}
	}
}

func (p *publisher) AddSubscriberCh() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return p.removeSubCh
}

func (p *publisher) PublishMessage() chan<- interface{} {
	return p.in
}

func (p *publisher) Stop() {
	close(p.stop)
}

/// NewWriterSubscriber accepts an ID and an io.Writer as the destination for writing the
func NewPublisher() Publisher {
	return &publisher{}
}
