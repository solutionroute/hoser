package bolthold

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/solutionroute/hoser"
	bh "github.com/timshannon/bolthold"
)

// Ensure that UserService fulfills the Interface
var _ hoser.UserService = &UserService{}

// UserService implementation for a bolthold datastore.
type UserService struct {
	client *Client
}

// User retrieves user by Email from the datastore or (nil, ErrKeyNotFound)
func (s *UserService) User(email string) (*hoser.User, error) {
	var user hoser.User
	err := s.client.db.FindOne(
		&user,
		bh.Where("Email").Eq(strings.ToLower(email)).Index("Email"))
	if err != nil {
		return nil, err // bh.ErrNotFound or others
	}
	return &user, nil
}

// UserID retrieves a user by User.ID from the datastore.
func (s *UserService) UserID(id int) (*hoser.User, error) {
	var user hoser.User
	// Allowing int as input so we can catch negative values here safely
	if id < 0 {
		return nil, bh.ErrNotFound
	}
	err := s.client.db.Get(uint64(id), &user)
	if err != nil {
		return nil, errors.Wrapf(err, "User not found by ID: %d", id) // bh.ErrNotFound or others
	}
	return &user, nil
}

// AddUser adds User to the datastore. User.ID will be autoincremented and
// populated.  Caller should perform User.Validate() before adding or updating
// users. Email is stored in lowercase.
func (s *UserService) AddUser(user *hoser.User) error {

	// New User must not have an ID
	if user.ID > 0 {
		return errors.Errorf("Can't create a new user. Non-nil ID %d", user.ID)
	}
	if err := user.Validate(); err != nil {
		return err
	}
	// normalize
	user.Email = strings.ToLower(user.Email)
	user.CreatedAt = time.Now()
	// email must not already be in datastore
	err := s.client.db.FindOne(&hoser.User{}, bh.Where("Email").Eq(user.Email))
	if err == nil {
		return errors.Errorf("Can't create new user; user email [%s] already exists.", user.Email)
	}
	// add the user
	if err = s.client.db.Insert(bh.NextSequence(), user); err != nil {
		return errors.Wrapf(err, "db.Insert failed for user: %#v", user)
	}
	return nil
}

// UpdateUser updates all fields of user stored within the datastore
// Caller should perform User.Validate() before adding or updating users.
func (s *UserService) UpdateUser(user *hoser.User) error {
	user.UpdatedAt = time.Now()
	// normalize
	user.Email = strings.ToLower(user.Email)
	user.UpdatedAt = time.Now()
	if err := s.client.db.Update(user.ID, user); err != nil {
		return errors.Wrapf(err, "db.Update failed for user: %#v", user)
	}
	return nil
}
