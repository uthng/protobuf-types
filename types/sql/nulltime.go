package sql

import (
	"database/sql/driver"
	"time"

	"github.com/spf13/cast"
)

///////// NullTime ///////////////////////
// Scan implements Scanner interface for database driver
func (n *NullTime) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Val, n.Valid = 0, false
		return nil
	}

	t, err := cast.ToTimeE(value)
	if err != nil {
		return err
	}

	n.Valid = true
	n.Val = t.Unix()

	return nil
}

// Value implements the driver Valuer interface for database driver
func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return time.Unix(n.Val, 0), nil
}
