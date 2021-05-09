// Definition of application components.
//
// @author TSS

package out

import (
	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type Updater interface {
	CheckForUpdate(timeout int64) (*domain.UpdateInfo, error)

	CheckTimestamp(dirPath string)

	DownloadFile(destination, url string, timeout int64) error

	ExtractArchive(src, dst string) error

	ReplaceBinary(src string) error

	ShouldCheck(period int, dirPath string) bool

	ValidateChecksum(binary, file string) error
}
