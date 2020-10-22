package shared

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//Organization represents a company in our plateform
type Product struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Slug        string     `json:"slug,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedByID string     `json:"createdByID,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

//Value marshalls to JSON value
func (p Product) Value() (driver.Value, error) {
	return json.Marshal(p)
}

//Scan unmarshalls from JSON value
func (p *Product) Scan(src interface{}) error {
	if bytes, ok := src.([]byte); ok {
		return json.Unmarshal(bytes, p)
	}
	return errors.New(fmt.Sprint("failed to unmarshal org JSONB from the DB", src))
}
