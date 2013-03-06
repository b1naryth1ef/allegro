package memfile
// #cgo LDFLAGS: -lallegro_memfile
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_memfile.h>
import "C"

import (
	"unsafe"
	"github.com/tapir/allegro"
)

func OpenMemfile(mem unsafe.Pointer, size int64, mode string) *allegro.File {
	m := C.CString(mode)
	defer C.free(unsafe.Pointer(m))
	return (*allegro.File)(unsafe.Pointer(C.al_open_memfile(mem, C.int64_t(size), m)))
}

func GetAllegroMemfileVersion() uint32 {
	return uint32(C.al_get_allegro_memfile_version())
}
