package augment

import (
	"fmt"
	"io"
	"os/exec"
)

// Runs an augmentation command.
// Gives you the option to customize how you want to log the command.
func RunCommand(cmd *exec.Cmd, logger io.Writer) error {
	cmd.Stderr = logger
	cmd.Stdout = logger
	err := cmd.Run() //blocks until sub process is complete
	if err != nil {
		return fmt.Errorf("'%s' failed with err: %s", cmd.String(), err.Error())
	}

	return nil
}
