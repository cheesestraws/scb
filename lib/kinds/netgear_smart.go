package kinds

import (
	"context"
	"io"
	"regexp"
	"strings"
	"time"

	expect "github.com/google/goexpect"
)

// NetgearSmart is for netgear smart switches
var NetgearSmart = "netgear_smart"

var netgearSmartTimeout = 10 * time.Second
var netgearSmartUserPrompt = regexp.MustCompile(regexp.QuoteMeta("Applying Interface configuration, please wait ..."))
var netgearSmartPasswordPrompt = regexp.MustCompile(regexp.QuoteMeta("Password:"))
var netgearSmartUnprivilegedPrompt = regexp.MustCompile(regexp.QuoteMeta("(Broadcom FASTPATH Switching) >"))
var netgearSmartPrivilegedPrompt = regexp.MustCompile(regexp.QuoteMeta("(Broadcom FASTPATH Switching) #"))
var netgearSmartMorePrompt = regexp.MustCompile(`--More-- or \(q\)uit`)

var netgearSmartCleanupMorePrompt = regexp.MustCompile(`\n\n--More-- or \(q\)uit\n\n`)
var netgearSmartCleanupCmd = regexp.MustCompile(`show running-config\n\n`)
var netgearSmartCleanupPrivilegedPrompt = regexp.MustCompile(regexp.QuoteMeta("\n\n(Broadcom FASTPATH Switching) #"))

func fetchNetgearSmart(ctx context.Context, device string, _ string, pass string, reserved interface{}) (io.ReadCloser, error) {
	e, _, err := telnetSpawn(device+":60000", -1)
	if err != nil {
		return nil, err
	}
	defer e.Close()

	_, _, err = e.Expect(netgearSmartUserPrompt, netgearSmartTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send("admin\n")
	if err != nil {
		return nil, err
	}

	_, _, err = e.Expect(netgearSmartPasswordPrompt, netgearSmartTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send(pass + "\n")
	if err != nil {
		return nil, err
	}

	_, _, err = e.Expect(netgearSmartUnprivilegedPrompt, netgearSmartTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send("enable\n")
	if err != nil {
		return nil, err
	}

	_, _, err = e.Expect(netgearSmartPasswordPrompt, netgearSmartTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send("\n")
	if err != nil {
		return nil, err
	}

	_, _, err = e.Expect(netgearSmartPrivilegedPrompt, netgearSmartTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send("show running-config\n")
	if err != nil {
		return nil, err
	}

	var config string
	i := 1
	for i == 1 {
		var result string
		result, _, i, err = e.ExpectSwitchCase(
			[]expect.Caser{
				&expect.Case{
					R: netgearSmartPrivilegedPrompt,
				},
				&expect.Case{
					R: netgearSmartMorePrompt,
				},
			},
			netgearSmartTimeout)
		if err != nil {
			return nil, err
		}
		config += result
		e.Send(" ")
	}

	// clean up the result
	// First, replace all \r\ns with \n\n to match what the web interface
	// gives us
	config = strings.Replace(config, "\r", "\n", -1)
	config = netgearSmartCleanupMorePrompt.ReplaceAllString(config, "")
	config = netgearSmartCleanupCmd.ReplaceAllString(config, "")
	config = netgearSmartCleanupPrivilegedPrompt.ReplaceAllString(config, "")

	return io.NopCloser(strings.NewReader(config)), nil
}

func init() {
	kinds[NetgearSmart] = fetchNetgearSmart
}
