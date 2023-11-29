# Opts

Opts is a simple Go library implementing unicode-aware getopt(3)- and
getopt_long(3) flag parsing in Go.  For details, check out the godoc
documentation.

Note that unlike the `getopt()` C function, the ‘:’ and ‘?’ flags are not
returned on errors — the errors are instead returned via the `err` return value
of `opts.Get()` and `opts.GetLong()`.  Additionally, a leading ‘:’ in the
opt-string provided to `opts.Get()` is not supported; you are responsible for
your own I/O.

## Example Usage

The following demonstrates an example usage of the `opts.Get()` function:

```go
package main

import (
    "fmt"
    "os"

    "git.sr.ht/~mango/opts"
)

func usage() {
    fmt.Fprintf(os.Stderr, "Usage: %s [-ßλ] [-a arg]\n", os.Args[0])
    os.Exit(1)
}

func main() {
    flags, optind, err := opts.Get(os.Args, "a:ßλ")
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
        usage()
    }

    for _, opt := range opts {
        switch opt.Key {
        case 'a':
            fmt.Println("-a given with argument", opt.Value)
        case 'ß':
            fmt.Println("-ß given")
        case 'λ':
            fmt.Println("-λ given")
        }
    }

    // The remaining arguments
    rest := os.Args[optind:]
}
```
