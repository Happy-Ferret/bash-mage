package mage

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

// FLU is a namespace.
type FLU mg.Namespace

// Deploy3 deploys stuff.
func (FLU) Deploy3() {
	fmt.Println("deploy3 from FLU")
}
