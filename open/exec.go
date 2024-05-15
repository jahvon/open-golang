//go:build !windows && !darwin
// +build !windows,!darwin

package open

import (
	"fmt"
	"os/exec"
)

// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-open/
// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-mime/

func open(uri string) *exec.Cmd {
	switch {
	case os.Getenv(OpenDisabledEnvKey) != "":
		return exec.Command("echo", fmt.Sprintf("\"xdg-open %s\"", uri))
	default:
		return exec.Command("xdg-open", uri)
	}
}

func openWith(input string, appName string) *exec.Cmd {
	switch {
	case os.Getenv(OpenDisabledEnvKey) != "":
		return exec.Command("echo", fmt.Sprintf("\"%s %s\"", appName, input))
	default:
		return exec.Command(appName, input)
	}
}
