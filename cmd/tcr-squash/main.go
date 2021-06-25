package main

import (
	"fmt"

	"github.com/mdwhatcott/tcr/exec"
	"github.com/mdwhatcott/tcr/exec/git"
)

var Version = "dev"

func main() {
	count := git.TCRCommitCount()
	fmt.Printf("tcr-squash [%s] resetting [%d] commits into a single, staged change set...\n", Version, count)
	fmt.Println(exec.RunFatal("git reset --soft " + fmt.Sprintf("HEAD~%d", count)))
	fmt.Println("Ready for commit!")
}
