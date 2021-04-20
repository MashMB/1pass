// Definition of repositories.
//
// @author TSS

package out

const (
	BandFilePattern string = "band_*.js"
	ProfileDir      string = "default"
)

type ItemRepo interface {
}

type ProfileRepo interface {
	GetIterations() int

	GetMasterKey() string

	GetOverviewKey() string

	GetSalt() string
}
