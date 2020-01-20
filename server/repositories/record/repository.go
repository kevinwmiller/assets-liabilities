package record

import (
	"assets-liabilities/entities"
	"context"
)

// Repository is the data access object for the Record type
type Repository interface {
	FindByID(ctx context.Context, id uint64) (entities.Record, error)
	List(ctx context.Context, where *entities.Record, params *entities.QueryParams) ([]entities.Record, error)
	Create(ctx context.Context, data entities.Record) (entities.Record, error)
	Update(ctx context.Context, data entities.Record) (entities.Record, error)
	Delete(ctx context.Context, id uint64) error
}
