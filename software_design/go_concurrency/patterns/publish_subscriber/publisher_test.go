package publish_subscriber

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

/// The mockSubscriber struct must implements Subscriber interface:

func (m *mockSubscriber) Notify(msg interface{}) error {
	fmt.Println("notifyTestingFunc in Notify ")
	m.notifyTestingFunc(msg)
	return nil
}

func (m *mockSubscriber) Close() {
	fmt.Println("closeTestingFunc in Close ")
	m.closeTestingFunc()
}

func TestPublisher(t *testing.T) {
	msg := "Message from TestPublisher"
	p := NewPublisher()

	var wg sync.WaitGroup

	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			fmt.Println("notifyTestingFunc")
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

	//Type assertion for getting the concrete class

	///adding the mockSubscriber
	fmt.Println("va a llamar AddSubscriberCh")
	p.AddSubscriberCh() <- sub

	fmt.Println("va a llamar PublishMessage")
	p.PublishMessage() <- msg
	wg.Wait()

	pubCon := p.(*publisher)
	// After calling 	p.AddSubscriberCh() the number of subscriber must be 1
	fmt.Println(" len subs => ", len(pubCon.subscribers))

	if len(pubCon.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}

	fmt.Println("va a llamar RemoveSubscriberCh")
	///
	wg.Add(1)
	p.RemoveSubscriberCh() <- sub
	wg.Wait()
	//Number of subscribers must be zero/0  after calling RemoveSubscriberCh
	fmt.Println(" After call RemoveSubscriberCh len subs => ", len(pubCon.subscribers))
	if len(pubCon.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}

	// Stop to avoid more messages coming and ensure stop the goRoutine
	p.Stop()

}
