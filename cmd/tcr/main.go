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
	runner := new(Runner)
	runner.TCR()
	fmt.Println(runner)
}

type Runner struct {
	started time.Time
	stopped time.Time

	testReport   string
	gitReport    string
	finalReport  *strings.Builder

	commitCount int
	testsPassed bool
}

func (this *Runner) TCR() {
	defer this.resetStopWatch()

	_ = this.Test() &&
		this.Commit() ||
		this.Revert()
}

func (this *Runner) resetStopWatch() {
	this.commitCount = exec.GetTCRCommitCount()
	URL := fmt.Sprintf(
		"http://localhost:7890/stopwatch/reset?commits=%d&passed=%t",
		this.commitCount,
		this.testsPassed,
	)
	_, _ = http.Get(URL)
}

func (this *Runner) Test() bool {
	this.start()
	defer this.stop()

	output, err := exec.Run(getRepositoryRoot(), "make")
	this.testReport = strings.TrimSpace(output)
	this.testsPassed = err == nil
	return this.testsPassed

}
func (this *Runner) start() { this.started = time.Now() }
func (this *Runner) stop()  { this.stopped = time.Now() }

func (this *Runner) Commit() bool {
	_, _ = exec.Run("", "git", "add", ".")
	output, _ := exec.Run("", "git", "commit", "-m", "tcr")
	this.gitReport = strings.TrimSpace(output)
	return true
}

func (this *Runner) Revert() bool {
	this.gitReport += exec.RunOrFatal("", "git", "clean", "-df")
	this.gitReport += exec.RunOrFatal("", "git", "reset", "--hard")
	this.gitReport += exec.RunOrFatal("", "pbcopy", "less is more")
	return true
}

func (this *Runner) String() string {
	this.finalReport = new(strings.Builder)
	this.printSummary()
	this.printReport(this.gitReport)
	this.printBanner("-- TEST --")
	this.printReport(this.testReport)
	return this.finalReport.String()
}

func (this *Runner) printReport(report string) {
	this.finalReport.WriteString(report)
}
func (this *Runner) printSummary() {
	this.printBanner("-- SUMMARY --")
	this.printData("Commits", this.commitCount)
	this.printData("Duration", this.stopped.Sub(this.started))
	this.printData("Location", workingDirectory())
	this.printPassOrFail()
}
func (this *Runner) printBanner(banner string) {
	this.finalReport.WriteString("\n\n")
	this.finalReport.WriteString(banner)
	this.finalReport.WriteString("\n\n")
}
func (this *Runner) printData(name string, data interface{}) {
	fmt.Fprintln(this.finalReport, name+":", data)
}
func workingDirectory() string {
	dir, _ := os.Getwd()
	return dir
}
func (this *Runner) printPassOrFail() {
	if this.testsPassed {
		this.printBanner("-- OK --")
	} else {
		this.printBanner("-- FAIL --")
	}
}

func getRepositoryRoot() string {
	return strings.TrimSpace(exec.RunOrFatal("", "git", "rev-parse", "--show-toplevel"))
}
