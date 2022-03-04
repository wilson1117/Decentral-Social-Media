package main

import "github.com/cronus6w6/Decentral-Social-Media/pkg/node"

var (
	configFile = "config.json"
)

func main() {
	node.Run(configFile)
}
