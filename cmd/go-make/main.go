package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mdwhatcott/tcr/exec"
	"github.com/mdwhatcott/tcr/exec/git"
	"github.com/mdwhatcott/tcr/gotest"
)

var Version = "dev"

const usageInfo = `
	This program runs the following go utilities in the working directory:

	- go version
	- go mod tidy
	- go fmt ./...
	- go vet ./...
	- go test {args}

	The go test command {args} default to '-cover ./...' if no args are provided.
`

func main() {
	log.SetFlags(log.Ltime | log.Llongfile)
	flags := flag.NewFlagSet("go-make", flag.ContinueOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(
			flags.Output(),
			"Usage of go-make (version: %s)\n%s",
			Version,
			usageInfo,
		)
		flags.PrintDefaults()
	}
	err := flags.Parse(os.Args)
	if err == flag.ErrHelp {
		os.Exit(1)
	}
	args := argsForGoTest(os.Args[1:])
	path := git.RepositoryRoot()

	fmt.Println(exec.RunFatal("go version"))
	output := run(path,
		"go mod tidy",
		"go fmt ./...",
		"go vet ./...",
		"go test "+args,
	)
	fmt.Println("----")
	fmt.Println(strings.TrimSpace(gotest.Format(output)))
}

func argsForGoTest(rawArgs []string) string {
	args := strings.Join(rawArgs, " ")
	if len(args) == 0 {
		args = "-cover ./..."
	}
	return args
}

func run(working string, commands ...string) string {
	b := new(bytes.Buffer)
	for _, command := range commands {
		fmt.Println(command)
		b.WriteString(exec.RunFatal(command, exec.At(working), exec.Out(os.Stdout)))
		b.WriteString("\n")
	}
	return b.String()
}
