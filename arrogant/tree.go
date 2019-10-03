package arrogant

// #include <modest/modest.h>
import "C"
import "errors"

// Tree struct
type Tree struct {
	myhtmlTree *C.myhtml_tree_t
	valid      bool
	parent     *Arrogant
}

// NewTree create instance
func NewTree(parent *Arrogant) (*Tree, error) {
	tree := C.myhtml_tree_create()
	if status := C.myhtml_tree_init(tree, parent.myhtml); status != 0 {
		return nil, errors.New("mythml_tree_init failed")
	}

	return &Tree{tree, true, parent}, nil
}

// Parse html
func (t *Tree) Parse(html string, encoding C.myencoding_t) error {
	if status := C.myhtml_parse(t.myhtmlTree, encoding, C.CString(html), C.ulong(len(html))); status != 0 {
		return errors.New("myhtml_parse failed")
	}

	return nil
}

// Release pointer
func (t *Tree) Release() {
	if t.valid {
		C.myhtml_tree_destroy(t.myhtmlTree)
		t.parent.Release()
		t.valid = false
	}
}

// Document get tree document
func (t *Tree) Document() *Node {
	return &Node{C.myhtml_tree_get_document(t.myhtmlTree), t}
}

// ByTagName get elements
func (t *Tree) ByTagName(name string) (*NodeRange, error) {
	return t.Document().ByTagName(name)
}
