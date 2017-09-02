package testutil

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("starting testutil test")
}

// TestPopFakeNhc dummy test for covering testutil pkg
func TestPopFakeNhc(t *testing.T) {
	t.Run("PopFakeNhc", func(t *testing.T) {
		PopFakeNhc()

		InitStubNHC()

	})
}
