package opts

import (
	"fmt"
	"testing"
)

func die(t *testing.T, name string, want, got any) {
	t.Fatalf("Expected %s to be ‘%s’ but got ‘%s’", name, want, got)
}

func assertGet(t *testing.T, args []string, fw, ow int, ew error) []Flag {
	flags, optind, err := Get(args, "abλc:dßĦ::")
	if err != ew {
		die(t, "err", ew, err)
	}
	if optind != ow {
		die(t, "optind", ow, optind)
	}
	if len(flags) != fw {
		die(t, "flags", fw, flags)
	}
	return flags
}

func TestNoArg(t *testing.T) {
	args := []string{}
	assertGet(t, args, 0, 0, nil)
}

func TestNoFlag(t *testing.T) {
	args := []string{"foo"}
	assertGet(t, args, 0, 1, nil)
}

func TestCNoArg(t *testing.T) {
	args := []string{"foo", "-c"}
	assertGet(t, args, 0, 0, NoArgumentError('c'))
}

func TestCWithArg(t *testing.T) {
	args := []string{"foo", "-c", "bar"}
	flags := assertGet(t, args, 1, 3, nil)
	if flags[0].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[0].Key)
	}
	if flags[0].Value != "bar" {
		die(t, "flags[0].Value", "bar", flags[0].Value)
	}
}

func TestCWithArgNoSpace(t *testing.T) {
	args := []string{"foo", "-cbar"}
	flags := assertGet(t, args, 1, 2, nil)
	if flags[0].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[0].Key)
	}
	if flags[0].Value != "bar" {
		die(t, "flags[0].Value", "bar", flags[0].Value)
	}
}

func TestANoArg(t *testing.T) {
	args := []string{"foo", "-a"}
	flags := assertGet(t, args, 1, 2, nil)
	if flags[0].Key != 'a' {
		die(t, "flags[0].Key", 'a', flags[0].Key)
	}
	if flags[0].Value != "" {
		die(t, "flags[0].Value", "", flags[0].Value)
	}
}

func TestAWithArg(t *testing.T) {
	args := []string{"foo", "-a", "bar"}
	flags := assertGet(t, args, 1, 2, nil)
	if flags[0].Key != 'a' {
		die(t, "flags[0].Key", 'a', flags[0].Key)
	}
	if flags[0].Value != "" {
		die(t, "flags[0].Value", "", flags[0].Value)
	}
}

func TestAAndCWithArg(t *testing.T) {
	args := []string{"foo", "-a", "-c", "bar"}
	flags := assertGet(t, args, 2, 4, nil)
	if flags[0].Key != 'a' {
		die(t, "flags[0].Key", 'a', flags[0].Key)
	}
	if flags[0].Value != "" {
		die(t, "flags[0].Value", "", flags[0].Value)
	}
	if flags[1].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[1].Key)
	}
	if flags[1].Value != "bar" {
		die(t, "flags[0].Value", "bar", flags[1].Value)
	}
}

func TestACWithArg(t *testing.T) {
	args := []string{"foo", "-ac", "bar"}
	flags := assertGet(t, args, 2, 3, nil)
	if flags[0].Key != 'a' {
		die(t, "flags[0].Key", 'a', flags[0].Key)
	}
	if flags[0].Value != "" {
		die(t, "flags[0].Value", "", flags[0].Value)
	}
	if flags[1].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[1].Key)
	}
	if flags[1].Value != "bar" {
		die(t, "flags[0].Value", "bar", flags[1].Value)
	}
}

func TestCAWithArg(t *testing.T) {
	args := []string{"foo", "-ca", "bar"}
	flags := assertGet(t, args, 1, 2, nil)
	if flags[0].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[0].Key)
	}
	if flags[0].Value != "a" {
		die(t, "flags[0].Value", "a", flags[0].Value)
	}
}

func TestBAfterDashDash(t *testing.T) {
	args := []string{"foo", "--", "-b"}
	assertGet(t, args, 0, 2, nil)
}

func TestCWithArgAfterDashDash(t *testing.T) {
	args := []string{"foo", "--", "-c", "bar", "baz"}
	assertGet(t, args, 0, 2, nil)
}

