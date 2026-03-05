package factory

import "testing"

func TestNewNotifier(t *testing.T) {
	n, err := NewNotifier("email", "a@b.com")
	if err != nil {
		t.Fatal(err)
	}
	got := n.Send("hello")
	if got != "Email to a@b.com: hello" {
		t.Errorf("email = %q", got)
	}

	n, _ = NewNotifier("sms", "+84123")
	if n.Send("hi") != "SMS to +84123: hi" {
		t.Error("sms wrong")
	}

	n, _ = NewNotifier("slack", "general")
	if n.Send("yo") != "Slack #general: yo" {
		t.Error("slack wrong")
	}

	_, err = NewNotifier("unknown", "x")
	if err == nil {
		t.Error("unknown type should error")
	}
}
