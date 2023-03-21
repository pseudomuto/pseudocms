package pseudocms

import "fmt"

// Version details
const (
	MajorVersion      = 0
	MinorVersion      = 1
	PatchVersion      = 0
	PreReleaseVersion = "alpha"
)

// Version returns the current version of pseudocms.
func Version() string {
	v := fmt.Sprintf("%d.%d.%d", MajorVersion, MinorVersion, PatchVersion)
	if PreReleaseVersion != "" {
		v = fmt.Sprintf("%s+%s", v, PreReleaseVersion)
	}

	return v
}
