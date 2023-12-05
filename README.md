## Update 6th December, 2023

The v2 of this library has been released, and you should probably use
that instead:

```sh
$ go get -u git.sr.ht/~mango/opts/v2
```

# Opts

Opts is a simple Go library implementing unicode-aware getopt(3)- and
getopt_long(3) flag parsing in Go.  For details, check out the godoc
documentation.

Note that unlike the `getopt()` C function, the ‘:’ and ‘?’ flags are not
returned on errors — the errors are instead returned via the `err` return
value of `opts.Get()` and `opts.GetLong()`.  Additionally, a leading ‘:’
in the opt-string provided to `opts.Get()` is not supported; you are
responsible for your own I/O.

## Example Usage

The following demonstrates an example usage of the `opts.Get()` function…

```go
package main

import (
	"fmt"
	"os"

	"git.sr.ht/~mango/opts/v2"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-ßλ] [-a arg]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	flags, rest, err := opts.Get(os.Args, "a:ßλ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		usage()
	}

	for _, f := range flags {
		switch f.Key {
		case 'a':
			fmt.Println("-a given with argument", f.Value)
		case 'ß':
			fmt.Println("-ß given")
		case 'λ':
			fmt.Println("-λ given")
		}
	}

	fmt.Println("The remaining arguments are:", rest)
}
```

…and the following demonstrates an example usage of the `opts.GetLong()`
function:

```go
package main

import (
	"fmt"
	"os"

	"git.sr.ht/~mango/opts"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-ßλ] [-a arg] [--no-short]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	// The fourth long-option has no short-option equivalent
	flags, rest, err := opts.GetLong(os.Args, []opts.LongOpt{
		{Short: 'a', Long: "add", Arg: opts.Required},
		{Short: 'ß', Long: "sheiße", Arg: opts.None},
		{Short: 'λ', Long: "λεωνίδας", Arg: opts.None},
		{Short: -1, Long: "no-short", Arg: opts.None},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		usage()
	}

	for _, f := range flags {
		switch f.Key {
		case 'a':
			fmt.Println("-a or --add given with argument", f.Value)
		case 'ß':
			fmt.Println("-ß or --sheiße given")
		case 'λ':
			fmt.Println("-λ or --λεωνίδας given")
		case -1:
			fmt.Println("--no-short given")
		}
	}

	// The remaining arguments
	fmt.Println("The remaining arguments are:", rest)
}
```
