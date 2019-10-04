package arrogant

// #include <modest/modest.h>
import "C"
import "unsafe"

// NodeRange struct
type NodeRange struct {
	myhtmlCollection *C.myhtml_collection_t
	myhtmlTree       *C.myhtml_tree_t
	parent           *Node
}

// First node
func (r *NodeRange) First() *Node {
	return r.At(0)
}

// At get node at given index
func (r *NodeRange) At(index int) *Node {
	length := int(r.myhtmlCollection.length)
	if length <= index {
		return nil
	}

	nodes := (*[1 << 30]*C.myhtml_tree_node_t)(unsafe.Pointer(r.myhtmlCollection.list))[:length:length]

	return &Node{nodes[index], r.parent.parent}
}
