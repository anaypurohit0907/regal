package capabilities

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/styrainc/regal/internal/test/assert"
)

func TestCanLookupCapabilitiesFromFile(t *testing.T) {
	t.Parallel()

	path := assert.Do(filepath.Abs("./testdata/capabilities.json")).Or(t, "no absolute path for capabilities file")
	caps := assert.Do(Lookup(context.TODO(), "file://"+path)).Or(t, "unexpected error from Lookup")

	assert.Equal(t, 1, len(caps.Builtins), "expected capabilities to have exactly 1 builtin")
	assert.Equal(t, "unittest123", caps.Builtins[0].Name, "builtin name is incorrect")
}

func TestCanLookupCapabilitiesFromEmbedded(t *testing.T) {
	t.Parallel()

	caps := assert.Do(Lookup(context.TODO(), "regal:///capabilities/opa/v0.55.0")).Or(t, "unexpected error from Lookup")

	assert.Equal(t, 193, len(caps.Builtins), "OPA v0.55.0 capabilities should have 193 builtins")
}

func TestSemverSort(t *testing.T) {
	t.Parallel()

	cases := []struct {
		note   string
		input  []string
		expect []string
	}{
		{
			note:   "should be able to correctly sort semver only",
			input:  []string{"1.2.3", "1.2.4", "1.0.1"},
			expect: []string{"1.2.4", "1.2.3", "1.0.1"},
		},
		{
			note:   "should be able to correctly sort non-semver only",
			input:  []string{"a", "b", "c"},
			expect: []string{"c", "b", "a"},
		},
		{
			note:   "should be able to correctly sort mixed semver and non-semver",
			input:  []string{"a", "b", "c", "4.0.7", "1.0.1", "2.1.1", "2.3.4"},
			expect: []string{"4.0.7", "2.3.4", "2.1.1", "1.0.1", "c", "b", "a"},
		},
	}

	for _, c := range cases {
		// Sorts the input in-place, which works since we won't re-visit the same test twice.
		semverSort(c.input)

		for j, x := range c.expect {
			if x != c.input[j] {
				t.Errorf("index=%d actual='%s' expected='%s'", j, c.input[j], x)
			}
		}
	}
}
