package aliasnimby

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"

	"github.com/mvdan/lint"
	"golang.org/x/tools/go/loader"
)

type Checker struct {
	prog *loader.Program
}

var _ lint.Checker = &Checker{}

func (c *Checker) Program(prog *loader.Program) {
	c.prog = prog
}

func (c *Checker) Check() ([]lint.Issue, error) {
	var issues []lint.Issue
	for _, pi := range c.prog.Imported {
		if pi.Errors != nil {
			fmt.Fprintf(os.Stderr, "Cannot check package: %s\n", pi.Pkg.Name())
			for _, err := range pi.Errors {
				fmt.Fprintf(os.Stderr, "\t%s\n", err)
			}
			os.Exit(1)
		}

		if !pi.TransitivelyErrorFree {
			fmt.Fprintf(os.Stderr, "Cannot check package %s: not error free\n", pi.Pkg.Name())
			os.Exit(1)
		}

		for _, file := range pi.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				typ, ok := n.(*ast.TypeSpec)
				if !ok {
					return true
				}
				if typ.Assign != 0 {
					issue := Issue{
						pos: n.Pos(),
						msg: fmt.Sprintf("%v is an alias", typ.Name),
					}
					issues = append(issues, issue)
				}
				return false
			})
		}
	}
	return issues, nil
}

type Issue struct {
	pos token.Pos
	msg string
}

var _ lint.Issue = Issue{}

func (i Issue) Pos() token.Pos {
	return i.pos
}

func (i Issue) Message() string {
	return i.msg
}
