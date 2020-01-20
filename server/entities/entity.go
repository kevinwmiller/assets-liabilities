package entities

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// UUID adds a column called ID and will set it to a new v4 uuid on creation
type UUID struct {
}

// BeforeCreate will set a UUID rather than numeric ID.
func (entity *UUID) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

// Entity is a copy of the gorm.Model struct with the exception that ID is a uuid type instead of an int. This struct also forces hard deletes
type Entity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;default:null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;default:null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (entity *Entity) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

// QueryParams contain information for paginating and ordering a FindAll query
type QueryParams struct {
	Limit  *int
	Offset *int
}

const (
	// MaxLimit is the maximum limit that can be specified per query requests
	MaxLimit = 1000
)
