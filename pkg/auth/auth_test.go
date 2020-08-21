package auth

import (
	"kocli/pkg/config"
	"log"
	"testing"
)

func TestWithCredential(t *testing.T) {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = WithCredential("admin", "kubeoperator@admin123")
	if err != nil {
		log.Fatal(err)
	}
}
