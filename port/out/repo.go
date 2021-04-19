// Definition of repositories.
//
// @author TSS

package out

const (
	ProfileDir string = "default"
)

type ProfileRepo interface {
	GetIterations() int

	GetMasterKey() string

	GetOverviewKey() string

	GetSalt() string
}