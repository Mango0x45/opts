// Package opts implements unicode-aware getopt(3)- and getopt_long(3)
// flag parsing.
//
// The opts package aims to provide as simple an API as possible.  If
// your usecase requires more advanced argument-parsing or a more robust
// API, this may not be the ideal package for you.
//
// While the opts package aims to closely follow the POSIX getopt(3) and
// GNU getopt_long(3) C functions, there are some notable differences.
// This package properly supports unicode flags, but also does not
// support a leading ‘:’ in the [Get] function’s option string; all
// user-facing I/O is delegrated to the caller.
package opts

// ArgMode represents the whether or not a long-option takes an argument.
type ArgMode int

// These tokens can be used to specify whether or not a long-option takes
// an argument.
const (
	None     ArgMode = iota // long opt takes no argument
	Required                // long opt takes an argument
	Optional                // long opt optionally takes an argument
)

// Flag represents a parsed command-line flag.  Key corresponds to the
// rune that was passed on the command-line, and Value corresponds to the
// flags argument if one was provided.  In the case of long-options Key
// will map to the corresponding short-code, even if a long-option was
// used.
type Flag struct {
	Key   rune   // the flag that was passed
	Value string // the flags argument
}

// LongOpt represents a long-option to attempt to parse.  All long
// options have a short-hand form represented by Short and a long-form
// represented by Long.  Arg is used to represent whether or not a takes
// an argument.
//
// In the case that you want to parse a long-option which doesn’t have a
// short-hand form, you can set Short to a negative integer.
type LongOpt struct {
	Short rune
	Long  string
	Arg   ArgMode
}

// Get parses the command-line arguments in args according to optstr.
// Unlike POSIX-getopt(3), a leading ‘:’ in optstr is not supported and
// will be ignored and no I/O is ever performed.
//
// Get will look for the flags listed in optstr (i.e., it will look for
// ‘-a’, ‘-ß’, and ‘λ’ given optstr == "aßλ").  The optstr need not be
// sorted in any particular order.  If an option takes a required
// argument, it can be suffixed by a colon.  If an option takes an
// optional argument, it can be suffixed by two colons.  As an example,
// optstr == "a::ßλ:" will search for ‘-a’ with an optional argument,
// ‘-ß’ with no argument, and ‘-λ’ with a required argument.
//
// A successful parse returns the flags in the flags slice and the index
// of the first non-option argument in optind.  In the case of failure,
// err will be one of [BadOptionError] or [NoArgumentError].
func Get(args []string, optstr string) (flags []Flag, optind int, err error) {
	argmap := make(map[rune]ArgMode)

	rs := []rune(optstr)
	if rs[0] == ':' {
		rs = rs[1:]
	}
	for len(rs) > 0 {
		switch r := rs[0]; {
		case len(rs) > 2 && rs[1] == ':' && rs[2] == ':':
			argmap[r] = Optional
			rs = rs[3:]
		case len(rs) > 1 && rs[1] == ':':
			argmap[r] = Required
			rs = rs[2:]
		default:
			argmap[r] = None
			rs = rs[1:]
		}
	}

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if len(arg) == 0 || arg == "-" || arg[0] != '-' {
			optind = i
			return
		} else if arg == "--" {
			optind = i + 1
			return
		}

		rs := []rune(arg[1:])
		for j, r := range rs {
			var s string
			am, ok := argmap[r]

			switch {
			case !ok:
				return nil, 0, BadOptionError(r)
			case am != None && j < len(rs)-1:
				s = string(rs[j+1:])
			case am == Required:
				i++
				if i >= len(args) {
					return nil, 0, NoArgumentError(r)
				}
				s = args[i]
			default:
				flags = append(flags, Flag{Key: r})
				continue
			}

			flags = append(flags, Flag{r, s})
			break
		}
	}

	optind = len(args)
	return
}
