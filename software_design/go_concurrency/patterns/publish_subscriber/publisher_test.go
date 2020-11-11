package publish_subscriber

import (
	"errors"
	"sync"
	"testing"
)

type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

/// The mockSubscriber struct must implements Subscriber interface:

func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

func TestPublisher(t *testing.T) {
	msg := "Message from TestPublisher"
	p := NewPublisher()

	var wg sync.WaitGroup
	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			defer wg.Done()
			s, ok := msg.(string) // preguntando si msg es de tipo string
			if !ok {
				t.Fatal(errors.New("Could not assert"))
			}

			if s != msg {
				t.Fail()
			}

		},
		closeTestingFunc: func() {
			wg.Done()
		},
	}

	///adding the mockSubscriber
	p.AddSubscriberCh() <- sub
	wg.Add(1)
	p.PublishMessage() <- msg
	wg.Wait()

	//Type assertion for getting the concrete class
	pubCon := p.(*publisher)

	// After calling 	p.AddSubscriberCh() the number of subscriber must be 1
	if len(pubCon.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}

	///
	wg.Add(1)
	p.RemoveSubscriberCh() <- sub
	wg.Wait()
	//Number of subscribers must be zero/0  after calling RemoveSubscriberCh
	if len(pubCon.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}

	// Stop to avoid more messages coming and ensure stop the goRoutine
	p.Stop()

}
