package main

import (
	"fmt"

	"github.com/mdwhatcott/tcr/exec"
)

var Version = "dev"

func main() {
	count := exec.GetTCRCommitCount()
	fmt.Printf("tcr-squash [%s] resetting [%d] commits into a single, staged change set...\n", Version, count)
	fmt.Println(exec.RunFatal("git", exec.Args("reset", "--soft", fmt.Sprintf("HEAD~%d", count))))
	fmt.Println("Ready for commit!")
}
