package mocking

// UserRepository interface for data access.
type UserRepository interface {
	GetByID(id string) (string, error)
	Save(id, name string) error
}

// EmailSender interface.
type EmailSender interface {
	Send(to, subject, body string) error
}

// UserService depends on UserRepository and EmailSender.
type UserService struct {
	repo  UserRepository
	email EmailSender
}

// NewUserService creates service with dependencies.
func NewUserService(repo UserRepository, email EmailSender) *UserService {
	// TODO: Implement this
	return nil
}

// GetUser retrieves user by ID.
func (s *UserService) GetUser(id string) (string, error) {
	// TODO: Implement this using s.repo.GetByID
	return "", nil
}

// CreateUser creates user and sends welcome email.
func (s *UserService) CreateUser(id, name, emailAddr string) error {
	// TODO: Implement this — save to repo, then send welcome email
	return nil
}

// MockUserRepository for testing.
type MockUserRepository struct {
	Users     map[string]string
	SaveCalls []struct{ ID, Name string }
}

func (m *MockUserRepository) GetByID(id string) (string, error) {
	// TODO: Implement this
	return "", nil
}

func (m *MockUserRepository) Save(id, name string) error {
	// TODO: Implement this — record the call
	return nil
}

// MockEmailSender for testing.
type MockEmailSender struct {
	SentEmails []struct{ To, Subject, Body string }
}

func (m *MockEmailSender) Send(to, subject, body string) error {
	// TODO: Implement this — record the call
	return nil
}
