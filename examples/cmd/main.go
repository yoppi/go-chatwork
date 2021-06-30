package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lettenj61/go-chatwork/examples"
)

var apiKey, cmd string

func getOptions() (*flag.FlagSet, error) {
	p := flag.NewFlagSet("examples", flag.ExitOnError)
	p.StringVar(&apiKey, "apikey", "", "Chatwork API key")
	p.StringVar(&cmd, "command", "", "Command name")

	err := p.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}
	if apiKey == "" {
		return p, fmt.Errorf("argument missing: -apikey (string)")
	}
	return p, nil
}

func main() {
	p, err := getOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: flag=%#v, error=%v\n", p, err)
		os.Exit(1)
	}

	switch cmd {
	case "me":
		examples.Me(apiKey)
	case "my":
		examples.My(apiKey)
	default:
		fmt.Println("missing or unknown example:", cmd)
		p.Usage()
	}
}
