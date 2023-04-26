package Node
import (
	"crypto/ed25519"

)

type Node struct{
	NodeID				uint64
	PubKey				ed25519.PublicKey		
    RingId				uint64
	Tor					string
	I2p					string

}


type nodes struct {
	allNodes map[uint64]Node
}


var Nodes = nodes{
	allNodes: make(map[uint64]Node),
}


	
func (nd *nodes) GetNode(NodeID uint64) Node	{
	return nd.allNodes[NodeID]

}


func (nd *nodes) GetNodes() map[uint64]Node {
	return nd.allNodes
}

func (nd *nodes) AddNode(newNode Node) {
	nd.allNodes[newNode.NodeID] = newNode
}