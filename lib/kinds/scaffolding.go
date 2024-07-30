package kinds

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/cheesestraws/telnet"
	expect "github.com/google/goexpect"
)

var ErrNoSuchKind = errors.New("no such kind of device")

type FetchFunc func(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.ReadCloser, error)

var kinds = make(map[string]FetchFunc)

func Fetch(ctx context.Context, kind string, device string, user string,
	pass string, reserved interface{}) (io.ReadCloser, error) {

	f := kinds[kind]
	if f == nil {
		return nil, ErrNoSuchKind
	}

	return f(ctx, device, user, pass, reserved)
}

func telnetSpawn(addr string, timeout time.Duration, opts ...expect.Option) (expect.Expecter, <-chan error, error) {
	conn, err := telnet.Dial("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	resCh := make(chan error)

	return expect.SpawnGeneric(&expect.GenOptions{
		In:  conn,
		Out: conn,
		Wait: func() error {
			return <-resCh
		},
		Close: func() error {
			close(resCh)
			return conn.Close()
		},
		Check: func() bool { return true },
	}, timeout, opts...)
}
