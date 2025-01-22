package assert

import (
	"cmp"
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, expected, got T, message ...string) {
	t.Helper()

	if expected != got {
		if len(message) > 0 {
			t.Errorf("%s: expected %v, got %v", strings.Join(message, " "), expected, got)
		} else {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
}

func AtLeast[T cmp.Ordered](t *testing.T, expectedAtLeast, got T, message ...string) {
	t.Helper()

	if got < expectedAtLeast {
		if len(message) > 0 {
			t.Errorf("%s: expected at least %v, got %v", strings.Join(message, " "), expectedAtLeast, got)
		} else {
			t.Errorf("expected at least %v, got %v", expectedAtLeast, got)
		}
	}
}

type result[T any] struct {
	value T
	err   error
}

func (r result[T]) Or(t *testing.T, message ...string) T {
	t.Helper()

	if r.err != nil {
		t.Fatalf("%s: %v", strings.Join(message, " "), r.err)
	}

	return r.value
}

func Do[T any](ret T, err error) result[T] {
	return result[T]{value: ret, err: err}
}
