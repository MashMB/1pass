// Implementation of cobra CLI control.
//
// @author TSS

package cobra

import (
	"github.com/mashmb/1pass/core/facade"
)

type cobraCliControl struct {
	vaultFacade facade.VaultFacade
}

func NewCobraCliControl(vaultFacade facade.VaultFacade) *cobraCliControl {
	return &cobraCliControl{
		vaultFacade: vaultFacade,
	}
}
