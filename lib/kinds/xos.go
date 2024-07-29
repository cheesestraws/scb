package kinds

import (
	"context"
	"io"
	"regexp"
	"strings"
	"time"
)

// XOS is for ExtremeXOS switches
var XOS = "xos"

var xosTimeout = 10 * time.Second
var xosUserPrompt = regexp.MustCompile(regexp.QuoteMeta("login:"))
var xosPasswordPrompt = regexp.MustCompile(regexp.QuoteMeta("password:"))
var xosPrompt = regexp.MustCompile(`\s?\*?\s?[-\w]+\s?[-\w.~]+(:\d+)? [#>] $`)

func fetchXOS(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.ReadCloser, error) {
	e, _, err := telnetSpawn(device+":23", -1)
	if err != nil {
		return nil, err
	}
	defer e.Close()

	_, _, err = e.Expect(xosUserPrompt, xosTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send(user + "\n")
	if err != nil {
		return nil, err
	}

	_, _, err = e.Expect(xosPasswordPrompt, xosTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send(pass + "\n")
	if err != nil {
		return nil, err
	}

	_, _, err = e.Expect(xosPrompt, xosTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send("disable clipaging\n")
	if err != nil {
		return nil, err
	}

	_, _, err = e.Expect(xosPrompt, xosTimeout)
	if err != nil {
		return nil, err
	}
	err = e.Send("show configuration\n")
	if err != nil {
		return nil, err
	}

	config, _, err := e.Expect(xosPrompt, xosTimeout)
	if err != nil {
		return nil, err
	}

	
	// Get rid of wandering CRs
	config = strings.Replace(config, "\r", "", -1)
	config = strings.TrimPrefix(config, "show configuration\n")
	config = xosPrompt.ReplaceAllString(config, "")

	return io.NopCloser(strings.NewReader(config)), nil
}

func init() {
	kinds[XOS] = fetchXOS
}
