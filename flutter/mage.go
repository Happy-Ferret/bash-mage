package mage

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

// FLU is a namespace.
type FLU mg.Namespace

// Deploy4 deploys stuff.
func (FLU) Deploy4() {
	fmt.Println("deploy4 from FLU")
}
