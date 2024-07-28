package telnetutils

import (
	"os"
	"testing"
)

func TestWhich(t *testing.T) {
	defer func(){
		if r := recover(); r != nil {
			t.Logf("no telnet installed?")
		}
	}()
	
	telnet := Which()
	t.Logf("telnet at: %s", telnet)
	
	// is it a real file
	info, err := os.Stat(telnet)
	if err != nil {
		t.Errorf("statting alleged telnet %q: %v", telnet, err)
	}
	
	if !info.Mode().IsRegular() {
		t.Errorf("alleged telnet %q is not a regular file?", telnet)
	}
}
