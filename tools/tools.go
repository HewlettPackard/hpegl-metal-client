//go:build tools
// +build tools

// (C) Copyright 2022 Hewlett Packard Enterprise Development LP

package tools

import (
	// required for interfaces & mocks generation
	_ "github.com/golang/mock/mockgen"
	_ "github.com/vburenin/ifacemaker"
)
