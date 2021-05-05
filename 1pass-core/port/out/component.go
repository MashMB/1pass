// Definition of application components.
//
// @author TSS

package out

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type Updater interface {
	CheckForUpdate() (*domain.UpdateInfo, error)
}
