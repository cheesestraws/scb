package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cheesestraws/scb/lib/kinds"
)

func download() error {
	args := os.Args
	if len(args) < 5 {
		return ErrTooFewParameters
	}

	ctx := context.Background()
	return downloadOne(ctx, os.Args[1], os.Args[2], os.Args[3], os.Args[4])
}

func downloadOne(ctx context.Context, kind string, device string, user string, pass string) error {
	r, e := kinds.Fetch(ctx, kind, device, user, pass, nil)
	if e != nil {
		return e
	}

	fn := fmt.Sprintf("%s-%s-%s.conf", device, kind, time.Now().Format("200601021504"))
	f, e := os.Create(fn)
	if e != nil {
		return e
	}
	defer f.Close()

	io.Copy(f, r)
	r.Close()
	fmt.Fprintf(f, "\n")

	return nil

}
