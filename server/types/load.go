package types

// LoadBool returns value if it is not nil and defaultValue otherwise
func LoadBool(value *bool, defaultValue bool) bool {
	if value != nil {
		return *value
	}
	return defaultValue
}

// LoadString returns value if it is not nil and defaultValue otherwise
func LoadString(value *string, defaultValue string) string {
	if value != nil {
		return *value
	}
	return defaultValue
}

// LoadInt returns value if it is not nil and defaultValue otherwise
func LoadInt(value *int, defaultValue int) int {
	if value != nil {
		return *value
	}
	return defaultValue
}

// LoadInt32 returns value if it is not nil and defaultValue otherwise
func LoadInt32(value *int32, defaultValue int) int {
	if value != nil {
		return int(*value)
	}
	return defaultValue
}
