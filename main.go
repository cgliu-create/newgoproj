package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)
var proj = flag.String("proj", "moo", "project name")
var git = flag.Bool("git", false, "init git repository")
var (
	gopath  = "GoChris/src"
	website = "github.com"
	user    = "cgliu-create"
)
var startFile =
`package main

import (
  "fmt"
)

func main() {
	fmt.Println("hello world")
}`
func main() {
	flag.Parse()
	fmt.Printf("working on project %v", *proj)
	fmt.Println()
	os.Chdir(gopath)
	os.Chdir(website)
	os.Chdir(user)
	_, err := os.Stat(*proj)
	if os.IsNotExist(err) {
		fmt.Println("creating new project")
		os.Mkdir(*proj, 0755)
	}
	fmt.Println("going to project")
	os.Chdir(*proj)
	if os.IsNotExist(err) {
		fmt.Println("initializing project")
		cmd := exec.Command("go", "mod", "init")
		cmd.Run()
		ioutil.WriteFile("main.go", []byte(startFile), 0644)
	}
  if *git {
		fmt.Println("initializing git")
    title := fmt.Sprintf("# %v", *proj)
		ioutil.WriteFile("README.md", []byte(title), 0644)
    cmdi := exec.Command("git", "init")
    cmdi.Run()
    cmdar := exec.Command("git", "add", "README.md")
    cmdar.Run()
    cmdam := exec.Command("git", "add", "go.mod")
    cmdam.Run()
    cmdaf := exec.Command("git", "add", "main.go")
    cmdaf.Run()
    cmdc := exec.Command("git", "commit", "-m", "\"first commit\"")
    cmdc.Run()
    gitlink := fmt.Sprintf("https://github.com/%v/%v.git", user, *proj)
    fmt.Println(gitlink)
    cmdo := exec.Command("git", "remote", "add", "origin", gitlink)
    cmdo.Run()
    cmdp := exec.Command("git", "push", "-u", "origin", "master")
    cmdp.Run()
  }
	fmt.Println("starting nvim")
	cmd := exec.Command("nvim", ".")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
