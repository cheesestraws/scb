package kinds

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/jlaffaye/ftp"
)

// APCFTP is a kind that pulls configuration from APC rPDUs over FTP
var APCFTP = "apcftp"

func fetchAPCFTP(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.ReadCloser, error) {
	c, err := ftp.Dial(fmt.Sprintf("%s:21", device),
		ftp.DialWithContext(ctx),
		ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}
	defer c.Quit()

	err = c.Login(user, pass)
	if err != nil {
		return nil, err
	}

	resp, err := c.Retr("/config.ini")
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	bs, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(bs)
	return io.NopCloser(buf), nil
}

func init() {
	kinds[APCFTP] = fetchAPCFTP
}
