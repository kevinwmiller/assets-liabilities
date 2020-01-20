package entities

import (
	"assets-liabilities/errors"
	"fmt"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

// RecordType is used for an enumeration of record types
type RecordType string

const (
	// Asset should be used as a record's financial type if the record provides current or future economic benefit
	Asset RecordType = "Asset"
	// Liability should be used as a record's financial type if the record is something that must be paid off in the future
	Liability RecordType = "Liability"
)

// ConvStrToRecordType converts a string to a record type
func ConvStrToRecordType(str string) (RecordType, error) {
	if str == string(Asset) {
		return Asset, nil
	} else if str == string(Liability) {
		return Liability, nil
	}
	return "", errors.NewErrorWithCode(http.StatusUnprocessableEntity, fmt.Sprintf("Invalid record type %s specified", str))
}

// Record represents a financial record either an asset or a liability.
type Record struct {
	Entity
	Type    RecordType
	Name    string
	Balance float64
}

// RecordSearchParams contains the fields available for searching for financial records
type RecordSearchParams struct {
	Type RecordType
}

// RecordQueryParams is where I would put record specific order by constants, but it is not really needed for this project
// Gorm's Order function doesn't escape parameters, so there is potential for SQL injection. Need to do more research into existing
// libraries for escaping SQL. Tried gorm.Expr, but that didn't seem to work. The correct solution is probably to switch off of gorm and onto something
// like sqlx
type RecordQueryParams struct {
	// Type RecordType
	// OrderBy RecordOrderBy
}

// SanitizeRecord sanitizes potential HTML elements in the model's string fields
func SanitizeRecord(data *Record) {
	strict := bluemonday.StrictPolicy()
	data.Name = strict.Sanitize(data.Name)
}
