package bolthold_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/solutionroute/hoser"
)

func populateUserDB(client *Client, t *testing.T) {
	s := client.UserService()
	frodo := &hoser.User{
		FirstName:   "Frodo",
		LastName:    "Baggins",
		Email:       "frodo@example.com",
		Password:    "29387sdakjh34",
		Permissions: hoser.Permissions{},
	}
	frodo.Permissions.Grant("lord-of-the-rings")
	err := s.AddUser(frodo)

	if err != nil {
		t.Error(err)
	}
	for i := 1; i <= NUMINSERTS; i++ {
		strNum := strconv.Itoa(i)
		err := s.AddUser(&hoser.User{
			FirstName:   "Mountain" + strNum,
			LastName:    "Dwarf" + strNum,
			Email:       strNum + "md@example.com",
			Password:    strNum + "fakehash",
			Permissions: hoser.Permissions{},
		})
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUser_Permissions(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.UserService()
	populateUserDB(c, t)

	frodo, _ := s.User("frodo@example.com")

	if frodo.Permissions.IsGranted("lord-of-the-rings") != false {
		t.Errorf("Frodo isn't the ring bearer, Grant permission must have failed: %#v", frodo)
	}
}

func TestUserService_User(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.UserService()
	populateUserDB(c, t)

	frodo, _ := s.User("frodo@example.com")

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *hoser.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{"ErrKeyNotFound",
			args{email: "nouser@example.com"},
			nil, true},
		{"valid",
			args{email: "frodo@example.com"},
			frodo, false},
		{"valid",
			args{email: "FRODO@example.com"},
			frodo, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.User(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.User() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.User() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserID(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.UserService()
	populateUserDB(c, t)

	frodo, _ := s.UserID(1)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *hoser.User
		wantErr bool
	}{
		{"nilID", args{id: 0}, nil, true},
		{"valid", args{id: 1}, frodo, false},
		{"keynotFound", args{id: 233}, nil, true},
		{"neg-id", args{id: -1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UserID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_AddUser(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.UserService()

	tests := []struct {
		name    string
		args    hoser.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{"valid",
			hoser.User{
				FirstName: "Frodo",
				LastName:  "Baggins",
				Email:     "frodo@example.com",
				Password:  "invalidpassword"},
			false,
		},
		{"valid",
			hoser.User{
				FirstName: "Gandalf",
				LastName:  "Greybeared",
				Email:     "wizard@example.com",
				Password:  "invalidpassword"},
			false,
		},
		{"ErrIDNotZero",
			hoser.User{
				ID:        uint64(123),
				FirstName: "Frodo",
				LastName:  "Baggins",
				Email:     "frodo@example.com",
				Password:  "invalidpassword"},
			true,
		},
		{"ErrKeyExists",
			hoser.User{
				FirstName: "Frodo",
				LastName:  "Baggins",
				Email:     "FRODO@example.com",
				Password:  "invalidpassword"},
			true,
		},
		{"ErrNilKey",
			hoser.User{
				FirstName: "Frodo",
				LastName:  "Baggins",
				Email:     "", // the actual key
				Password:  "invalidpassword"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.AddUser(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("UserService.AddUser() wantErr %v\nact error = %v", tt.wantErr, err)
			}
		})
	}
	if _, err := s.User("WIZARD@example.com"); err != nil {
		t.Error("Wizard not found, should be.", err)
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.UserService()
	populateUserDB(c, t)

	// retrieve valid users
	frodo, err := s.User("frodo@example.com")
	if err != nil {
		t.Error("frodo here", err)
	}
	// change email
	frodo.Email = "frodo@fooBAZ.com"
	oneMd, err := s.User("1md@example.com")
	if err != nil {
		t.Error("here", err)
	}
	// set ID to some ridiculous value
	oneMd.ID = 987654321
	// fake user
	fake := &hoser.User{
		ID:        uint64(12345678),
		FirstName: "Sauron",
		LastName:  "TheTerrible",
		Email:     "theeye@example.com",
		Password:  "invalidpassword"}

	tests := []struct {
		name    string
		user    *hoser.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{"valid frodo", frodo, false},
		{"ErrKeyNotFound oneMd", oneMd, true},
		{"ErrKeyNotFound fake", fake, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.UpdateUser(tt.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if _, err := s.User("frodo@foobaz.COM"); err != nil {
		t.Error("Frodo@foobaz.com not found, should be.", err)
	}
}
