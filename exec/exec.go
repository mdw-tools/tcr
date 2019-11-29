package exec

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func Run(directory string, args ...string) (output string, err error) {
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

func RunOrFatal(directory string, args ...string) string {
	output, err := Run(directory, args...)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
