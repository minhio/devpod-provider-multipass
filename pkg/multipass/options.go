package multipass

import "io"

type Options struct {
	Env    []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type OptionSetter func(*Options)

func Env(env []string) OptionSetter {
	return func(args *Options) {
		args.Env = env
	}
}

func Stdin(stdin io.Reader) OptionSetter {
	return func(args *Options) {
		args.Stdin = stdin
	}
}

func Stdout(stdout io.Writer) OptionSetter {
	return func(args *Options) {
		args.Stdout = stdout
	}
}

func Stderr(stdout io.Writer) OptionSetter {
	return func(args *Options) {
		args.Stdout = stdout
	}
}
