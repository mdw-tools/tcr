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

var Version = "dev"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		makefile  string
		gitignore string
		editor    string
		module    string
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of tcr-init [%s]:\n", Version)
		fmt.Fprintln(os.Stderr,
			"  The path in which to operate defaults to the current "+
				"directory or is the first non-flag argument supplied.")
		flag.PrintDefaults()
	}

	flag.StringVar(&makefile, "makefile", defaultMakefileContents, "default makefile content")
	flag.StringVar(&gitignore, "gitignore", ".idea/", ".gitignore file contents")
	flag.StringVar(&editor, "editor", "goland", "editor to invoke")
	flag.StringVar(&module, "module", "", "the go module name (defaults to base name of created directory)")

	flag.Parse()

	path := resolvePath()
	if module == "" {
		module = filepath.Base(path)
	}

	log.Printf("tcr-init [%s]\n", Version)
	createDirectory(path)
	createMakefile(path, makefile)
	createGitIgnore(path, gitignore)
	initializeGoModule(path, module)
	initializeGitRepository(path)
	startEditor(editor, path)
	log.Println("Finished.")
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
	log.Println("Creating:", directory)
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
	log.Println("Creating: Makefile")
	createFile(filepath.Join(path, "Makefile"), makefile)
}
func createFile(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
func createGitIgnore(path, gitignore string) {
	log.Println("Creating:", ".gitignore")
	createFile(filepath.Join(path, ".gitignore"), gitignore)
}
func initializeGoModule(path string, module string) {
	name := filepath.Base(path)
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, " ", "_")

	if module == "" {
		module = path
	}
	log.Println("Initializing go module:", module)
	fmt.Println(exec.RunFatal("go mod init " + module, exec.At(path)))
	fmt.Println(exec.RunFatal("go get github.com/smartystreets/gunit", exec.At(path)))
	fmt.Println(exec.RunFatal("go get github.com/smartystreets/assertions", exec.At(path)))

	createFile(filepath.Join(path, name+".go"), "package "+name)
	createFile(filepath.Join(path, name+"_test.go"), "package "+name)
}
func initializeGitRepository(path string) {
	log.Println("Initializing git repository...")
	fmt.Println(exec.RunFatal("git init", exec.At(path)))
	fmt.Println(exec.RunFatal("git add .", exec.At(path)))
	fmt.Println(exec.RunFatal("git commit -m 'Initial commit.'", exec.At(path)))
}
func startEditor(editor string, path string) {
	log.Println("Starting editor...")
	fmt.Println(exec.RunFatal(editor + " .", exec.At(path)))
}

const (
	defaultMakefileContents = `#!/usr/bin/make -f
test:
	go fmt ./...
	go test -cover -count=1 -timeout=1s -race -v ./...
`
)
