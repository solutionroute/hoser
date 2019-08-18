package hoser

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// TODO - Abstract user and user management later into a separate package.

// User is a representation of an application user
// Email is used as the primary lookup key
type User struct {
	ID        uint64    `json:"id" boltholdKey:"ID"` // Primary key
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email" boltholdIndex:"Email"`
	Password  string    `json:"-" validate:"required"` // Hashed password
	Salt      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate user data according to User 'validate' struct tags
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)

}

// UserService retrieves, creates and manages Users
type UserService interface {
	User(string) (*User, error) // by key, email
	UserID(int) (*User, error)
	AddUser(*User) error
	UpdateUser(*User) error
}

// AuthService provides mechanisms to authenticate a client user
// TODO add an AuthService / locate it appropriately within the source files. i.e.
// type AuthService interface {
//     HashMatchesPassword(string, string) (bool, error)
//     // JWT / other means
// }
