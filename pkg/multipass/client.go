package multipass

import (
	"io"
)

type options struct {
	environ []string
	stdin   io.Reader
	stdout  io.Writer
	stderr  io.Writer
}

type optionSetter func(*options)

type client struct {
	executablePath string
	environ        []string
	stdin          io.Reader
	stdout         io.Writer
	stderr         io.Writer
}

func (c *client) GetInstance(name string) (*instance, error) {
	listResult, err := c.List()
	if err != nil {
		return nil, err
	}

	for _, item := range listResult.List {
		if item.Name == name {
			inst := &instance{
				Name:  item.Name,
				State: item.State,
				Ipv4:  item.Ipv4,
			}
			return inst, nil
		}
	}

	return nil, &instanceNotFound{name: name}
}

func NewClient(executablePath string, optionSetters ...optionSetter) *client {
	opts := &options{}

	for _, setter := range optionSetters {
		setter(opts)
	}

	return &client{
		executablePath: executablePath,
		environ:        opts.environ,
		stdin:          opts.stdin,
		stdout:         opts.stdout,
		stderr:         opts.stderr,
	}
}

func SetClientEnviron(environ []string) optionSetter {
	return func(args *options) {
		args.environ = environ
	}
}

func SetClientStdin(stdin io.Reader) optionSetter {
	return func(args *options) {
		args.stdin = stdin
	}
}

func SetClientStdout(stdout io.Writer) optionSetter {
	return func(args *options) {
		args.stdout = stdout
	}
}

func SetClientStderr(stdout io.Writer) optionSetter {
	return func(args *options) {
		args.stdout = stdout
	}
}
