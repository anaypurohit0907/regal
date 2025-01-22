//go:build integration
// +build integration

package capabilities

import (
	"context"
	"testing"

	"github.com/styrainc/regal/internal/test/assert"
)

func TestCanLookupCapabilitiesFromURL(t *testing.T) {
	t.Parallel()

	cURL := "https://raw.githubusercontent.com/open-policy-agent/opa/main/capabilities/v0.55.0.json"
	caps := assert.
		Do(Lookup(context.TODO(), cURL)).
		Or(t, "unexpected error from Lookup")

	assert.Equal(t, 193, len(caps.Builtins), "OPA v0.55.0 capabilities should have 193 builtins")

}
