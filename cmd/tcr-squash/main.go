package main

import (
	"fmt"

	"github.com/mdwhatcott/tcr/exec"
)

func main() {
	count := exec.GetTCRCommitCount()
	fmt.Printf("Squashing [%d] commits into single, staged change set and opening smerge...\n", count)
	fmt.Println(exec.RunFatal("git", exec.Args("reset", "--soft", fmt.Sprintf("HEAD~%d", count))))
	output, err := exec.Run("smerge", exec.Args("."))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Output:")
	fmt.Println(output)
}
