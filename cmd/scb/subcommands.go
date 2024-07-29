package main

import (
	"errors"
)

var ErrTooFewParameters = errors.New("too few parameters")
var ErrBadSubcommand = errors.New("bad subcommand")

var subcommands = map[string]func([]string) error{
	"fetch": fetch,
	"download": download,
}

func extractSubcommand(args []string) (string, []string, error) {
	if len(args) < 1 {
		return "", nil, ErrTooFewParameters
	}

	subcommand := args[0]
	rest := args[1:]

	return subcommand, rest, nil
}

func runSubcommand(subcommand string, args []string) error {
	f, ok := subcommands[subcommand]
	if !ok {
		return ErrBadSubcommand
	}
	return f(args)
}
