package testutil

import "testing"

// TestPopFakeNhc dummy test for covering testutil pkg
func TestPopFakeNhc(t *testing.T) {
	t.Run("PopFakeNhc", func(t *testing.T) {
		PopFakeNhc()

		InitStubNHC()

	})
}
