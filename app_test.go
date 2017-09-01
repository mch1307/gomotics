package main

import (
	"path/filepath"
	"testing"
)

func TestSub(t *testing.T) {
	Sub(filepath.Join(".", "config", "test.toml"))
	// test listening on port, sleep close
}
