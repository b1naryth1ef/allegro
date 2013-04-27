package font

// #cgo LDFLAGS: -lallegro_font
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_font.h>
import "C"

import (
	"github.com/tapir/allegro"
	"unsafe"
)

const (
	AlignLeft   = C.ALLEGRO_ALIGN_LEFT
	AlignCentre = C.ALLEGRO_ALIGN_CENTRE
	AlignRight  = C.ALLEGRO_ALIGN_RIGHT
	AlignInteger = C.ALLEGRO_ALIGN_INTEGER
)

type Font C.ALLEGRO_FONT

func InitFont() bool {
	return bool(C.al_init_font_addon())
}

func ShutdownFont() {
	C.al_shutdown_font_addon()
}

func GetVersion() uint32 {
	return uint32(C.al_get_allegro_font_version())
}

func LoadFont(fileName string, size, flags int32) *Font {
	f := C.CString(fileName)
	defer C.free(unsafe.Pointer(f))
	return (*Font)(unsafe.Pointer(C.al_load_font(f, C.int(size), C.int(flags))))
}

func (f *Font) Destroy() {
	C.al_destroy_font((*C.ALLEGRO_FONT)(unsafe.Pointer(f)))
}

func (f *Font) GetLineHeight() int32 {
	return int32(C.al_get_font_line_height((*C.ALLEGRO_FONT)(unsafe.Pointer(f))))
}

func (f *Font) GetAscent() int32 {
	return int32(C.al_get_font_ascent((*C.ALLEGRO_FONT)(unsafe.Pointer(f))))
}

func (f *Font) GetDescent() int32 {
	return int32(C.al_get_font_descent((*C.ALLEGRO_FONT)(unsafe.Pointer(f))))
}

func (f *Font) GetTextWidth(text string) int32 {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	return int32(C.al_get_text_width((*C.ALLEGRO_FONT)(unsafe.Pointer(f)), t))
}

func (f *Font) DrawText(c *allegro.Color, x, y float32, flags int32, text string) {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	C.al_draw_text((*C.ALLEGRO_FONT)(unsafe.Pointer(f)), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(x), C.float(y), C.int(flags), t)
}

func (f *Font) DrawJustifiedText(c *allegro.Color, x1, x2, y, diff float32, flags int32, text string) {
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	C.al_draw_justified_text((*C.ALLEGRO_FONT)(unsafe.Pointer(f)), *((*C.ALLEGRO_COLOR)(unsafe.Pointer(c))), C.float(x1), C.float(x2), C.float(y), C.float(diff), C.int(flags), t)
}

//func (f *Font) DrawTextf()

//func (f *Font) DrawJustifiedTextf()

func (f *Font) GetTextDimensions(text string) (int32, int32, int32, int32) {
	var x, y, w, h C.int
	t := C.CString(text)
	defer C.free(unsafe.Pointer(t))
	C.al_get_text_dimensions((*C.ALLEGRO_FONT)(unsafe.Pointer(f)), t, &x, &y, &w, &h)
	return int32(x), int32(y), int32(w), int32(h)
}

//func (f *Font) GetRanges()

/****************/
/* Bitmap fonts */
/****************/

//func (f *Font) GrabFromBitmap()

func LoadBitmapFont(fname string) *Font {
	f := C.CString(fname)
	defer C.free(unsafe.Pointer(f))
	return (*Font)(unsafe.Pointer(C.al_load_bitmap_font(f)))
}

func LoadBitmapFontFlags(fname string, flags int32) *Font {
	f := C.CString(fname)
	defer C.free(unsafe.Pointer(f))
	return (*Font)(unsafe.Pointer(C.al_load_bitmap_font_flags(f, C.int(flags))))
}

func CreateBuiltinFont() *Font {
	return (*Font)(unsafe.Pointer(C.al_create_builtin_font()))
}
