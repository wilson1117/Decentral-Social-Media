package p2p

import (
	"bufio"
	"context"
	"fmt"

	"github.com/cronus6w6/Decentral-Social-Media/pkg/config"
	"github.com/cronus6w6/Decentral-Social-Media/pkg/p2p"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

type P2P struct {
	Host   host.Host
	Config map[string]interface{}
	Ctx    context.Context
	PubSub *p2p.PubSub
}

func New(config *config.NodeConfig) (*P2P, error) {
	p := &P2P{
		Config: map[string]interface{}{
			"bindMultiAddresses": []string{},
		},
		Ctx: context.Background(),
	}

	config.GetConfig("p2p", &p.Config)
	var err error
	if p.Host, err = libp2p.New(
		libp2p.ListenAddrStrings(p.Config["bindMultiAddresses"].([]string)...),
	); err != nil {
		return nil, err
	}

	if p.PubSub, err = p2p.NewPubSub(p.Ctx, p.Host); err != nil {
		return nil, err
	}

	p.Info()

	return p, nil
}

func (p *P2P) Info() {
	peerInfo := peerstore.AddrInfo{
		ID:    p.Host.ID(),
		Addrs: p.Host.Addrs(),
	}
	addrs, _ := peerstore.AddrInfoToP2pAddrs(&peerInfo)

	println("libp2p node address:")
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}

func (p *P2P) Handle(pid protocol.ID, handler func(*bufio.ReadWriter)) {
	p.Host.SetStreamHandler(pid, func(s network.Stream) {
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
		go handler(rw)
	})
}
