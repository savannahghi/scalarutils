package go_utils

import (
	"fmt"
)

// ID is fulfilled by all stringifiable types.
// A valid Relay ID must fulfill this interface.
type ID interface {
	fmt.Stringer
}

// Node is a Relay (GraphQL Relay) node.
// Any valid type in this server should be a node.
type Node interface {
	IsNode()
	GetID() ID
	SetID(string)
}

// IDValue represents GraphQL object identifiers
type IDValue string

func (val IDValue) String() string { return string(val) }

// PageInfo is used to add pagination information to Relay edges.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

//IsEntity ...
func (p PageInfo) IsEntity() {}

// NewString returns a pointer to the supplied string.
func NewString(s string) *string {
	return &s
}
