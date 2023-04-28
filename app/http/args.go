package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type (
	Args struct {
		Config string `short:"c" long:"config" description:"the config line" required:"true"`
	}
)

func parseArgs() Args {
	var args Args
	_, err := flags.ParseArgs(&args, os.Args)
	if err != nil {
		fmt.Println("parse arguments error:" + err.Error())
		os.Exit(-1)
	}
	return args
}
