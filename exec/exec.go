package exec

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
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
	return buffer.String(), err
}
func RunFatal(program string, options ...option) string {
	output, err := Run(program, options...)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func RunOld(directory string, args ...string) (output string, err error) {
	buffer := new(bytes.Buffer)
	writer := io.MultiWriter(os.Stdout, buffer)
	command := exec.Command(args[0], args[1:]...)
	command.Stdout = writer
	command.Stderr = writer
	if directory != "" {
		command.Dir = directory
	}
	err = command.Run()
	return buffer.String(), err
}

func RunOrFatalOld(directory string, args ...string) string {
	output, err := RunOld(directory, args...)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
