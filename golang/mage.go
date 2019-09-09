package mage

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

// GO is a namespace.
type GO mg.Namespace

// Deploy2 deploys stuff.
func (GO) Deploy2() {
	fmt.Println("deploy2 from go")
}
