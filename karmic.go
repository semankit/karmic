// Package karmic is a wrapper around git command line made for Semankit
//
// sources available at https://github.com/semankit/karmic
package karmic

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Git struct {
	Path string
}

func New(repoPath *string) (k Git) {
	if repoPath != nil {
		k.Path = *repoPath
	} else {
		execPath, _ := os.Getwd()
		k.Path = execPath
	}

	return
}

func (receiver Git) Version() (version string, err error) {
	version = "undetermined"
	cmd := exec.Command("git", "--versioning")
	cmd.Dir = receiver.Path
	output, cmdErr := cmd.Output()
	if cmdErr != nil {
		err = errors.New("could not determine git versioning")
		return
	}

	buffer := strings.Split(string(output), "\n")
	if len(buffer) == 1 {
		version = buffer[0]
	}

	if len(buffer) == 1 || len(buffer) == 2 {
		version = buffer[0]
	}

	return
}

func (receiver Git) IsInstalled() (isInstalled bool, err error) {
	cmd := exec.Command("which", "git")
	_, cmdErr := cmd.Output()
	if cmdErr != nil {
		if cmdErr.(*exec.ExitError).ExitCode() == 1 {
			err = errors.New("error, git is not installed")
		} else {
			err = errors.New("error, could not determine if git is installed")
		}

		return
	}

	isInstalled = true

	return
}
