// +build linux

package clipboard

import (
	"errors"
	"os/exec"
)

var command []string

func init() {
	if _, err := exec.LookPath("xclip"); err == nil {
		command = []string{"xclip", "--input", "--clipboard"}
		return
	}

	if _, err := exec.LookPath("xsel"); err == nil {
		command = []string{"xsel", "-in", "-selection", "clipboard"}
		return
	}
}

func copy(text string) error {
	if command == nil || len(command) == 0 {
		return errors.New("no xclip or xsel commands not found")
	}

	cmd := exec.Command(command[0], command[1:]...)

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
