package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdwhatcott/tcr/exec"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		makefile  string
		gitignore string
		editor    string
	)

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			"  The path in which to operate defaults to the current "+
				"directory or is the first non-flag argument supplied.")
		flag.PrintDefaults()
	}
	flag.StringVar(&makefile, "makefile", defaultMakefileContents, "default makefile content")
	flag.StringVar(&gitignore, "gitignore", ".idea/", ".gitignore file contents")
	flag.StringVar(&editor, "editor", "goland", "editor to invoke")

	flag.Parse()

	path := resolvePath()
	createDirectory(path)
	createMakefile(path, makefile)
	createGitIgnore(path, gitignore)
	initializeGoModule(path)
	initializeGitRepository(path)
	startEditor(editor, path)
}
func resolvePath() string {
	path := flag.Arg(0)
	if path == "" {
		path, _ = os.Getwd()
	}
	path, _ = filepath.Abs(path)
	return path
}
func createDirectory(directory string) {
	exists, err := pathExists(directory)
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		return
	}
	err = os.MkdirAll(directory, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
func createMakefile(path, makefile string) {
	createFile(filepath.Join(path, "Makefile"), makefile)
}
func createFile(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
func createGitIgnore(path, gitignore string) {
	createFile(filepath.Join(path, ".gitignore"), gitignore)
}
func initializeGoModule(path string) {
	name := filepath.Base(path)
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, " ", "_")

	fmt.Println(exec.RunFatal("go", exec.Args("mod", "init", name), exec.At(path)))
	fmt.Println(exec.RunFatal("go", exec.Args("get", "github.com/smartystreets/gunit"), exec.At(path)))
	fmt.Println(exec.RunFatal("go", exec.Args("get", "github.com/smartystreets/assertions"), exec.At(path)))

	createFile(filepath.Join(path, name+".go"), "package "+name)
	createFile(filepath.Join(path, name+"_test.go"), "package "+name)
}
func initializeGitRepository(path string) {
	fmt.Println(exec.RunFatal("git", exec.Args("init"), exec.At(path)))
	fmt.Println(exec.RunFatal("git", exec.Args("add", "."), exec.At(path)))
	fmt.Println(exec.RunFatal("git", exec.Args("commit", "-m", "Initial commit."), exec.At(path)))
}
func startEditor(editor string, path string) {
	fmt.Println(exec.RunFatal(editor, exec.Args("."), exec.At(path)))
}

const (
	defaultMakefileContents = `#!/usr/bin/make -f
test:
	go fmt ./...
	go test -cover -count=1 -timeout=1s -race -v ./...
`
)
