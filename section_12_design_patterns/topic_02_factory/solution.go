package factory

// Notifier interface
type Notifier interface {
	Send(message string) string
}

// EmailNotifier sends via email.
type EmailNotifier struct{ Address string }

func (e EmailNotifier) Send(message string) string {
	// TODO: Implement this — return "Email to <Address>: <message>"
	return ""
}

// SMSNotifier sends via SMS.
type SMSNotifier struct{ Phone string }

func (s SMSNotifier) Send(message string) string {
	// TODO: Implement this — return "SMS to <Phone>: <message>"
	return ""
}

// SlackNotifier sends via Slack.
type SlackNotifier struct{ Channel string }

func (s SlackNotifier) Send(message string) string {
	// TODO: Implement this — return "Slack #<Channel>: <message>"
	return ""
}

// NewNotifier tạo Notifier theo type.
// "email" → EmailNotifier{Address: target}
// "sms" → SMSNotifier{Phone: target}
// "slack" → SlackNotifier{Channel: target}
// Trả về (nil, error) nếu type không hợp lệ.
func NewNotifier(nType, target string) (Notifier, error) {
	// TODO: Implement this
	return nil, nil
}
