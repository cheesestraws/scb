package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	toml "github.com/pelletier/go-toml/v2"

	"github.com/cheesestraws/scb/lib/kinds"
)

func download(args []string) error {
	if len(args) < 4 {
		return ErrTooFewParameters
	}

	ctx := context.Background()
	return downloadOne(ctx, args[0], args[1], args[2], args[3])
}

type device struct {
	Kind string
	User string
	Password string
}

func checkToml(args []string) error {
	if len(args) < 1 {
		return ErrTooFewParameters
	}
	
	f, e := os.Open(args[0])
	if e != nil {
		fmt.Printf("%+v", e)
		os.Exit(1)
		return e
	}
	
	m := make(map[string]device)
	d := toml.NewDecoder(f)
	e = d.Decode(&m)
	if e != nil {
		fmt.Printf("%+v", e)
		os.Exit(1)
		return e
	}

	os.Exit(0)
	return nil
}

func downloadAll(args []string) error {
	if len(args) < 1 {
		return ErrTooFewParameters
	}
	
	f, e := os.Open(args[0])
	if e != nil {
		return e
	}
	
	m := make(map[string]device)
	d := toml.NewDecoder(f)
	e = d.Decode(&m)
	if e != nil {
		return e
	}
	
	ctx := context.Background()

	for k, v := range m {
		err := downloadOne(ctx, v.Kind, k, v.User, v.Password)
		if err != nil{
			fmt.Printf("%+v\n", err)
		}
	}

	return nil
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