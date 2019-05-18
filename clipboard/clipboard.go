package clipboard

import "os/exec"

func Copy(text string) error {
	return copy(text)
}

func unixCopy(cmd *exec.Cmd, text string) error {
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
