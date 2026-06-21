//go:build windows

package clipboard

import "os/exec"

func copy(text string) error {
	cmd := exec.Command("clip")

	return unixCopy(cmd, text)
}
