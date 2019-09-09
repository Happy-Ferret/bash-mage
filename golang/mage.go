package mage

import (
	"fmt"

	"os"

	"github.com/chirino/hawtgo/sh"
	"github.com/magefile/mage/mg"
)

// GO is a namespace.
type GO mg.Namespace

// Deploy2 deploys stuff.
func (GO) Deploy2() {
	fmt.Println("deploy2 from go")

	// call flu ?

}

var mysh = sh.New().
	CommandLog(os.Stdout).
	CommandLogPrefix("> ").
	Env(map[string]string{
		"CGO_ENABLED": "0",
		"GOOS":        "linux",
		"GOARCH":      "amd64",
	}).
	Dir("./target")

// BuildExecutable builds
func (GO) BuildExecutable() {
	mysh.Line(`go build -o "my app" github.com/chirino/cmd/myapp`).MustExec()
	mysh.Line(`go build -o otherapp github.com/chirino/cmd/otherapp`).MustExec()
}
