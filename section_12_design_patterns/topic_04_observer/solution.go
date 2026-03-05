package observer

// Event đại diện cho một sự kiện.
type Event struct {
	Type    string
	Payload string
}

// Observer interface
type Observer interface {
	OnEvent(event Event)
}

// EventBus quản lý observers và phát events.
type EventBus struct {
	observers map[string][]Observer
}

// NewEventBus tạo EventBus mới.
func NewEventBus() *EventBus {
	// TODO: Implement this
	return nil
}

// Subscribe đăng ký observer cho event type cụ thể.
func (eb *EventBus) Subscribe(eventType string, observer Observer) {
	// TODO: Implement this
}

// Publish phát event đến tất cả observers đã đăng ký cho event type đó.
func (eb *EventBus) Publish(event Event) {
	// TODO: Implement this
}

// LogObserver ghi lại tất cả events nhận được.
type LogObserver struct {
	Events []Event
}

// OnEvent implements Observer.
func (l *LogObserver) OnEvent(event Event) {
	// TODO: Implement this
}
