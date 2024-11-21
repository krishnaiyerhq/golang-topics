package mock

import (
	"context"
	"krishnaiyerhq/golang/topics/mocks/api"
	"krishnaiyerhq/golang/topics/mocks/database/errors"
)

// UserClient is a mock user client.
type UserClient struct {
	// Local store
	users map[string]api.User // The key is the unique user ID.
}

// NewUser returns a new user client.
func NewUser() UserClient {
	return UserClient{
		users: make(map[string]api.User, 0),
	}
}

// Create implements database.UserClient.
func (client UserClient) Create(ctx context.Context, user api.User) error {
	_, ok := client.users[user.ID]
	if ok {
		// The user already exists, return an error.
		return errors.ErrUserAlreadyExists
	}
	// Create the user; creating here means to write it to the map.
	client.users[user.ID] = user
	return nil
}

// Create implements database.UserClient.
func (client UserClient) Get(ctx context.Context, id string) (*api.User, error) {
	user, ok := client.users[id]
	if !ok {
		// User does not exist in the map, return an error.
		return nil, errors.ErrUserNotFound
	}
	return &user, nil
}

// Update implements database.UserClient.
func (client UserClient) Update(ctx context.Context, id string, update api.User) error {
	return nil
}

// Delete implements database.UserClient.
func (client UserClient) Delete(ctx context.Context, id string) error {
	return nil
}
