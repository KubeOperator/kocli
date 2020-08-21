package main

import (
	"kocli/cmd"
	"kocli/pkg/config"
)

func main() {
	_ = config.LoadConfig()
	_ = cmd.Execute()
}
