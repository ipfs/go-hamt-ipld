// Deprecated: use github.com/filecoin-project/go-hamt-ipld
package hamt

import (
	"context"

	upstream "github.com/filecoin-project/go-hamt-ipld"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)

// Deprecated: use github.com/filecoin-project/go-hamt-ipld
var (
	ErrNotFound      = upstream.ErrNotFound
	ErrMaxDepth      = upstream.ErrMaxDepth
	ErrMalformedHamt = upstream.ErrMalformedHamt
)

// Deprecated: use github.com/filecoin-project/go-hamt-ipld
type (
	Node    = upstream.Node
	Pointer = upstream.Pointer
	KV      = upstream.KV
	Option  = upstream.Option
)

// Deprecated: use github.com/filecoin-project/go-hamt-ipld
func UseTreeBitWidth(bitWidth int) Option {
	return upstream.UseTreeBitWidth(bitWidth)
}

// Deprecated: use github.com/filecoin-project/go-hamt-ipld
func UseHashFunction(hash func([]byte) []byte) Option {
	return upstream.UseHashFunction(hash)
}

// Deprecated: use github.com/filecoin-project/go-hamt-ipld
func NewNode(cs cbor.IpldStore, options ...Option) *Node {
	return upstream.NewNode(cs, options...)
}

// Deprecated: use github.com/filecoin-project/go-hamt-ipld
func LoadNode(ctx context.Context, cs cbor.IpldStore, c cid.Cid, options ...Option) (*Node, error) {
	return upstream.LoadNode(ctx, cs, c, options...)
}
