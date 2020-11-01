package sql

import (
	"database/sql/driver"
	"github.com/spf13/cast"
)

// Scan implements the Scanner interface.
func (n *NullString) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Val, n.Valid = "null", false
		return nil
	}

	n.Valid = true
	n.Val, err = cast.ToStringE(value)
	if err != nil {
		return err
	}

	return nil
}

// Value implements the driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Val, nil
}
