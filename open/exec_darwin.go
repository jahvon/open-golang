//go:build darwin
// +build darwin

package open

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func open(uri string) *exec.Cmd {
	args := []string{uri}
	if os.Getenv(OpenInBackgroundEnvKey) != "" {
		args = append([]string{"-g"}, args...)
	}
	switch {
	case os.Getenv(OpenDisabledEnvKey) != "":
		return exec.Command("echo", fmt.Sprintf("\"open %s\"", strings.Join(args, " ")))
	default:
		return exec.Command("open", args...)
	}
}

func openWith(uri string, appName string) *exec.Cmd {
	args := []string{"-a", appName, uri}
	if os.Getenv(OpenInBackgroundEnvKey) != "" {
		args = append([]string{"-g"}, args...)
	}

	switch {
	case os.Getenv(OpenDisabledEnvKey) != "":
		return exec.Command("echo", fmt.Sprintf("\"open -a %s %s\"", appName, strings.Join(args, " ")))
	default:
		return exec.Command("open", args...)
	}
}
