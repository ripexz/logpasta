// +build windows

package clipboard

import "errors"

func copy(text string) error {
	return errors.New("windows not yet supported")
}
