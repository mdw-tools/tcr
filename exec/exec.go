package exec

import (
	"log"
	"os/exec"
)

func Run(directory string, args ...string) (output string, err error) {
	command := exec.Command(args[0], args[1:]...)
	if directory != "" {
		command.Dir = directory
	}
	out, err := command.CombinedOutput()
	return string(out), err
}

func RunOrFatal(directory string, args ...string) string {
	output, err := Run(directory, args...)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
