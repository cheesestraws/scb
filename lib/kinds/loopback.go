package kinds

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

// Loopback is a virtual device that just echoes back what it's given
var Loopback = "loopback"

func fetchLoopback(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.ReadCloser, error) {
	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, "# device: %s\n", device)
	fmt.Fprintf(buf, "# user: %s\n", user)

	return io.NopCloser(buf), nil
}

func init() {
	kinds[Loopback] = fetchLoopback
}
