package user

import (
	"assets-liabilities/entities"
	"assets-liabilities/repositories/user"
	"context"
)

// Model contains the business logic for all user related actions
type Model struct {
	r user.Repository
}

// New returns a new instance of a user model with a connection to the given repository
func New(r user.Repository) *Model {
	return &Model{
		r,
	}
}

// FindByID returns the user matching the given id
func (m *Model) FindByID(ctx context.Context, id string) (entities.User, error) {
	return m.r.FindByID(ctx, id)
}

// Create creates a new user
func (m *Model) Create(ctx context.Context, data entities.User) (entities.User, error) {
	entities.SanitizeUser(&data)
	var err error
	data.Password, err = HashPassword(ctx, data.Password)
	if err != nil {
		return entities.User{}, err
	}
	return m.r.Create(ctx, data)
}

// Update updates the user matching the given id. ID must be set in data.
// TODO: Define a different method to handle changing passwords. The new method should take the old password and the new password
// for extra security
func (m *Model) Update(ctx context.Context, data entities.User) (entities.User, error) {
	entities.SanitizeUser(&data)
	if data.Password != "" {
		var err error
		data.Password, err = HashPassword(ctx, data.Password)
		if err != nil {
			return entities.User{}, err
		}
	}
	return m.r.Update(ctx, data)
}

// Delete deletes the user matching the given id
func (m *Model) Delete(ctx context.Context, id string) error {
	return m.r.Delete(ctx, id)
}
