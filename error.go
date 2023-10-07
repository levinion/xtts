package xtts

import "fmt"

func newRequestError(code int) error {
	return fmt.Errorf("request were not handled rightly, server return: %d", code)
}
