package routes

import "net/http"

const (
	// Get represents an http GET request
	Get = "GET"

	// Post represents an http POST request
	Post = "POST"

	// Patch represents an http PATCH request
	Patch = "PATCH"

	// Put represents an http PUT request
	Put = "PUT"

	// Delete represents an http DELETE request
	Delete = "DELETE"
)

// Methods maps a set of http methods to their http handlers
type Methods map[string]http.HandlerFunc

// Routes maps a route to a map of Method handlers
type Routes map[string]Methods

// Router handles mapping routes their to handlers
type Router interface {
	List() Routes
}
