package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

type DHT struct {
	dht *dht.IpfsDHT
}

func New(ctx context.Context, h host.Host) (*DHT, error) {
	d, err := dht.New(ctx, h)
	if err != nil {
		return nil, err
	}

	return &DHT{
		dht: d,
	}, nil
}
