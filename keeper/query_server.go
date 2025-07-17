package keeper

import (
	"context"

	"template.dev/types"
)

var _ types.QueryServer = &queryServer{}

type queryServer struct {
	*Keeper
}

func NewQueryServer(keeper *Keeper) types.QueryServer {
	return queryServer{Keeper: keeper}
}

// Status implements types.QueryServer.
func (q queryServer) Status(context.Context, *types.QueryStatus) (*types.QueryStatusResponse, error) {
	return &types.QueryStatusResponse{
		IsPaused: false,
	}, nil
}
