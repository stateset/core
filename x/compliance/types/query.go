package types

import "context"

// QueryServer defines the gRPC query server interface for the compliance module.
type QueryServer interface {
	Profile(context.Context, *QueryProfileRequest) (*QueryProfileResponse, error)
	Profiles(context.Context, *QueryProfilesRequest) (*QueryProfilesResponse, error)
	ProfilesByStatus(context.Context, *QueryProfilesByStatusRequest) (*QueryProfilesByStatusResponse, error)
}

// QueryProfileRequest is the request for a single profile
type QueryProfileRequest struct {
	Address string `json:"address"`
}

// QueryProfileResponse is the response containing a profile
type QueryProfileResponse struct {
	Profile Profile `json:"profile"`
}

// QueryProfilesRequest is the request for all profiles
type QueryProfilesRequest struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

// QueryProfilesResponse is the response containing profiles
type QueryProfilesResponse struct {
	Profiles []Profile `json:"profiles"`
	Total    uint64    `json:"total"`
}

// QueryProfilesByStatusRequest is the request for profiles by status
type QueryProfilesByStatusRequest struct {
	Status ProfileStatus `json:"status"`
	Limit  uint64        `json:"limit"`
	Offset uint64        `json:"offset"`
}

// QueryProfilesByStatusResponse is the response containing filtered profiles
type QueryProfilesByStatusResponse struct {
	Profiles []Profile `json:"profiles"`
	Total    uint64    `json:"total"`
}
