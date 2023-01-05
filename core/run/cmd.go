package run

import (
	"fmt"
	"io"
	"os/exec"
)

// Runs an augmentation command.
// Gives you the option to customize how you want to log the command output.
func Cmd(cmd *exec.Cmd, writer io.Writer) error {
	cmd.Stderr = writer
	cmd.Stdout = writer
	err := cmd.Run() //blocks until sub process is complete
	if err != nil {
		return fmt.Errorf("'%s' failed with err: %s", cmd.String(), err.Error())
	}

	return nil
}
