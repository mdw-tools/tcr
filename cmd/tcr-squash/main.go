package main

import (
	"fmt"

	"github.com/mdwhatcott/tcr/exec"
)

func main() {
	count := exec.GetTCRCommitCount()
	fmt.Printf("Resetting [%d] commits into a single, staged change set...\n", count)
	fmt.Println(exec.RunFatal("git", exec.Args("reset", "--soft", fmt.Sprintf("HEAD~%d", count))))
	fmt.Println("Ready for commit!")
}
