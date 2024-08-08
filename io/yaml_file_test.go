package io

import (
	"fmt"
	"testing"
)

func TestChargeConfigFile(t *testing.T) {
	path := "/Users/luismanueltorrestrevino/workspaces/gocrity/resources/alacritty.yaml"
	config, err := ChargeConfigFile(path)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(config)
}
