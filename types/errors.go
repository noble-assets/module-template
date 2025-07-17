package types

import "cosmossdk.io/errors"

var (
	ErrUnauthorized = errors.Register(ModuleName, 1, "signer must be the authority")
)
