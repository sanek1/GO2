package main

import (
	config "github.com/rowdyroad/go-yaml-config"
)

func main() {
	var cfg struct{}
	config.LoadConfig(&cfg, "test.yaml", nil)
}
