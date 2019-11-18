package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TCR: test && commit || revert
// codified by Kent Beck
// https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864

func main() {
	started := time.Now()
	_ = Test() && Commit() || Revert()
	printBanner(time.Since(started).String())
}

func Test() bool {
	resetStopwatch()
	printWorkingDirectory()
	printBanner("-- TEST --")
	return executeTests()
}
func resetStopwatch() {
	_, _ = http.Get("http://localhost:7890/stopwatch/reset")
}
func printWorkingDirectory() {
	fmt.Println()
	dir, _ := os.Getwd()
	fmt.Println(dir)
}
func printBanner(banner string) {
	fmt.Println()
	fmt.Println(banner)
}
func executeTests() bool {
	output, err := execute(exec.Command("make"), getRepositoryRoot())
	fmt.Println(output)
	return err == nil
}
func getRepositoryRoot() string {
	return strings.TrimSpace(executeOrFatal(exec.Command("git", "rev-parse", "--show-toplevel"), ""))
}
func execute(command *exec.Cmd, directory string) (output string, err error) {
	if directory != "" {
		command.Dir = directory
	}
	out, err := command.CombinedOutput()
	return string(out), err
}
func executeOrFatal(command *exec.Cmd, directory string) string {
	output, err := execute(command, directory)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func Commit() bool {
	printBanner("-- COMMIT --")
	_, _ = execute(exec.Command("git", "add", "."), "")
	fmt.Println(execute(exec.Command("git", "commit", "-m", "tcr"), ""))
	printBanner(fmt.Sprintf("TCR commit count: %d", getTCRCommitCount()))
	printBanner("-- OK --")
	return true
}
func getTCRCommitCount() int {
	output := executeOrFatal(exec.Command("git", "log", "--oneline"), "")
	lines := strings.Split(output, "\n")
	var count = 0
	for _, line := range lines {
		if strings.HasSuffix(line, " tcr") {
			count++
		} else {
			break
		}
	}
	return count
}

func Revert() bool {
	printBanner("-- REVERT --")
	fmt.Println(executeOrFatal(exec.Command("git", "clean", "-df"), ""))
	fmt.Println(executeOrFatal(exec.Command("git", "reset", "--hard"), ""))
	fmt.Println(executeOrFatal(exec.Command("pbcopy", "less is more"), ""))
	printBanner("-- ERROR --")
	return true
}
