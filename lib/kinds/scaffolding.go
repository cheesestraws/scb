package kinds

import (
	"context"
	"errors"
	"io"
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