func TestCWithArgThenDAfterDashDash(t *testing.T) {
	args := []string{"foo", "-c", "bar", "baz", "--", "-d"}
	flags := assertGet(t, args, 1, 3, nil)
	if flags[0].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[0].Key)
	}
	if flags[0].Value != "bar" {
		die(t, "flags[0].Value", "bar", flags[0].Value)
	}
}

func TestCWithArgThenDAfterEmpty(t *testing.T) {
	args := []string{"foo", "-c", "bar", "baz", "", "-d"}
	flags := assertGet(t, args, 1, 3, nil)
	if flags[0].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[0].Key)
	}
	if flags[0].Value != "bar" {
		die(t, "flags[0].Value", "bar", flags[0].Value)
	}
}

func TestBChainedThrice(t *testing.T) {
	args := []string{"foo", "-bbb"}
	flags := assertGet(t, args, 3, 2, nil)
	for i := 0; i < 3; i++ {
		s := fmt.Sprintf("flags[%d].", i)
		if flags[i].Key != 'b' {
			die(t, s+"Key", 'b', flags[i].Key)
		}
		if flags[i].Value != "" {
			die(t, s+"Value", "", flags[i].Value)
		}
	}
}

func TestẞChainedTwice(t *testing.T) {
	args := []string{"foo", "-ßß"}
	flags := assertGet(t, args, 2, 2, nil)
	for i := 0; i < 2; i++ {
		s := fmt.Sprintf("flags[%d].", i)
		if flags[i].Key != 'ß' {
			die(t, s+"Key", 'ß', flags[i].Key)
		}
		if flags[i].Value != "" {
			die(t, s+"Value", "", flags[i].Value)
		}
	}
}

func TestΛAsArgToC(t *testing.T) {
	args := []string{"foo", "-c", "-λ"}
	flags := assertGet(t, args, 1, 3, nil)
	if flags[0].Key != 'c' {
		die(t, "flags[0].Key", 'c', flags[0].Key)
	}
	if flags[0].Value != "-λ" {
		die(t, "flags[0].Value", "-λ", flags[0].Value)
	}
}

func TestInvalidFlag(t *testing.T) {
	args := []string{"foo", "-X"}
	assertGet(t, args, 0, 0, BadOptionError('X'))
}

func TestInvalidFlagWithArg(t *testing.T) {
	args := []string{"foo", "-X", "bar"}
	assertGet(t, args, 0, 0, BadOptionError('X'))
}

func TestAAfterArg(t *testing.T) {
	args := []string{"foo", "bar", "-a"}
	assertGet(t, args, 0, 1, nil)
}

func TestXAfterDash(t *testing.T) {
	args := []string{"foo", "-", "-x"}
	assertGet(t, args, 0, 1, nil)
}

func TestĦWithSpaceAndArg(t *testing.T) {
	args := []string{"foo", "-Ħ", "bar"}
	flags := assertGet(t, args, 1, 2, nil)
	if flags[0].Key != 'Ħ' {
		die(t, "flags[0].Key", 'Ħ', flags[0].Key)
	}
	if flags[0].Value != "" {
		die(t, "flags[0].Value", "", flags[0].Value)
	}
}

func TestĦWithArg(t *testing.T) {
	args := []string{"foo", "-Ħbar"}
	flags := assertGet(t, args, 1, 2, nil)
	if flags[0].Key != 'Ħ' {
		die(t, "flags[0].Key", 'Ħ', flags[0].Key)
	}
	if flags[0].Value != "bar" {
		die(t, "flags[0].Value", "bar", flags[0].Value)
	}
}

func TestΛĦWithArg(t *testing.T) {
	args := []string{"foo", "-λĦbar"}
	flags := assertGet(t, args, 2, 2, nil)
	if flags[0].Key != 'λ' {
		die(t, "flags[0].Key", 'λ', flags[0].Key)
	}
	if flags[0].Value != "" {
		die(t, "flags[0].Value", "", flags[0].Value)
	}
	if flags[1].Key != 'Ħ' {
		die(t, "flags[1].Key", 'Ħ', flags[1].Key)
	}
	if flags[1].Value != "bar" {
		die(t, "flags[1].Value", "bar", flags[1].Value)
	}
}
