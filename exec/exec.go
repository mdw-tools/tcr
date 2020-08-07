package exec

import (
	"bytes"
	"io"
	"log"
	"os/exec"
	"strings"
)

type option func(*exec.Cmd)

func At(dir string) option {
	return func(command *exec.Cmd) {
		command.Dir = dir
	}
}

func Out(writers ...io.Writer) option {
	return func(command *exec.Cmd) {
		command.Stdout = io.MultiWriter(append(writers, command.Stdout)...)
		command.Stderr = command.Stdout
	}
}

func Args(args ...string) option {
	return func(command *exec.Cmd) {
		command.Args = append(command.Args, args...)
	}
}

func Run(program string, options ...option) (output string, err error) {
	buffer := new(bytes.Buffer)
	command := exec.Command(program)
	command.Stdout = buffer
	command.Stderr = buffer
	for _, option := range options {
		option(command)
	}
	err = command.Run()
	return strings.TrimSpace(buffer.String()), err
}
func RunFatal(program string, options ...option) string {
	output, err := Run(program, options...)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
