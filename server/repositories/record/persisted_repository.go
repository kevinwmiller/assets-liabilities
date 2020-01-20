package record

import (
	"assets-liabilities/entities"
	"assets-liabilities/errors"
	"assets-liabilities/logging"
	"assets-liabilities/types"
	"context"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

// PersistedRepository stores and retrieves data from a persisted data store
type PersistedRepository struct {
	db *gorm.DB
}

// NewPersistedRepository created a new PersistedRepository instance
func NewPersistedRepository(db *gorm.DB) *PersistedRepository {
	return &PersistedRepository{
		db,
	}
}

// FindByID fetches the financial record from the database that matches the given id
func (r *PersistedRepository) FindByID(ctx context.Context, id uint64) (entities.Record, error) {
	record := entities.Record{}

	result := r.db.Where("id = ?", id).Find(&record)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return record, errors.NewErrorWithCode(http.StatusNotFound, fmt.Sprintf("No record found with id %d", id))
		}
		return record, result.Error
	}
	return record, nil
}

// List returns all records matching the given search params
func (r *PersistedRepository) List(ctx context.Context, where *entities.Record, params *entities.QueryParams) ([]entities.Record, error) {
	records := make([]entities.Record, 0)

	if where == nil {
		where = &entities.Record{}
	}

	if params == nil {
		params = &entities.QueryParams{}
	}
	result := r.db.Where(where).
		// Gorm's Order function doesn't escape parameters
		// Order(types.LoadString(params.OrderBy, "")).
		Offset(types.LoadInt(params.Offset, -1)).
		Limit(types.LoadInt(params.Limit, -1)).
		Find(&records)

	if result.Error != nil {
		logging.Logger(ctx).Error(result.Error)
		return records, errors.NewErrorWithCode(http.StatusInternalServerError, result.Error.Error())
	}
	return records, nil
}

// Create creates a new financial record with the given values
func (r *PersistedRepository) Create(ctx context.Context, data entities.Record) (entities.Record, error) {
	record := entities.Record{}
	result := r.db.Create(&data)
	if result.Error != nil {
		return record, errors.NewErrorWithCode(http.StatusInternalServerError, result.Error.Error())
	}
	// Looking up the newly created object because I have had issues in the past with gorm setting the createdAt and updatedAt times in an inconsistent format
	// after creating a new object
	newData, err := r.FindByID(ctx, data.ID)
	return newData, err

}

// Update updates the given financial record if it exists
func (r *PersistedRepository) Update(ctx context.Context, data entities.Record) (entities.Record, error) {
	record, err := r.FindByID(ctx, data.ID)
	if err != nil {
		return record, err
	}
	result := r.db.Model(&record).Updates(entities.Record{
		Name:    data.Name,
		Type:    data.Type,
		Balance: data.Balance,
	})

	if result.Error != nil {
		logging.Logger(ctx).Error(result.Error)
		return record, errors.NewErrorWithCode(http.StatusInternalServerError, result.Error.Error())
	}
	newData, err := r.FindByID(ctx, data.ID)
	return newData, err

}

// Delete deletes the financial record with the given ID from the database
func (r *PersistedRepository) Delete(ctx context.Context, id uint64) error {
	record, err := r.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return r.db.Delete(&record).Error
}
