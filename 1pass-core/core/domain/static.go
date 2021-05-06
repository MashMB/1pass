// Definition of 1Pass application statics.
//
// @author TSS

package domain

const (
	AppName         string = "1pass"
	Archive         string = "1pass.tar.gz"
	BandFilePattern string = "band_*.js"
	CacheDir        string = "/tmp/1pass"
	Checksum        string = "checksum.md5"
	ConfigFile      string = "1pass.yml"
	GithubReleases  string = "https://api.github.com/repos/mashmb/1pass/releases"
	ProfileDir      string = "default"
	ProfileFile     string = "profile.js"
	Timeout         int64  = 2
	Version         string = "1.0.0"
)
