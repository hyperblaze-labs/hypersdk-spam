// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package auth

import (
	"github.com/hyperblaze-labs/hypersdk-spam/examples/morpheusvm/consts"
	"github.com/hyperblaze-labs/hypersdk-spam/vm"
)

func Engines() map[uint8]vm.AuthEngine {
	return map[uint8]vm.AuthEngine{
		// Only ed25519 batch verification is supported
		consts.ED25519ID: &ED25519AuthEngine{},
	}
}
