package user

import (
	"assets-liabilities/entities"
	"context"
)

// Repository is the data access object for the Record type
type Repository interface {
	FindByID(ctx context.Context, id string) (entities.User, error)
	Create(ctx context.Context, data entities.User) (entities.User, error)
	Update(ctx context.Context, data entities.User) (entities.User, error)
	Delete(ctx context.Context, id string) error
}
