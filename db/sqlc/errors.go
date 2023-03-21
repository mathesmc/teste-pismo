package db

import (
	"fmt"
)

func ErrEmptyParam(s string) error {
	return fmt.Errorf("creating resource: %s should not be empty", s)

}
