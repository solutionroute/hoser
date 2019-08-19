package hoser

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// User is a representation of an application user
// Email is used as the primary lookup key
// Create new user instances using the NewUser() constructor.
type User struct {
	ID          uint64      `json:"id" boltholdKey:"ID"` // Primary key
	FirstName   string      `json:"first_name" validate:"required"`
	LastName    string      `json:"last_name" validate:"required"`
	Email       string      `json:"email" validate:"required,email" boltholdIndex:"Email"`
	Password    string      `json:"-" validate:"required"` // Hashed password
	Salt        string      `json:"-"`
	Permissions Permissions `json:"-"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// NewUser constructs and returns a User
func NewUser() *User {
	user := &User{}
	user.CreatedAt = time.Now()
	user.Permissions = Permissions{}
	return user
}

// Validate user data according to User 'validate' struct tags
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// Permissions is a map[string]bool representing the permissions or roles a user may have.
// p["is_admin"] = true
type Permissions map[string]bool

// Grant adds the permission to the user
func (p Permissions) Grant(permission string) {
	p[permission] = true
}

// Ungrant removes the permission from the user
func (p Permissions) Ungrant(permission string) {
	delete(p, permission)
}

// IsGranted returns true if the user has the permission
func (p Permissions) IsGranted(permission string) bool {
	perm, _ := p[permission]
	return perm
}

// UserService retrieves, creates and manages Users in a storage
type UserService interface {
	User(string) (*User, error) // Retrieve user by the key which is Email
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
