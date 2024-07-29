// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package controller

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils/logging"

	"github.com/hyperblaze-labs/hypersdk-spam/codec"
	"github.com/hyperblaze-labs/hypersdk-spam/examples/tokenvm/genesis"
	"github.com/hyperblaze-labs/hypersdk-spam/examples/tokenvm/orderbook"
	"github.com/hyperblaze-labs/hypersdk-spam/examples/tokenvm/storage"
	"github.com/hyperblaze-labs/hypersdk-spam/fees"
)

func (c *Controller) Genesis() *genesis.Genesis {
	return c.genesis
}

func (c *Controller) Logger() logging.Logger {
	return c.inner.Logger()
}

func (c *Controller) Tracer() trace.Tracer {
	return c.inner.Tracer()
}

func (c *Controller) GetTransaction(
	ctx context.Context,
	txID ids.ID,
) (bool, int64, bool, fees.Dimensions, uint64, error) {
	return storage.GetTransaction(ctx, c.metaDB, txID)
}

func (c *Controller) GetAssetFromState(
	ctx context.Context,
	asset ids.ID,
) (bool, []byte, uint8, []byte, uint64, codec.Address, error) {
	return storage.GetAssetFromState(ctx, c.inner.ReadState, asset)
}

func (c *Controller) GetBalanceFromState(
	ctx context.Context,
	addr codec.Address,
	asset ids.ID,
) (uint64, error) {
	return storage.GetBalanceFromState(ctx, c.inner.ReadState, addr, asset)
}

func (c *Controller) Orders(pair string, limit int) []*orderbook.Order {
	return c.orderBook.Orders(pair, limit)
}

func (c *Controller) GetOrderFromState(
	ctx context.Context,
	orderID ids.ID,
) (
	bool, // exists
	ids.ID, // in
	uint64, // inTick
	ids.ID, // out
	uint64, // outTick
	uint64, // remaining
	codec.Address, // owner
	error,
) {
	return storage.GetOrderFromState(ctx, c.inner.ReadState, orderID)
}
