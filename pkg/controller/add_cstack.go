package controller

import (
	"github.com/cstack/operator/pkg/controller/cstack"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cstack.Add)
}
