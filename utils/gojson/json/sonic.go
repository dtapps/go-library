// Copyright 2022 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build sonic && avx && (linux || windows || darwin) && amd64
// +build sonic
// +build avx
// +build linux windows darwin
// +build amd64

package json

import "github.com/bytedance/sonic"

var (
	json = sonic.ConfigStd
	// Marshal is exported by gojson/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gojson/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gojson/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gojson/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gojson/json package.
	NewEncoder = json.NewEncoder
)
