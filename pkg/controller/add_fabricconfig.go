package controller

import (
	"github.com/stevenchiu30801/onos-bandwidth-operator/pkg/controller/fabricconfig"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, fabricconfig.Add)
}
