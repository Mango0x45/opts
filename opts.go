package opts

type ArgMode int

const (
	None ArgMode = iota
	Required
	Optional
)

type Flag struct {
	Key   rune
	Value string
}

type LongOpt struct {
	Short rune
	Long  string
	Arg   ArgMode
}

func Get(args []string, optstr string) (flags []Flag, optind int, err error) {
	argmap := make(map[rune]bool)

	rs := []rune(optstr)
	for i, r := range rs {
		if r != ':' {
			argmap[r] = false
		} else if i > 0 {
			argmap[rs[i-1]] = true
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
			argp, ok := argmap[r]

			switch {
			case !ok:
				return nil, 0, ErrBadOption(r)
			case argp && j < len(rs)-1:
				s = string(rs[j+1:])
			case argp:
				i++
				if i >= len(args) {
					return nil, 0, ErrNoArgument(r)
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
