package embedded

import (
	"testing"

	"github.com/styrainc/regal/internal/test/assert"
)

func TestEmbeddedEOPAHasAtLeastExpectedNumberOfCapabilities(t *testing.T) {
	t.Parallel()

	versions := assert.Do(LoadCapabilitiesVersions("eopa")).Or(t, "error loading EOPA capabilities versions")

	assert.AtLeast(t, 47, len(versions), "EOPA capabilities in the embedded database")

	for _, v := range versions {
		caps := assert.Do(LoadCapabilitiesVersion("eopa", v)).Or(t, "error loading EOPA capabilities version", v)

		assert.AtLeast(t, 1, len(caps.Builtins), "EOPA capabilities version", v)
	}
}
