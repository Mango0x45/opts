package opts

import "fmt"

type ErrBadOption rune

func (e ErrBadOption) Error() string {
	return fmt.Sprintf("unknown option ‘%c’", e)
}

type ErrNoArgument rune

func (e ErrNoArgument) Error() string {
	return fmt.Sprintf("expected argument for option ‘%c’", e)
}
