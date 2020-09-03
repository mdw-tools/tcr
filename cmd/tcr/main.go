// TCR: test && commit || revert
// codified by Kent Beck
// https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mdwhatcott/tcr/exec"
)

var Version = "dev"

func main() {
	flag.Usage = func() {

	}
	flag.Parse()

	runner := newRunner(Version)
	runner.TCR()
	fmt.Print(runner.FinalReport())
}

func newRunner(version string) *Runner {
	builder := new(strings.Builder)
	fmt.Fprintf(builder, "tcr [%s]\n", version)
	return &Runner{finalReport: builder}
}

type Runner struct {
	started time.Time
	stopped time.Time

	testReport  string
	gitReport   string
	finalReport *strings.Builder

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

	output, err := exec.Run("make", exec.At(getRepositoryRoot()), exec.Out(os.Stdout))
	this.testReport = strings.TrimSpace(output)
	this.testsPassed = err == nil
	return this.testsPassed

}
func (this *Runner) start() { this.started = time.Now() }
func (this *Runner) stop()  { this.stopped = time.Now() }
func (this *Runner) elapsed() time.Duration {
	return this.stopped.Sub(this.started).Round(time.Millisecond)
}

func (this *Runner) Commit() bool {
	_, _ = exec.Run("git", exec.Args("add", "."))
	output, _ := exec.Run("git", exec.Args("commit", "-m", "tcr"))
	this.gitReport = strings.TrimSpace(output)
	return true
}

func (this *Runner) Revert() bool {
	this.gitReport += exec.RunFatal("git", exec.Args("clean", "-df"), exec.Out(os.Stdout))
	this.gitReport += exec.RunFatal("git", exec.Args("reset", "--hard"), exec.Out(os.Stdout))
	this.gitReport += exec.RunFatal("pbcopy", exec.Args("less is more"), exec.Out(os.Stdout))
	return true
}

func (this *Runner) FinalReport() string {
	this.finalReport = new(strings.Builder)
	if !this.testsPassed {
		this.printSummary()
	}
	this.printReport(this.gitReport)
	if !this.testsPassed {
		this.printBanner("Test failures repeated for convenience:")
		fmt.Fprintln(this.finalReport)
		this.printReport(filterPassingPackages(this.testReport))
	}
	if this.testsPassed {
		this.printSummary()
	}
	return strings.TrimSpace(this.finalReport.String())
}

func filterPassingPackages(report string) string {
	var builder strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(report))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ok  \t") {
			continue
		}
		if strings.HasPrefix(line, "?   \t") {
			continue
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return builder.String()
}

func (this *Runner) printReport(report string) {
	this.finalReport.WriteString(report)
}
func (this *Runner) printSummary() {
	fmt.Fprintln(this.finalReport)
	fmt.Fprintln(this.finalReport)
	fmt.Fprintf(this.finalReport,
		"%s in [%v] with [%d] tcr commit(s) at %s",
		this.passOrFail(),
		this.elapsed(),
		this.commitCount,
		workingDirectory(),
	)
}
func (this *Runner) printBanner(banner string) {
	this.finalReport.WriteString("\n\n")
	this.finalReport.WriteString(banner)
	this.finalReport.WriteString("\n\n")
}
func workingDirectory() string {
	dir, _ := os.Getwd()
	return dir
}
func (this *Runner) passOrFail() string {
	if this.testsPassed {
		return "OK"
	} else {
		return "FAIL"
	}
}

func getRepositoryRoot() string {
	return strings.TrimSpace(exec.RunFatal("git", exec.Args("rev-parse", "--show-toplevel")))
}
