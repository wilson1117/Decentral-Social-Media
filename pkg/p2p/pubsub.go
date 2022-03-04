package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type PubSub struct {
	ps   *pubsub.PubSub
	host host.Host
	ctx  context.Context
}

func NewPubSub(ctx context.Context, h host.Host) (*PubSub, error) {
	ps := &PubSub{
		ctx:  ctx,
		host: h,
	}
	var err error
	if ps.ps, err = pubsub.NewGossipSub(ctx, h); err != nil {
		return nil, err
	}

	return ps, nil
}

func (p PubSub) Subscribe(topic string, bufferSize ...int) (chan []byte, error) {
	var bs int

	if len(bufferSize) > 0 {
		bs = bufferSize[0]
	} else {
		bs = 1024
	}

	t, err := p.ps.Join(topic)
	if err != nil {
		return nil, err
	}

	sub, err := t.Subscribe()
	if err != nil {
		return nil, err
	}

	mc := make(chan []byte, bs)

	go p.subReadLoop(sub, mc)

	return mc, nil
}

func (p *PubSub) subReadLoop(sub *pubsub.Subscription, mc chan []byte) {
	for {
		msg, err := sub.Next(p.ctx)
		if err != nil {
			close(mc)
			return
		}

		if msg.ReceivedFrom == p.host.ID() {
			continue
		}

		mc <- msg.Data
	}
}
