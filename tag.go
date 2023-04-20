package karmic

import (
	"fmt"
	"os/exec"
	"strings"
)

type Tag struct {
	git   Git
	value string
}

func (receiver Git) Tags() []Tag {
	cmd := exec.Command("git", "tag")
	cmd.Dir = receiver.Path
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}

	// otherwise, print the output from running the command
	buffer := strings.Split(string(out), "\n")
	if len(buffer[len(buffer)-1]) == 0 {
		buffer = buffer[:len(buffer)-1]
	}

	var tags = make([]Tag, len(buffer))
	for idx, e := range buffer {
		tags[idx] = Tag{
			git:   receiver,
			value: e,
		}
	}

	return tags
}

// Commits ...
func (tag Tag) Commits() []Commit {
	cmd := exec.Command("git", "log", fmt.Sprintf("%s...HEAD", tag), "--pretty=oneline")
	cmd.Dir = tag.git.Path
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

		c := Commit{}
		c.Hash = cursor[0:40]
		c.Message = cursor[40:totalChar]
		commits = append(commits, c)
	}

	return commits
}

// String ...
func (tag Tag) String() string {
	return tag.value
}
