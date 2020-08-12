package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var proj = flag.String("proj", "moo", "project name")
var (
	gopath  = "GoChris/src"
	website = "github.com"
	user    = "cgliu-create"
)

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
		startFile :=
`package main

import (
  "fmt"
)

func main() {
	fmt.Println("hello world")
}`
		ioutil.WriteFile("main.go", []byte(startFile), 0644)
	}
	fmt.Println("starting nvim")
	cmd := exec.Command("nvim", ".")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
