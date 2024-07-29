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

	e.Expect(xosUserPrompt, xosTimeout)
	e.Send(user + "\n")
	e.Expect(xosPasswordPrompt, xosTimeout)
	e.Send(pass + "\n")
	e.Expect(xosPrompt, xosTimeout)
	e.Send("disable clipaging\n")
	e.Expect(xosPrompt, xosTimeout)
	e.Send("show configuration\n")

	config, _, _ := e.Expect(xosPrompt, xosTimeout)
	
	// Get rid of wandering CRs
	config = strings.Replace(config, "\r", "", -1)
	config = strings.TrimPrefix(config, "show configuration\n")
	config = xosPrompt.ReplaceAllString(config, "")

	return io.NopCloser(strings.NewReader(config)), nil
}

func init() {
	kinds[XOS] = fetchXOS
}
