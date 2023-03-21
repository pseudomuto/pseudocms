package testutil

import (
	"fmt"
	"testing"

	"gotest.tools/v3/golden"
)

// AssertGolden ensures the actual string matches the contents of the golden file
// at ./testdata/${test.Name}.golden.
func AssertGolden(t *testing.T, actual string) {
	AssertGoldenExt(t, actual, "golden")
}

// AssertGoldenExt functions identically to AssetGolden, only it adds .<ext> to the
// resulting file name.
func AssertGoldenExt(t *testing.T, actual, ext string) {
	golden.Assert(t, actual, fmt.Sprintf("%s.%s", t.Name(), ext))
}
