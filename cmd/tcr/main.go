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
	"github.com/mdwhatcott/tcr/exec/git"
)

var Version = "dev"

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "Usage of tcr:")
		flag.CommandLine.PrintDefaults()
	}
	wd, _ := os.Getwd()
	command := flag.String("command", orDefault(os.Getenv("TCR_EXECUTABLE"), "make"), "The 'test' command.")
	working := flag.String("working", orDefault(git.RepositoryRoot(), wd), "The working directory.")
	dryRun := flag.Bool("dry-run", false, "When set, test but don't commit or revert.")
	flag.Parse()

	runner := newRunner(Version, *command, *working, *dryRun)
	runner.TCR()
	fmt.Print(runner.FinalReport())
}

func orDefault(value, fallback string) string {
	if value != "" {
		return value
	}
	return fallback
}

func newRunner(version, program, working string, dryRun bool) *Runner {
	builder := new(strings.Builder)
	_, _ = fmt.Fprintf(builder, "tcr [%s]\n", version)
	return &Runner{
		program:     program,
		working:     working,
		dryRun:      dryRun,
		finalReport: builder,
	}
}

type Runner struct {
	program string
	working string
	dryRun  bool

	started time.Time
	stopped time.Time

	testReport  string
	gitReport   string
	finalReport *strings.Builder

	commitCount int
	testsPassed bool
}

func (this *Runner) TCR() {
	_ = this.Test() &&
		this.Commit() ||
		this.Revert()

	this.finalizeReport()
	this.resetStopWatch()
}

func (this *Runner) Test() bool {
	this.start()
	defer this.stop()

	output, err := exec.Run(this.program, exec.At(this.working), exec.Out(os.Stdout))
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
	if this.dryRun {
		this.gitReport += "- dry-run enabled, skipping git commit"
		return true
	}
	_, _ = exec.Run("git add .")
	output, _ := exec.Run("git commit -m tcr")
	this.gitReport = strings.TrimSpace(output)
	this.commitCount = git.TCRCommitCount()
	return true
}

func (this *Runner) Revert() bool {
	if this.dryRun {
		this.gitReport += "- dry-run enabled, skipping git revert"
		return true
	}
	this.gitReport += exec.RunFatal("git clean -df", exec.Out(os.Stdout))
	this.gitReport += exec.RunFatal("git reset --hard", exec.Out(os.Stdout))
	return true
}

func (this *Runner) FinalReport() string {
	return strings.TrimSpace(this.finalReport.String()) + "\n"
}

func (this *Runner) finalizeReport() {
	this.finalReport = new(strings.Builder)
	if !this.testsPassed {
		this.printSummary()
	}

	this.printBanner("---")
	this.printReport(strings.TrimSpace(this.gitReport))

	this.printBanner("Test output reformatted for convenience:")
	this.printReport(this.testReport)

	if !this.testsPassed {
		this.printBanner("---")
		this.printBanner("Test failures repeated for convenience:")
		this.printReport(filterPassingPackages(this.testReport))
	}
	if this.testsPassed {
		this.printSummary()
	}
}

func filterPassingPackages(report string) string {
	var builder strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(report))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ok ") {
			continue
		}
		if strings.HasPrefix(line, "?  ") {
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
	this.printBanner("---")
	_, _ = fmt.Fprintf(this.finalReport,
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

func (this *Runner) resetStopWatch() {
	_, _ = http.Get("http://localhost:7890/stopwatch/reset")
}
