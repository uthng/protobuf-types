package null

import (
	"database/sql/driver"
	"github.com/spf13/cast"
)

///////// NullInt32 ///////////////////////
// Scan implements Scanner interface for database driver
func (n *NullInt32) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Value, n.Valid = 0, false
		return nil
	}

	n.Valid = true
	v.Value, err = cast.ToInt32E(value)
	if err != nil {
		return err
	}

	return nil
}

// Value implements the driver Valuer interface for database driver
func (n NullInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Value, nil
}

//////////// NullInt64 ///////////////////////
// Scan implements Scanner interface for database driver
func (n *NullInt64) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Value, n.Valid = 0, false
		return nil
	}

	n.Valid = true
	v.Value, err = cast.ToInt64E(value)
	if err != nil {
		return err
	}

	return nil
}

// Value implements the driver Valuer interface for database driver
func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Value, nil
}
