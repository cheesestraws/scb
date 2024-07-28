package telnetutils

import (
	"strings"
	"os/exec"
)

// Which() returns the path to telnet or gtelnet.  Panics if no telnet
// is installed.
func Which() string {
	telnet_names := []string{"telnet", "gtelnet"}

	for _, name := range telnet_names {
		out, err := exec.Command("which", name).Output()
		if err == nil {
			return strings.TrimSpace(string(out))
		}
	}
	
	panic("no telnet installed")
}
