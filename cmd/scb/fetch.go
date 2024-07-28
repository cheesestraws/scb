package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cheesestraws/scb/lib/kinds"
)

func fetch() error {
	args := os.Args
	if len(args) < 5 {
		return ErrTooFewParameters
	}

	ctx := context.Background()

	r, e := kinds.Fetch(ctx, os.Args[1], os.Args[2], os.Args[3], os.Args[4], nil)
	if e != nil {
		return e
	}

	io.Copy(os.Stdout, r)
	r.Close()
	fmt.Printf("\n")

	return nil
}
