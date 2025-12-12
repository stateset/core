package types

import proto "github.com/cosmos/gogoproto/proto"

// String implements gogoproto's proto.Message interface for Params.
func (m Params) String() string { return proto.CompactTextString(&m) }
