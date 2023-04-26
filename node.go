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
	allNodes map[NodeID]Node
}


var Nodes = nodes{
	allNodes: make(map[NodeID]Node),
}


	
func (nd *Nodes) GetNode(NodeID uint64) Node	{
	return nd.allNodes[NodeID]

}


func (nd *Nodes) GetNodes() map[NodeID]Node {
	return nd.allNodes
}

func (nd *Nodes) AddNode(newNode Node) {
	nd.allNodes[newNode.NodeID] = newNode
}