# newgoproj
this is a simple tool for accessing and creating go projects

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

modify the gopath, website, user var in lazy.go to your workspace  
use the command, go build lazy.go, to build the executable  

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
the github link is assumed to be  
https://github.com/user/projectname.git  