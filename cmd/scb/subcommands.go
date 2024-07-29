package main

import "errors"

var ErrTooFewParameters = errors.New("too few parameters")
var ErrBadSubcommand = errors.New("bad subcommand")

var subcommands = map[string]func() error{
	"fetch": fetch,
}

func extractSubcommand(args []string) (string, []string, error) {
	if len(args) < 2 {
		return "", nil, ErrTooFewParameters
	}

	subcommand := args[1]
	rest := []string{args[0]}
	rest = append(rest, args[2:]...)

	return subcommand, rest, nil
}

func runSubcommand(subcommand string) error {
	f, ok := subcommands[subcommand]
	if !ok {
		return ErrBadSubcommand
	}
	return f()
}