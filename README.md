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
