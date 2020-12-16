package gotest

import (
	"fmt"
	"strings"
)

func Format(output string) string {
	lines := strings.Split(output, "\n")
	bulkTestOutputLines := make(map[int]BulkGoTestLine)
	for l, line := range lines {
		if strings.HasPrefix(line, "ok  \t") || strings.HasPrefix(line, "?   \t") || strings.HasPrefix(line, "FAIL\t") {
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
			lines[l] = bulk.Format(maxResult, maxPackage, maxDuration, maxCoverage)
		}
	}
	return strings.Join(lines, "\n")
}

func ParseBulkGoTestLine(line string) BulkGoTestLine {
	fields := strings.Fields(line)
	result := BulkGoTestLine{
		Original:    line,
		Result:      fields[0],
		PackageName: fields[1],
	}
	if strings.Contains(line, "[no test files]") {
		result.Duration = "[no tests files]"
	} else {
		result.Duration = fields[2]
	}
	if strings.Contains(line, "\tcoverage: ") && strings.Contains(line, "% of statements") {
		result.Coverage = fields[4]
	} else if strings.Contains(line, "[no tests to run]") {
		result.Coverage = "[no tests to run]"
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
