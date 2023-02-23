package main

import (
	"fmt"
	"os"

	"github.com/jamesliu96/faux"
)

func main() {
	var key []byte
	if len(os.Args) > 1 {
		key = []byte(os.Args[1])
	}
	if len(key) == 0 {
		fmt.Fprintln(os.Stderr, "warning: zero-length key")
	}
	fmt.Fprintf(os.Stderr, "\"%s\"=[%x]\n", key, key)
	if err := faux.Faux(os.Stdin, os.Stdout, key); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
