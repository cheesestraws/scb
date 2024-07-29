package kinds

import (
	"context"
	"io"
	"regexp"
	"strings"
	"time"
)

// RouterOS is for mikrotik routeros things
var RouterOS = "routeros"

var rosTimeout = 10 * time.Second
var rosUserPrompt = regexp.MustCompile(regexp.QuoteMeta("Login:"))
var rosPasswordPrompt = regexp.MustCompile(regexp.QuoteMeta("Password:"))
var rosPrompt = regexp.MustCompile(`(?U)\[\w+@\S+(\s+\S+)*\]\s?>\s+$`)
var rosFirstLine = regexp.MustCompile("^[^\n]*\\n")


func fetchRouterOS(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.ReadCloser, error) {
	e, _, err := telnetSpawn(device+":23", -1)
	if err != nil {
		return nil, err
	}
	defer e.Close()
	
	e.Expect(rosUserPrompt, rosTimeout)
	e.Send(user + "+cet\n") // +cet = dumb terminal, no colours, don't probe terminal
	e.Expect(rosPasswordPrompt, rosTimeout)
	e.Send(pass + "\n")
	e.Expect(rosPrompt, rosTimeout)
	e.Send("/export\r\n")

	config, _, _ := e.Expect(rosPrompt, rosTimeout)

	// Get rid of wandering CRs
	config = strings.Replace(config, "\r", "", -1)
	config = rosFirstLine.ReplaceAllString(config, "")
	config = rosPrompt.ReplaceAllString(config, "")


	return io.NopCloser(strings.NewReader(config)), nil
}

func init() {
	kinds[RouterOS] = fetchRouterOS
}
