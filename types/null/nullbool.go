package null

import (
	"database/sql/driver"
	"github.com/spf13/cast"
)

//////////// NullBool ///////////////////////
// Scan implements Scanner interface for database driver
func (n *Bool) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Val, n.Valid = false, false
		return nil
	}

	n.Valid = true
	n.Val, err = cast.ToBoolE(value)
	if err != nil {
		return err
	}

	return nil
}

// Value implements the driver Valuer interface for database driver
func (n Bool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}
