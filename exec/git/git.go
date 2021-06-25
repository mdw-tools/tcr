package git

import (
	"strings"

	"github.com/mdwhatcott/tcr/exec"
)

func TCRCommitCount() (count int) {
	rawLog := exec.RunFatal("git log --oneline")
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

func RepositoryRoot() string {
	return strings.TrimSpace(exec.RunFatal("git rev-parse --show-toplevel"))
}
