package mocking

import "testing"

func TestGetUser(t *testing.T) {
	repo := &MockUserRepository{Users: map[string]string{"1": "Alice"}}
	svc := NewUserService(repo, &MockEmailSender{})
	if svc == nil {
		t.Fatal("nil service")
	}
	name, err := svc.GetUser("1")
	if err != nil {
		t.Fatal(err)
	}
	if name != "Alice" {
		t.Errorf("got %q", name)
	}
	_, err = svc.GetUser("999")
	if err == nil {
		t.Error("should error for unknown user")
	}
}

func TestCreateUser(t *testing.T) {
	repo := &MockUserRepository{Users: map[string]string{}}
	email := &MockEmailSender{}
	svc := NewUserService(repo, email)
	err := svc.CreateUser("1", "Bob", "bob@test.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(repo.SaveCalls) != 1 {
		t.Error("Save should be called once")
	}
	if len(email.SentEmails) != 1 {
		t.Error("Email should be sent once")
	}
	if email.SentEmails[0].To != "bob@test.com" {
		t.Error("wrong email recipient")
	}
}
