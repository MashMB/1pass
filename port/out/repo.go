// Definition of repositories.
//
// @author TSS

package out

import (
	"github.com/mashmb/1pass/core/domain"
)

const (
	BandFilePattern string = "band_*.js"
	ProfileDir      string = "default"
)

type ItemRepo interface {
	FindByCategoryAndTrashed(category *domain.ItemCategory, trashed bool) []*domain.RawItem
}

type ProfileRepo interface {
	GetIterations() int

	GetMasterKey() string

	GetOverviewKey() string

	GetSalt() string
}
