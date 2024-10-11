package main

import (
	"fmt"
	"regexp"
)

// User is a user.
type User struct {
	ID   string
	Name string
	Age  int
}

const (
	maxUserNameLength = 50
	minUserAge        = 18
	maxUserAge        = 120
)

var (
	userIDRegex = regexp.MustCompile(`^[a-z0-9]{3,16}$`)

	// Error definitions.
	errInvalidUserID = fmt.Errorf("invalid user ID")
	errNameTooLong   = fmt.Errorf("user name longer than 50 characters")
	errAgeOutOfRange = fmt.Errorf("user age must be between 18 and 120")
)

// Validate validates the user.
// ID: Lowercase letters and numbers, between 3 and 16 characters.
// Name: 50 character limit; no other restrictions.
// Age: 18-120
func (u User) Validate() error {
	if !userIDRegex.MatchString(u.ID) {
		return errInvalidUserID
	}
	if len(u.Name) > maxUserNameLength {
		return errNameTooLong
	}
	if u.Age > maxUserAge || u.Age < minUserAge {
		return errAgeOutOfRange
	}

	return nil
}

func main() {

}
