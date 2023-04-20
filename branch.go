package karmic

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// Branch represents a branch
//
// https://git-scm.com/docs/git-branch
type Branch string

// Branches returns a list of branches
func (receiver Git) Branches() (branches []Branch, err error) {
	cmd := exec.Command("git", "branch", "-l")
	cmd.Dir = receiver.Path
	buffer, err := cmd.Output()
	if err != nil {
		return
	}

	buffer = bytes.TrimSuffix(buffer, []byte{10})
	buffer = bytes.TrimPrefix(buffer, []byte{42})
	buffer = bytes.TrimPrefix(buffer, []byte{32})

	for idx, cursor := range strings.Split(string(buffer), "\n") {
		bCursor := []byte(cursor)
		bCursor = bytes.TrimPrefix(bCursor, []byte{42})
		bCursor = bytes.TrimPrefix(bCursor, []byte{32})
		bCursor = bytes.TrimPrefix(bCursor, []byte{42})
		bCursor = bytes.TrimPrefix(bCursor, []byte{32})
		branches[idx] = Branch(bCursor)
	}

	return
}

// Checkout checks out a branch
func (receiver Git) Checkout(branchName string) (err error) {
	cmd := exec.Command("git", "checkout", branchName)
	cmd.Dir = receiver.Path
	_, err = cmd.Output()
	if err != nil {
		err = errors.New("error, could not checkout branch")
	}

	return
}

// CurrentBranch returns checked-out branch
func (receiver Git) CurrentBranch() Branch {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = receiver.Path
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command

	branches := strings.Split(string(out), "\n")
	if len(branches[len(branches)-1]) == 0 {
		branches = branches[:len(branches)-1]
	}

	return Branch(branches[0])
}

// Exist checks if a branch exists
func (receiver Git) Exist(branchName string) (exist bool, err error) {
	branches, err := receiver.Branches()
	if err != nil {
		err = errors.New("error, could not list branches")
	}

	for _, cursor := range branches {
		if string(cursor) != branchName {
			continue
		}

		exist = true
		break
	}

	return
}
