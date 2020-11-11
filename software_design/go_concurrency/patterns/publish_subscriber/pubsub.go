package publish_subscriber

/*Con este patron podemos:
Offer an event-driven architecture where one event can trigger one or more
actions.

(separar publishers/emisor del evento del subsribers/receiver qu ejecuta la acción)
Uncoupling the actions that are performed from the event that triggers them


Provide more than one source event that triggers the same action

*/

type Subscriber interface {
	/// input interface for allowing any kind of type
	Notify(interface{}) error

	/// Close stops the goRoutine where this subscriber is listening
	Close()
}

type Publisher interface {
	start()
	///AddSubscriberCh returns a channels that receives Subscriber types
	AddSubscriberCh() chan<- Subscriber

	///RemoveSubscriberCh returns a channels that receives Subscriber types
	RemoveSubscriberCh() chan<- Subscriber

	/// PublishingCh set messages
	PublishMessage() chan<- interface{}

	/// Stops(the goRoutines) publisher and all subscribers related
	Stop()
}
