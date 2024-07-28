package kinds

import (
	"context"
	"errors"
	"io"
)

// Loopback is a virtual device that just echoes back what it's given
var Everunready = "everunready"

func fetchEverunready(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.Reader, error) {
	return errors.New("fail")
}

func init() {
	kinds[Everunready] = fetchEverunready
}
