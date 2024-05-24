package gotest

import (
	"fmt"
	"strings"
)

func Format(output string) string {
	lines := strings.Split(output, "\n")
	bulkTestOutputLines := make(map[int]BulkGoTestLine)
	for l, line := range lines {
		if isBulkGoTestLine(line) {
			bulkTestOutputLines[l] = ParseBulkGoTestLine(line)
		}
	}
	var (
		maxResult   int
		maxPackage  int
		maxDuration int
		maxCoverage int
	)
	for _, line := range bulkTestOutputLines {
		if len(line.Result) > maxResult {
			maxResult = len(line.Result)
		}
		if len(line.PackageName) > maxPackage {
			maxPackage = len(line.PackageName)
		}
		if len(line.Duration) > maxDuration {
			maxDuration = len(line.Duration)
		}
		if len(line.Coverage) > maxCoverage {
			maxCoverage = len(line.Coverage)
		}
	}
	for l := range lines {
		bulk, ok := bulkTestOutputLines[l]
		if ok {
			lines[l] = bulk.Format(maxResult, maxPackage, maxDuration, 6)
		}
	}
	return strings.Join(lines, "\n")
}

func isBulkGoTestLine(line string) bool {
	return strings.HasPrefix(line, "ok  \t") ||
		strings.HasPrefix(line, "?   \t") ||
		strings.HasPrefix(line, "FAIL\t") ||
		(strings.HasPrefix(line, "\t") && strings.Contains(line, "coverage:"))
}

func ParseBulkGoTestLine(line string) BulkGoTestLine {
	if strings.HasPrefix(line, "\t") {
		line = "- " + strings.Replace(line, "of statements", "", 1) + " hi [missing test files]"
	}
	fields := strings.Fields(line)
	result := BulkGoTestLine{
		Original:    line,
		Result:      fields[0],
		PackageName: fields[1],
	}
	if strings.Contains(line, "[no test files]") {
		result.Duration = "[no test files]"
	} else if strings.Contains(line, "[missing test files]") {
		result.Duration = "[missing test files]"
	} else {
		result.Duration = fields[2]
	}
	if strings.Contains(line, "\tcoverage: ") && strings.Contains(line, "% of statements") {
		result.Coverage = fields[4]
		for len(result.Coverage) < 6 {
			result.Coverage = " " + result.Coverage
		}
	} else if strings.Contains(line, "[no tests to run]") {
		result.Coverage = "[nope]"
	}
	return result
}

type BulkGoTestLine struct {
	Original    string
	Result      string
	PackageName string
	Duration    string
	Coverage    string
}

func (this BulkGoTestLine) Format(result, packageName, duration, coverage int) string {
	format := fmt.Sprintf("%%-%ds %%-%ds %%-%ds %%%ds", result, packageName, coverage, duration)
	line := fmt.Sprintf(format, this.Result, this.PackageName, this.Coverage, this.Duration)
	return strings.TrimSpace(line)
}
