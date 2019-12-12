package exec

import "strings"

func GetTCRCommitCount() (count int) {
	rawLog := RunFatal("git", Args("log", "--oneline"))
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
