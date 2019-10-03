package arrogant

// #include <modest/modest.h>
// #include <myhtml/serialization.h>
import "C"
import (
	"errors"
)

// Node struct
type Node struct {
	myhtmlTreeNode *C.myhtml_tree_node_t
	parent         *Tree
}

// ByTagName get elements
func (n *Node) ByTagName(name string) (*NodeRange, error) {
	var status C.mystatus_t
	myhtmlCollection := C.myhtml_collection_create(0, nil)
	collection := NodeRange{
		C.myhtml_get_nodes_by_name_in_scope(
			n.parent.myhtmlTree,
			myhtmlCollection,
			n.myhtmlTreeNode,
			C.CString(name),
			C.ulong(len(name)),
			&status),
		n.parent.myhtmlTree,
		n,
	}
	if status != 0 {
		return nil, errors.New("myhtml_get_nodes_by_name_in_scope failed")
	}
	return &collection, nil
}

// String html
func (n *Node) String() string {
	var raw C.mycore_string_raw_t
	C.mycore_string_raw_clean_all(&raw)
	defer C.mycore_string_raw_destroy(&raw, false)

	if C.myhtml_serialization_tree_buffer(n.myhtmlTreeNode, &raw) != 0 {
		return ""
	}

	return C.GoString(raw.data)
}

// InnerText set text
func (n *Node) InnerText(s string) {
	node := C.myhtml_node_create(n.parent.myhtmlTree, 1, 1)
	C.myhtml_node_text_set(node, C.CString(s), C.ulong(len(s)), 0)
	C.myhtml_node_remove(node) // detach

	// delete all children
	for current := C.myhtml_node_child(n.myhtmlTreeNode); current != nil; {
		C.myhtml_node_delete_recursive(current)
		current = C.myhtml_node_next(current)
	}

	C.myhtml_node_append_child(n.myhtmlTreeNode, node)
}
