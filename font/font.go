package font

// #cgo LDFLAGS: -lallegro_font -lallegro_ttf
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_font.h>
// #include <allegro5/allegro_ttf.h>
import "C"

import (
	"github.com/tapir/allegro"
	"unsafe"	
)

const (
	AlignLeft   = C.ALLEGRO_ALIGN_LEFT
	AlignCentre = C.ALLEGRO_ALIGN_CENTRE
	AlignRight  = C.ALLEGRO_ALIGN_RIGHT
)

type Font C.ALLEGRO_FONT

func InitFont() {
	C.al_init_font_addon()
}

func InitTTF() bool {
	return bool(C.al_init_ttf_addon())
}

func LoadFont(fileName string, size, flags int) *Font {
	f := C.CString(fileName)
	defer C.free(unsafe.Pointer(f))
	return (*Font)(unsafe.Pointer(C.al_load_font(f, C.int(size), C.int(flags))))
}

func (f *Font) Destroy() {
	C.al_destroy_font((*C.ALLEGRO_FONT)(unsafe.Pointer(f)))
}

func (f *Font) DrawText(text string, c *allegro.Color, x, y float32, flags int) {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	C.al_draw_text((*C.ALLEGRO_FONT)(unsafe.Pointer(f)), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(x), C.float(y), C.int(flags), t)
}

func (f *Font) TextWidth(text string) int {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	return int(C.al_get_text_width((*C.ALLEGRO_FONT)(unsafe.Pointer(f)), t))
}

func (f *Font) TextDimensions(text string) (int, int, int, int) {
	var x, y, w, h C.int
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	C.al_get_text_dimensions((*C.ALLEGRO_FONT)(unsafe.Pointer(f)), t, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}
