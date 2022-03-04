package node

import (
	"github.com/cronus6w6/Decentral-Social-Media/core/p2p"
	"github.com/cronus6w6/Decentral-Social-Media/pkg/config"
)

type Node struct {
	config *config.NodeConfig
	P2P    *p2p.P2P
}

func New(configFilePath string) (*Node, error) {
	node := &Node{}
	var err error

	if node.config, err = config.New(configFilePath); err != nil {
		return nil, err
	}

	if node.P2P, err = p2p.New(node.config); err != nil {
		return nil, err
	}

	return node, nil
}

func Run(configFilePath string) (*Node, error) {
	node, err := New(configFilePath)

	if err != nil {
		return nil, err
	}

	return node, err
}
