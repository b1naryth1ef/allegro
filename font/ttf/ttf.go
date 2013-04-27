package ttf

// #cgo LDFLAGS: -lallegro_ttf
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_ttf.h>
import "C"

import (
	"github.com/tapir/allegro"
	"github.com/tapir/allegro/font"
	"unsafe"
)

const (
	NoKerning = C.ALLEGRO_TTF_NO_KERNING
	Monochrome = C.ALLEGRO_TTF_MONOCHROME
	NoAutohint = C.ALLEGRO_TTF_NO_AUTOHINT
)

func InitTtf() bool {
	return bool(C.al_init_ttf_addon())
}

func ShutdownTtf() {
	C.al_shutdown_ttf_addon()
}

func GetVersion() uint32 {
	return uint32(C.al_get_allegro_ttf_version())
}

func LoadFont(filename string, size, flags int32) *font.Font {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return (*font.Font)(unsafe.Pointer(C.al_load_ttf_font(f, C.int(size), C.int(flags))))
}

func LoadFontF(file *allegro.File, filename string, size, flags int32) *font.Font {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return (*font.Font)(unsafe.Pointer(C.al_load_ttf_font_f((*C.ALLEGRO_FILE)(unsafe.Pointer(file)), f, C.int(size), C.int(flags))))
}

func LoadFontStretch(filename string, w, h, flags int32) *font.Font {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return (*font.Font)(unsafe.Pointer(C.al_load_ttf_font_stretch(f, C.int(w), C.int(h), C.int(flags))))
}

func LoadFontStretchF(file *allegro.File, filename string, w, h, flags int32) *font.Font {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return (*font.Font)(unsafe.Pointer(C.al_load_ttf_font_stretch_f((*C.ALLEGRO_FILE)(unsafe.Pointer(file)), f, C.int(w), C.int(h), C.int(flags))))
}
