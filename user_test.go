package hoser

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		u       *User
		wantErr bool
	}{

		{"valid",
			&User{
				FirstName: "Frodo",
				LastName:  "Baggins",
				Email:     "frodo@example.com",
				Password:  "29387sdakjh34",
			},
			false,
		},
		{"missing email",
			&User{
				FirstName: "Frodo",
				LastName:  "Baggins",
				Password:  "29387sdakjh34",
			},
			true,
		},
		{"invalid email",
			&User{
				FirstName: "Frodo",
				LastName:  "Baggins",
				Email:     "frodo@example",
			},
			true,
		},
		{"missing firstname",
			&User{
				FirstName: "",
				LastName:  "Baggins",
				Email:     "frodo@example.com",
				Password:  "29387sdakjh34",
			},
			true,
		},
		{"missing lastname",
			&User{
				FirstName: "Bilbo",
				Email:     "bilbo@example.com",
				Password:  "29387sdakjh34",
			},
			true,
		},
		{"missing password",
			&User{
				FirstName: "",
				LastName:  "Baggins",
				Email:     "frodo@example.com",
				Password:  "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foo, _ := json.Marshal(tt.u)
			fmt.Println(string(foo))
			if err := tt.u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
