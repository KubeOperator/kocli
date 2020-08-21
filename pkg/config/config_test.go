package config

import (
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}
