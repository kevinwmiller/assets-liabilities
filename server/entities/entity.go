package entities

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sony/sonyflake"
)

// Entity is a copy of the gorm.Model struct with the exception that ID is a sonyflake id type instead of a sequential int.
// This struct also forces hard deletes
type Entity struct {
	ID        string    `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;default:null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;default:null"`
}

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

// BeforeCreate will set a SonyFlake ID rather than a sequential ID.
func (entity *Entity) BeforeCreate(scope *gorm.Scope) error {
	id, err := sf.NextID()
	if err != nil {
		// We have been running our app for 174 years
		panic("Reset the sonyflake start date")
	}
	return scope.SetColumn("ID", strconv.FormatUint(id, 10))
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
