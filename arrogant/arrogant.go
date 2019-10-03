package arrogant

// #cgo CFLAGS: -I../modest/include
// #cgo LDFLAGS: -L../modest/lib -lmodest_static -lm
// #include <modest/modest.h>
import "C"
import "errors"

const (
	opt         uint32  = 0
	threadCount C.ulong = 1
	queueSize   C.ulong = 0

	encoding C.myencoding_t = 0
)

// Arrogant struct
type Arrogant struct {
	myhtml *C.myhtml_t
	valid  bool
}

// New arrogant instance
func New() *Arrogant {
	return &Arrogant{}
}

// Parse html string
func (a *Arrogant) Parse(html string) (*Tree, error) {
	if a.myhtml == nil {
		a.myhtml = C.myhtml_create()
		if status := C.myhtml_init(a.myhtml, opt, threadCount, queueSize); status != 0 {
			return nil, errors.New("myhtml_init failed")
		}
		a.valid = true
	}
	tree, err := NewTree(a)
	if err != nil {
		return nil, err
	}
	if err := tree.Parse(html, encoding); err != nil {
		return nil, err
	}

	return tree, nil
}

// // ParseFragment TODO
// func (a *Arrogant) ParseFragment(html string) (*Tree, error) {
// 	return nil, nil
// }

// Release pointer
func (a *Arrogant) Release() {
	if a.valid {
		C.myhtml_destroy(a.myhtml)
		a.valid = false
	}
}
