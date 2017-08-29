package testutil

import "testing"

func TestPopFakeNhc(t *testing.T) {
	t.Run("PopFakeNhc", func(t *testing.T) {
		PopFakeNhc()

		InitStubNHC()

	})
}
