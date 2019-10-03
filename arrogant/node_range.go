package arrogant

// #include <modest/modest.h>
import "C"

// NodeRange struct
type NodeRange struct {
	myhtmlCollection *C.myhtml_collection_t
	myhtmlTree       *C.myhtml_tree_t
	parent           *Node
}

// First node
func (r *NodeRange) First() *Node {
	// TODO: check length
	node := C.myhtml_collection_get(r.myhtmlCollection, 0)
	return &Node{node, r.parent.parent}
}
