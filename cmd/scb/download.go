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

	if *into != "." {
		err := os.Chdir(*into)
		if err != nil {
			return err
		}
	}

	ctx := context.Background()
	return downloadOne(ctx, args[0], args[1], args[2], args[3], false)
}

type device struct {
	Kind     string
	User     string
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

	if *into != "." {
		err := os.Chdir(*into)
		if err != nil {
			return err
		}
	}
	
	os.Mkdir("log", 0755)
	resfn := fmt.Sprintf("log/results-scb-%s.out", time.Now().Format("200601021504"))
	res, e := os.Create(resfn)
	if e != nil {
		return e
	}
	defer res.Close()

	ctx := context.Background()

	for k, v := range m {
		err := downloadOne(ctx, v.Kind, k, v.User, v.Password, true)
		if err != nil {
			fmt.Printf("%s: %+v\n", k, err)
			fmt.Fprintf(res, "%s: %+v\n", k, err)
		}
	}
	
	fmt.Fprintf(res, "ok\n")

	return nil
}

func downloadOne(ctx context.Context, kind string, device string, user string, pass string, mkdir bool) error {
	r, e := kinds.Fetch(ctx, kind, device, user, pass, nil)
	if e != nil {
		return e
	}

	fn := fmt.Sprintf("%s-%s-%s.conf", device, kind, time.Now().Format("200601021504"))
	
	if mkdir {
		os.Mkdir(device, 0755)
		fn = device+"/"+fn
	}

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
