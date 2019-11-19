// TCR: test && commit || revert
// codified by Kent Beck
// https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mdwhatcott/tcr/exec"
)

func main() {
	resetStopwatch()
	printSummary(timed(TCR))
}
func resetStopwatch() {
	_, _ = http.Get("http://localhost:7890/stopwatch/reset")
}
func TCR() {
	_ = Test() && Commit() || Revert()
}
func timed(action func()) time.Duration {
	started := time.Now()
	action()
	return time.Since(started)
}
func Test() bool {
	printBanner("-- TEST --")
	return executeTests()
}
func printBanner(banner string) {
	fmt.Println()
	fmt.Println(banner)
	fmt.Println()
}
func executeTests() bool {
	output, err := exec.Run(getRepositoryRoot(), "make")
	fmt.Println(strings.TrimSpace(output))
	return err == nil
}
func getRepositoryRoot() string {
	return strings.TrimSpace(exec.RunOrFatal("", "git", "rev-parse", "--show-toplevel"))
}
func Commit() bool {
	printBanner("-- COMMIT --")
	commitChanges()
	printBanner("-- OK --")
	return true
}
func commitChanges() {
	_, _ = exec.Run("", "git", "add", ".")
	output, _ := exec.Run("", "git", "commit", "-m", "tcr")
	fmt.Println(strings.TrimSpace(output))
}
func Revert() bool {
	printBanner("-- REVERT --")
	revertState()
	printBanner("-- ERROR --")
	return true
}
func revertState() {
	revertToPreviousCommit()
	clearClipboardContents()
}
func revertToPreviousCommit() {
	fmt.Println(exec.RunOrFatal("", "git", "clean", "-df"))
	fmt.Println(exec.RunOrFatal("", "git", "reset", "--hard"))
}
func clearClipboardContents() {
	fmt.Println(exec.RunOrFatal("", "pbcopy", "less is more"))
}
func printSummary(duration time.Duration) {
	fmt.Println("Location:", workingDirectory())
	fmt.Println("Duration:", duration)
	fmt.Println("Commits:", exec.GetTCRCommitCount())
}
func workingDirectory() string {
	dir, _ := os.Getwd()
	return dir
}
