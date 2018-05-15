package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var output = make(map[io.Writer]bool)

func main() {
	stdout := flag.Bool("stdout", true, "enable stdout")
	stderr := flag.Bool("stderr", false, "enable stderr")
	flag.Parse()

	output[os.Stdout] = *stdout
	output[os.Stderr] = *stderr

	b := new(bytes.Buffer)
	output[b] = true

	for w, enabled := range output {
		if enabled {
			fmt.Fprintf(w, "some output\n")
		}
	}

	fmt.Printf("%d bytes in buffer\n", b.Len())
}
