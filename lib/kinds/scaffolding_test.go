package kinds

import (
	"context"
	"io"
	"strings"
	"testing"
)

func TestDispatch(t *testing.T) {
	dummyFetchFunc := func(ctx context.Context, device string, user string, pass string, reserved interface{}) (io.Reader, error) {
		return strings.NewReader("honk honk"), nil
	}

	kinds["dummy"] = dummyFetchFunc

	r, e := Fetch(context.Background(), "dummy", "", "", "", nil)
	if e != nil {
		t.Errorf("dispatch error: %v", e)
	}

	bs, e := io.ReadAll(r)
	if e != nil {
		t.Errorf("read error: %v", e)
	}

	if string(bs) != "honk honk" {
		t.Errorf("garbled config returned")
	}
}
