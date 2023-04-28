package Node
import (
	"crypto/ed25519"

)

type VidarNode struct{
	NodeID				uint64
	PubKey				ed25519.PublicKey		
    RingId				uint64
	Tor					string
	I2p					string

}


type nodes struct {
	allNodes map[uint64]VidarNode
}


var Nodes = nodes{
	allNodes: make(map[uint64]VidarNode),
}


	
func (nd *nodes) GetNode(NodeID uint64) VidarNode	{
	return nd.allNodes[NodeID]

}


func (nd *nodes) GetNodes() map[uint64]VidarNode {
	return nd.allNodes
}

func (nd *nodes) AddNode(newNode VidarNode) {
	nd.allNodes[newNode.NodeID] = newNode
}