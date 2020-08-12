# newgoproj
this is a simple tool for accessing and creating go projects

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:
```bash
go get -u github.com/cgliu-create/newgoproj
```

## Usage
use the go tool
```bash
go run lazy.go
```
use the executable
```bash
./lazy
```
running this command will either  
A) create a project with go.mod and main.go  
B) open an existing project  

modify the variables in lazy.go to your workspace  
```go
var (
	gopath  = "gopath/location"
	website = "github.com"
	user    = "cgliu-create"
)
```
use the following command to build the executable  
```bash
go build lazy.go
```
## Flag Variables
```bash
./lazy
```
starts Neovim for a project called moo
```bash
./lazy -proj="projectname" 
```
starts Neovim for the project
```bash
./lazy -proj="projectname" -git=true
```
starts Neovim for the project  
initalizes a git repository  
-the github link is assumed to be-  
https://github.com/user/projectname.git  
