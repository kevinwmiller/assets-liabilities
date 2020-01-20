package types

import "strconv"

// CreateBool is a helper function that can be used to get a pointer to a literal value
func CreateBool(value bool) *bool {
	return &value
}

// CreateString is a helper function that can be used to get a pointer to a literal value
func CreateString(value string) *string {
	return &value
}

// CreateInt is a helper function that can be used to get a pointer to a literal value
func CreateInt(value int) *int {
	return &value
}

// CreateIntFromString is a helper function that can be used to get a pointer to a literal value
func CreateIntFromString(valueStr string) *int {
	if valueStr != "" {
		if v, err := strconv.Atoi(valueStr); err == nil {
			return &v
		}
	}
	return nil
}

// CreateInt32 is a helper function that can be used to get a pointer to a literal value
func CreateInt32(value int32) *int32 {
	return &value
}

// CreateIntFromInt32 converts an int32 pointer to an int pointer
func CreateIntFromInt32(value *int32) *int {
	var v *int
	if value != nil {
		converted := int(*value)
		v = &converted
	}
	return v
}
