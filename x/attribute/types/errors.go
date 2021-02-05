package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/attribute module errors
var (
	// ErrInvalidAddress indicates the address given does not match an existing account.
	ErrInvalidAddress = sdkerrors.Register(ModuleName, 1, "address does not match an existing account")
	// ErrAttributeContainsSegments indicates a multi-segment name in a single segment context.
	ErrAttributeContainsSegments = sdkerrors.Register(ModuleName, 2, "invalid attribute: \".\" is reserved")
)
