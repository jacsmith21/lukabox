package db

import (
	"errors"

	"github.com/jacsmith21/lukabox/domain"
	log "github.com/jacsmith21/lukabox/ext/logrus"
)

var users = []*domain.User{
	{ID: 1, Email: "jacob.smith@unb.ca", Password: "password", FirstName: "Jacob", LastName: "Smith", Archived: false},
	{ID: 2, Email: "j.a.smith@live.ca", Password: "password", FirstName: "Jacob", LastName: "Smith", Archived: false},
	{ID: 3, Email: "jacobsmithunb@gmail.com", Password: "password", FirstName: "Jacob", LastName: "Smith", Archived: false},
}

//UserService represents an implementation UserService
type UserService struct {
}

// ValidateUser validatese a user
func (s *UserService) ValidateUser(user *domain.User) error {
	log.WithField("user", user).Info("validate user")
	if user.Email == "" {
		return errors.New("a user must have an email")
	}
	if user.FirstName == "" {
		return errors.New("a user must have a first name")
	}
	if user.LastName == "" {
		return errors.New("a user must have a last name")
	}
	if user.Password == "" {
		return errors.New("a user must have a password")
	}
	return nil
}

// InsertUser creates a user in the database
func (s *UserService) InsertUser(user *domain.User) error {
	if user.ID != 0 {
		return errors.New("user id must equal 0")
	}
	user.ID = users[len(users)-1].ID + 1
	users = append(users, user)
	return nil
}

//Users retrieves a user from the database
func (s *UserService) Users() ([]*domain.User, error) {
	return users, nil
}

//UserByID retrieves a user from the database using their ID
func (s *UserService) UserByID(id int) (*domain.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

//UserByEmail retrieves a user from the database using their email
func (s *UserService) UserByEmail(email string) (*domain.User, error) {
	for _, u := range users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

//UpdateUser updates a user in the datbase
func (s *UserService) UpdateUser(id int, user *domain.User) error {
	for i, u := range users {
		if u.ID == id {
			users[i] = user
			return nil
		}
	}
	return errors.New("article not found")
}
