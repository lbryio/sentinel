package meta

import (
	"os/exec"
	"strings"
)

// version and commitMsg get filled in using -ldflags when the binary gets built
var version string
var commitMsg string

// GetVersion returns the software version as commit hash
func GetVersion() string {
	if version != "" {
		return version
	}
	out, err := exec.Command("git", "describe", "--always", "--dirty", "--long").Output()
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(out))
}

// GetCommitMessage returns the commit message for the commit this version was built from.
func GetCommitMessage() string {
	if commitMsg != "" {
		return commitMsg
	}

	out, err := exec.Command("git", "show", "-s", "--format=%s", "HEAD").Output()
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(out))
}
