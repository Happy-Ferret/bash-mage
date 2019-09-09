// +build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
	// mage:import test
	_ "github.com/LUSHDigital/core-mage/targets/tests"
	// mage:import goo
	_ "github.com/joeblew99/bash-mage/golang"
	// mage:import flu
	_ "github.com/joeblew99/bash-mage/flutter"
)

func init() {
	targets.ProjectName = "example"
	targets.ProjectType = "test"
	targets.DockerComposeDevDependencies = []string{"redis"}
	targets.DockerComposeTestDependencies = []string{"redis"}
}
