package main

import (
	"fmt"
	"github.com/aliskhannn/go-cut/internal/cut"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	flags := cut.InitFlags()
	pflag.Parse()

	args := pflag.Args()
	if (flags.Fields == nil || len(*flags.Fields) == 0) && len(args) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: gocut [OPTIONS] -f fields [-d delimiter] [FILE...]")
		os.Exit(1)
	}

	if flags.Fields == nil || len(*flags.Fields) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "At least one field must be specified with -f")
		os.Exit(1)
	}

	cfg := cut.Config{
		Fields:    *flags.Fields,
		Delimiter: *flags.Delimiter,
		Separated: *flags.Separated,
	}

	files := args
	if len(files) == 0 {
		if err := cut.Process(os.Stdin, os.Stdout, cfg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		return
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Failed to open file:", f, err)
			os.Exit(1)
		}

		if err := cut.Process(file, os.Stdout, cfg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error processing file:", f, err)
		}

		// Close the file if it was opened.
		if file != nil {
			_ = file.Close()
		}
	}
}
