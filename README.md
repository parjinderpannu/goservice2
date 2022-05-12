# goservice2
test1
test2
https://github.com/ardanlabs/service

[goservice2] export GOPATH="/Users/ppannu/code/personal/go-tutorial/goservice2/go" 
[goservice2] env | grep gopath 
[goservice2] env | grep -i gopath
GOPATH=/Users/ppannu/code/personal/go-tutorial/goservice2/go

Explaining go module
go mod tidy #found the pkg and selected the version
- look at GOPROXY variable
- check with proxy server and request for module of code for that github.com/pkg/errors. Also latest version.
- google proxy is check and go to github for latest tag version and pull it and bundle it to mod and zip
- first time takes longer
- go.mod has that information as well.
- go.sum : hashcode 1. content of zip file 2. mod file
- go mod build : going to google and get pkg with TIDY
- google is keeping request for 30 days
- change GOPROXY direct
- private mod mirros : Athens (open source)
- private mod mirros : jfrog (propietory)
- change mod GOPROXY and use , direct because if you 404,401 then it could go to github
- private mod mirrors config: either go to github or some stuff go to google (privacy concern)
- GOPRIVATE variable 
- GONOPROXY means don't go to mod mirror but go direct

go env 
GOPROXY

## gitmvsAlgorithm
- dependency A needs indirect dep.. D
- D 1.9
- go.mod of A shares that D 1.2
- we need to look for go sum mod using cmd: go list mod
- scenario : A1.4 -> D1.2
- scenario : B1.6 -> D1.4
- we have to choose one version of D therefore go choose D1.4
- Remove scenario : B1.6 -> D1.4
- But now we are left with D1.4 not D1.2

- what if you want to use D1.9 cmd:: go get to update to latest 
- latest of direct
- scenario : A1.6 -> D1.7
- NOTE: we only have to test our own test not of dependencies
- author : likes to get latest of everything

## checksumDB

- GOPRIVATE -> GONOPROXY GONOSUMDB
- checksumdb : something like writeonly and tell write like blockchain
- mod mirrors tells checksumdb and checksumdb maintain hash code

## majorVersion
- https://github.com/dimfeld/httptreemux
- go mod tidy
- go build main.go
- latest of treemux is 5.4.0 but tooling choose v5.0.1+incompatible
-- there is no go.mod file in 5.0.1
--- 

## vendoring
