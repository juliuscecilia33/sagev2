package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONMap map[string]string

// Custom MarshalJSON and UnmarshalJSON methods to handle the map serialization
func (m JSONMap) Value() (driver.Value, error) {
	// Marshal the map into a JSON string for the database
	return json.Marshal(m)
}

func (m *JSONMap) Scan(value interface{}) error {
	// Scan the JSON string from the database back into the map
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, m)
}