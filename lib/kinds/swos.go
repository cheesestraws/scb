package kinds

import (
	"context"
	"fmt"
	"io"

	"net/http"

	"github.com/icholy/digest"
)

// SWOS is a kind that polls Mikrotik SWOS devices
var SWOS = "swos"

func fetchSWOS(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.ReadCloser, error) {
	client := &http.Client{
		Transport: &digest.Transport{
			Username: user,
			Password: pass,
		},
	}

	req, err := http.NewRequestWithContext(ctx, "GET", "http://"+device+"/backup.swb", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got code %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func init() {
	kinds[SWOS] = fetchSWOS
}
