package main

import (
	"fmt"

	"github.com/mdwhatcott/tcr/exec"
)

func main() {
	count := exec.GetTCRCommitCount()
	fmt.Printf("Squashing [%d] commits into single, staged change set and opening smerge...\n", count)
	fmt.Println(exec.RunOrFatal("", "git", "reset", "--soft", fmt.Sprintf("HEAD~%d", count)))
	_, _ = exec.Run("", "smerge", ".")
}
