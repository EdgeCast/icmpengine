// Copyright 2021 Edgecast Inc

// +build !race

// Package israce reports if the Go race detector is enabled.
package norace

// Enabled reports if the race detector is enabled.
const Enabled = false
