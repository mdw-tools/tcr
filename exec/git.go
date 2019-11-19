package exec

import "strings"

func GetTCRCommitCount() (count int) {
	rawLog := RunOrFatal("", "git", "log", "--oneline")
	logLines := strings.Split(rawLog, "\n")
	for _, line := range logLines {
		if strings.HasSuffix(line, " tcr") {
			count++
		} else {
			break
		}
	}
	return count
}
