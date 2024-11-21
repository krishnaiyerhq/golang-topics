package api

// User is a user.
type User struct {
	ID   string // This is the primary ID. Must be unique.
	Age  int
	Name string
}
