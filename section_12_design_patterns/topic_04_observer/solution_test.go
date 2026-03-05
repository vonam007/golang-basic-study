package observer

import "testing"

func TestEventBus(t *testing.T) {
	bus := NewEventBus()
	if bus == nil {
		t.Fatal("nil")
	}
	log := &LogObserver{}
	bus.Subscribe("user.created", log)
	bus.Publish(Event{Type: "user.created", Payload: "alice"})
	bus.Publish(Event{Type: "user.deleted", Payload: "bob"}) // not subscribed
	if len(log.Events) != 1 {
		t.Errorf("events = %d; want 1", len(log.Events))
	}
	if log.Events[0].Payload != "alice" {
		t.Error("wrong payload")
	}
}

func TestMultipleObservers(t *testing.T) {
	bus := NewEventBus()
	l1, l2 := &LogObserver{}, &LogObserver{}
	bus.Subscribe("click", l1)
	bus.Subscribe("click", l2)
	bus.Publish(Event{Type: "click", Payload: "button"})
	if len(l1.Events) != 1 || len(l2.Events) != 1 {
		t.Error("both observers should receive event")
	}
}
