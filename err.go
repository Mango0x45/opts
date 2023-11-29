package opts

import "fmt"

// A BadOptionError describes an option that the user attempted to pass
// which the developer did not register.
type BadOptionError rune

func (e BadOptionError) Error() string {
	return fmt.Sprintf("unknown option ‘%c’", e)
}

// A NoArgumentError describes an option that the user attempted to pass
// without an argument, which required an argument.
type NoArgumentError rune

func (e NoArgumentError) Error() string {
	return fmt.Sprintf("expected argument for option ‘%c’", e)
}
