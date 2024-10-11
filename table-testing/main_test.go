package main

import (
	"errors"
	"testing"
)

func TestUser(t *testing.T) {
	// Valid case.
	u := User{
		ID:   "alice",
		Name: "Alice Smith",
		Age:  25,
	}

	err := u.Validate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Invalid case; invalid ID
	u.ID = "&#($())"
	if !errors.Is(u.Validate(), errInvalidUserID) {
		t.Fatalf("unexpected error: %v", err)
	}

	// Invalid Name
	u = User{
		ID:   "alice",
		Age:  25,
		Name: "sjfklsdjfksljfkdsljfdksl;jfsdkjfdksfjsdakfj;asdjfds",
	}
	if !errors.Is(u.Validate(), errNameTooLong) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUsers(t *testing.T) {

	for _, testcase := range []struct {
		// Define the fields.
		Name          string
		User          User
		ExpectedError error
	}{
		// Define the cases.
		{
			Name: "Valid",
			User: User{
				ID:   "alice",
				Name: "Alice Smith",
				Age:  25,
			},
			ExpectedError: nil,
		},
		{
			Name: "UserIDTooShort",
			User: User{
				ID:   "al",
				Name: "Alice Smith",
				Age:  25,
			},
			ExpectedError: errInvalidUserID,
		},
		{
			Name: "UserIDTooLong",
			User: User{
				ID:   "abcdefghijklmnopq",
				Name: "Alice Smith",
				Age:  25,
			},
			ExpectedError: errInvalidUserID,
		},
		{
			Name: "UserIDInvalidCharacters",
			User: User{
				ID:   "abd&(*())",
				Name: "Alice Smith",
				Age:  25,
			},
			ExpectedError: errInvalidUserID,
		},
		{
			Name: "UserNameTooLong",
			User: User{
				ID:   "alice",
				Name: "sjfklsdjfksljfkdsljfdksl;jfsdkjfdksfjsdakfj;asdjfds",
				Age:  25,
			},
			ExpectedError: errNameTooLong,
		},
		{
			Name: "UserAgeTooLow",
			User: User{
				ID:   "alice",
				Name: "Alice Smith",
				Age:  17,
			},
			ExpectedError: errAgeOutOfRange,
		},
		{
			Name: "UserAgeTooHigh",
			User: User{
				ID:   "alice",
				Name: "Alice Smith",
				Age:  121,
			},
			ExpectedError: errAgeOutOfRange,
		},
	} {
		// Run the tests.
		t.Run(testcase.Name, func(t *testing.T) {
			err := testcase.User.Validate()
			if !errors.Is(err, testcase.ExpectedError) {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}

}
