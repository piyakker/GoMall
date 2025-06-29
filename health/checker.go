package health

import (
	"context"
	"errors"
	"fmt"
)

type Checker interface {
	Check(ctx context.Context) error
}

func RunAll(ctx context.Context, checkers ...Checker) error {
	var errMsg string
	for _, c := range checkers {
		if err := c.Check(ctx); err != nil {
			errMsg += fmt.Sprintf("%T error: %v\n", c, err)
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}
