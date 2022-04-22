package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go.husin.dev/differ"
)

var (
	useColor   = flag.Bool("color", false, "show colors")
	useUnified = flag.Bool("unified", false, "show unified (print the same lines too)")
)

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) != 2 {
		errormsg("error: expected 2 arguments, found %d\n", len(files))
		os.Exit(1)
	}

	rawLeft, err := os.ReadFile(files[0])
	if err != nil {
		errormsg("error: parsing '%s': %v\n", files[0], err.Error())
	}
	rawRight, err := os.ReadFile(files[1])
	if err != nil {
		errormsg("error: parsing '%s': %v\n", files[1], err.Error())
	}

	left := strings.Split(string(rawLeft), "\n")
	right := strings.Split(string(rawRight), "\n")

	diffs := differ.Diff[string](left, right)

	p := differ.NewPrinter(differ.StrconvString)

	if *useColor {
		p.Configure(differ.PrintColor)
	}
	if *useUnified {
		p.Configure(differ.PrintUnified)
	} else {
		p.Configure(differ.PrintNumber)
	}

	p.Add(diffs...)

	fmt.Printf(p.String())
}

func errormsg(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
}
