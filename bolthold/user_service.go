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
	user := hoser.NewUser()
	err := s.client.db.FindOne(
		user,
		bh.Where("Email").Eq(strings.ToLower(email)).Index("Email"))
	if err != nil {
		return nil, errors.Wrapf(err, "User not found by email: [%s]", email) // bh.ErrNotFound or others
	}
	return user, nil
}

// UserID retrieves a user by User.ID from the datastore.
func (s *UserService) UserID(id int) (*hoser.User, error) {
	user := hoser.NewUser()
	// Allowing int as input so we can catch negative values here safely
	if id < 0 {
		return nil, errors.New("Invalid: can't query User ID < 0")
	}
	err := s.client.db.Get(uint64(id), user)
	if err != nil {
		return nil, errors.Wrapf(err, "User not found by ID: %d", id) // bh.ErrNotFound or others
	}
	return user, nil
}

// AddUser adds User to the datastore.
// - User.Validate() is executed and any errors are returned immediately without adding to the db
// - User.ID will be autoincremented and populated
func (s *UserService) AddUser(user *hoser.User) error {
	userExists := hoser.NewUser()

	// New User must not have an ID
	if user.ID > 0 {
		return errors.Errorf("Can't create new user; non-nil ID %d", user.ID)
	}
	if err := user.Validate(); err != nil {
		return err
	}
	// normalize
	user.Email = strings.ToLower(user.Email)
	user.CreatedAt = time.Now()
	// email must not already be in datastore
	err := s.client.db.FindOne(userExists, bh.Where("Email").Eq(user.Email))
	if err == nil {
		return errors.Errorf("Can't create new user; user email [%s] already exists.", user.Email)
	}
	// add the user
	if err = s.client.db.Insert(bh.NextSequence(), user); err != nil {
		return errors.Wrapf(err, "Add user failed for user: %#v", user)
	}
	return nil
}

// UpdateUser updates all fields of user stored within the datastore
func (s *UserService) UpdateUser(user *hoser.User) error {
	// Existing user must have a valid ID
	if user.ID < 1 {
		return errors.Errorf("Can't update this user; nil or invalid ID %d\n%#v", user.ID, user)
	}
	if err := user.Validate(); err != nil {
		return err
	}
	user.UpdatedAt = time.Now()
	// normalize
	user.Email = strings.ToLower(user.Email)
	user.UpdatedAt = time.Now()
	if err := s.client.db.Update(user.ID, user); err != nil {
		return errors.Wrapf(err, "Update user failed for user: %#v", user)
	}
	return nil
}
