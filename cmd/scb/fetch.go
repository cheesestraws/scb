package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cheesestraws/scb/lib/kinds"
)

func fetch(args []string) error {
	if len(args) < 4 {
		return ErrTooFewParameters
	}

	ctx := context.Background()

	r, e := kinds.Fetch(ctx, args[0], args[1], args[2], args[3], nil)
	if e != nil {
		return e
	}

	io.Copy(os.Stdout, r)
	r.Close()
	fmt.Printf("\n")

	return nil
}
