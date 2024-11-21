package database

import (
	"context"
	"krishnaiyerhq/golang/topics/mocks/api"
	"testing"

	"github.com/smarty/assertions"
)

func TestUsers(t *testing.T) {
	ctx := context.Background()
	// Use assertions for better testing.
	a := assertions.New(t)

	// Create a client.
	client, err := NewUser("mock")
	if err != nil {
		t.Fatal(err)
	}

	// Create a user.
	user := api.User{
		ID:   "alice",
		Name: "Alice",
		Age:  25,
	}

	err = client.Create(ctx, user)
	if !a.So(err, assertions.ShouldBeNil) {
		// Assertion failed.
		t.Fatalf("unexpected error: %v", err)
	}

	// Get the created user.
	got, err := client.Get(ctx, user.ID)
	if !a.So(err, assertions.ShouldBeNil) {
		// Assertion failed.
		t.Fatalf("unexpected error: %v", err)
	}
	if !a.So(got, assertions.ShouldEqual, &user) {
		// Assertion failed.
		t.Fatalf("unexpected error: %v", err)
	}
	// Add tests for Update, Delete etc
}
