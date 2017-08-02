package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bradleyfalzon/aliasnimby"
	"github.com/kisielk/gotool"
	"golang.org/x/tools/go/loader"
)

func main() {
	flag.Parse()

	// Use gotool to default blank import path to "." and handle recursion
	paths := gotool.ImportPaths(flag.Args())

	var conf loader.Config
	if _, err := conf.FromArgs(paths, true); err != nil {
		fmt.Fprintf(os.Stderr, "Could not check %v: %s\n", os.Args[1:], err)
		os.Exit(1)
	}

	prog, err := conf.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not check %v: %s\n", os.Args[1:], err)
		os.Exit(1)
	}

	var ok = true

	var checker aliasnimby.Checker
	checker.Program(prog)

	issues, err := checker.Check()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	for _, issue := range issues {
		ok = false
		// TODO add relative path not abs
		fmt.Fprintf(os.Stderr, "%s: %v\n", prog.Fset.Position(issue.Pos()), issue.Message())
	}

	if !ok {
		fmt.Fprint(os.Stderr, "Argh fark, it's in my backyard!\n")
		os.Exit(2)
	}
}
