// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"github.com/hyperblaze-labs/hypersdk-spam/codec"
	"github.com/hyperblaze-labs/hypersdk-spam/examples/tokenvm/cmd/token-feed/manager"
)

type Manager interface {
	GetFeedInfo(context.Context) (codec.Address, uint64, error)
	GetFeed(context.Context) ([]*manager.FeedObject, error)
}
