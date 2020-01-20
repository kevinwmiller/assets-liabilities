package user

import (
	"assets-liabilities/entities"
	"assets-liabilities/errors"
	"assets-liabilities/logging"
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

// FindByID fetches the financial user from the database that matches the given id
func (r *PersistedRepository) FindByID(ctx context.Context, id uint64) (entities.User, error) {
	user := entities.User{}
	result := r.db.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return user, errors.NewErrorWithCode(http.StatusNotFound, fmt.Sprintf("No user found with id %d", id))
		}
		return user, result.Error
	}
	return user, nil
}

// Create creates a new user with the given values
func (r *PersistedRepository) Create(ctx context.Context, data entities.User) (entities.User, error) {
	user := entities.User{}
	result := r.db.Create(&data)
	if result.Error != nil {
		return user, errors.NewErrorWithCode(http.StatusInternalServerError, result.Error.Error())
	}

	// Looking up the newly created object because I have had issues in the past with gorm setting the createdAt and updatedAt times in an inconsistent format
	// after creating a new object
	newData, err := r.FindByID(ctx, data.ID)
	return newData, err
}

// Update updates the given financial user if it exists
func (r *PersistedRepository) Update(ctx context.Context, data entities.User) (entities.User, error) {
	user, err := r.FindByID(ctx, data.ID)
	if err != nil {
		return user, err
	}
	result := r.db.Model(&user).Updates(entities.User{
		FullName: data.FullName,
		Password: data.Password,
	})

	if result.Error != nil {
		logging.Logger(ctx).Error(result.Error)
		return user, errors.NewErrorWithCode(http.StatusInternalServerError, result.Error.Error())
	}
	newData, err := r.FindByID(ctx, data.ID)
	return newData, err

}

// Delete deletes the financial user with the given ID from the database
func (r *PersistedRepository) Delete(ctx context.Context, id uint64) error {
	user, err := r.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}
