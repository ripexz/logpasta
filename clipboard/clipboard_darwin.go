// +build darwin

package clipboard

import (
	"os/exec"
)

func copy(text string) error {
	cmd := exec.Command("pbcopy")

	return unixCopy(cmd, text)
}
