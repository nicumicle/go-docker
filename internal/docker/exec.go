package docker

import (
	"os"
	"os/exec"
)

func run(args ...string) error {
	cmd := execCommand(args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func getOutput(args ...string) ([]byte, error) {
	return execCommand(args...).Output()
}

func execCommand(args ...string) *exec.Cmd {
	return exec.Command("docker", args...)
}
