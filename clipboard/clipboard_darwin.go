// +build darwin

package clipboard

import (
	"os/exec"
)

func copy(text string) error {
	cmd := exec.Command("pbcopy")

	input, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	_, err = input.Write([]byte(text))
	if err != nil {
		return err
	}

	err = input.Close()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
