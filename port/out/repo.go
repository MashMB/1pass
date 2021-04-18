// Definition of repositories.
//
// @author TSS

package out

const (
	ProfileDir string = "default"
)

type ProfileRepo interface {
	GetSalt() string
}
