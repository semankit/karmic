package karmic

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Commit struct {
	Hash    string
	Message string
}

func (receiver Commit) String() string {
	return fmt.Sprintf("%s %s", receiver.Hash, receiver.Message)
}

// IsBreakingChange returns true if commit message contains "BREAKING CHANGE"
func (receiver Commit) IsBreakingChange() (is bool, err error) {
	if strings.Contains(receiver.Message, "BREAKING CHANGE") {
		is = true
	}

	return
}

// List all commits, can be limited by depth
func (receiver Git) List(depth uint8) []Commit {
	arg := []string{
		"log",
		"--pretty=oneline",
	}

	if depth != 0 {
		arg = append(arg, "-n", strconv.Itoa(int(depth)))
	}

	cmd := exec.Command("git", arg...)
	cmd.Dir = receiver.Path
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command

	commits := make([]Commit, 0)
	buffer := strings.Split(string(out), "\n")

	for _, cursor := range buffer {
		totalChar := len(cursor)
		if totalChar == 0 {
			continue
		}

		commit := Commit{}
		commit.Hash = cursor[0:40]
		commit.Message = cursor[40:totalChar]
		commits = append(commits, commit)
	}

	return commits
}
