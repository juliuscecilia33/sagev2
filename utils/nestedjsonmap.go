package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONMap allows for nested JSON objects by using `interface{}` for values
type NestedJSONMap map[string]interface{}

// Value method for database storage
func (m NestedJSONMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan method for retrieving and parsing JSON data from the database
func (m *NestedJSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, m)
}
