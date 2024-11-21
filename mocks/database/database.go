package database

import (
	"context"
	"fmt"
	"krishnaiyerhq/golang/topics/mocks/api"
	"krishnaiyerhq/golang/topics/mocks/database/mock"
)

var (
	ErrUnknownClientType = fmt.Errorf("unknown error type")
)

// UserClient abstracts user functions.
type UserClient interface {
	// Create creates a user.
	Create(ctx context.Context, user api.User) error

	// Get gets a user using the ID.
	Get(ctx context.Context, id string) (*api.User, error)

	// Update updates a user.
	Update(ctx context.Context, id string, update api.User) error

	// Delete deletes a user using the ID.
	Delete(ctx context.Context, id string) error
}

// NewUser returns a new user client based on the configuration.
func NewUser(typ string) (UserClient, error) {
	switch typ {
	case "mock":
		return mock.NewUser(), nil
	default:
		return nil, ErrUnknownClientType
	}
}
