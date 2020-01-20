package record

import (
	"assets-liabilities/entities"
	"assets-liabilities/repositories/record"
	"context"
)

// Model contains the business logic for all record related actions
type Model struct {
	r record.Repository
}

// New returns a new instance of a record model with a connection to the given repository
func New(r record.Repository) *Model {
	return &Model{
		r,
	}
}

// FindByID returns the financial record matching the given id
func (m *Model) FindByID(ctx context.Context, id string) (entities.Record, error) {
	return m.r.FindByID(ctx, id)
}

// List returns all financial records matching the given search parameters
func (m *Model) List(ctx context.Context, where entities.SearchParams, params entities.QueryParams) ([]entities.Record, error) {
	return m.r.List(ctx, where, params)
}

// Create creates a new financial record
func (m *Model) Create(ctx context.Context, data entities.Record) (entities.Record, error) {
	entities.SanitizeRecord(&data)
	return m.r.Create(ctx, data)
}

// Update updates the financial record matching the given id. ID must be set in data
func (m *Model) Update(ctx context.Context, data entities.Record) (entities.Record, error) {
	entities.SanitizeRecord(&data)
	return m.r.Update(ctx, data)
}

// Delete deletes the record matching the given id
func (m *Model) Delete(ctx context.Context, id string) error {
	return m.r.Delete(ctx, id)
}
