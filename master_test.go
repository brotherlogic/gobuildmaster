package main

import (
	"testing"
)

func TestLoadMainConfig(t *testing.T) {
	_, err := loadConfig()
	if err != nil {
		t.Errorf("Config load failed: %v", err)
	}
}
