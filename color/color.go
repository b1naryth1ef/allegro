package color

// #cgo LDFLAGS: -lallegro_color
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_color.h>
import "C"

import (
	"allegro"
	"unsafe"
)

func GetVersion() uint32 {
	return uint32(C.al_get_allegro_color_version())
}

func Cmyk(c, m, y, k float32) *allegro.Color {
	r := C.al_color_cmyk(C.float(c), C.float(m), C.float(y), C.float(k))
	return (*allegro.Color)(unsafe.Pointer(&r))
}

func CmykToRgb(cyan, magenta, yellow, key float32) (float32, float32, float32) {
	var r, g, b C.float
	C.al_color_cmyk_to_rgb(C.float(cyan), C.float(magenta), C.float(yellow), C.float(key), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func RgbToCmyk(r, g, b float32) (float32, float32, float32, float32) {
	var c, m, y, k C.float
	C.al_color_rgb_to_cmyk(C.float(r), C.float(g), C.float(b), &c, &m, &y, &k)
	return float32(c), float32(m), float32(y), float32(k)
}

func Hsl(h, s, l float32) *allegro.Color {
	r := C.al_color_hsl(C.float(h), C.float(s), C.float(l))
	return (*allegro.Color)(unsafe.Pointer(&r))
}

func HslToRgb(h, s, l float32) (float32, float32, float32) {
	var r, g, b C.float
	C.al_color_hsl_to_rgb(C.float(h), C.float(s), C.float(l), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func RgbToHsl(r, g, b float32) (float32, float32, float32) {
	var h, s, l C.float
	C.al_color_rgb_to_hsl(C.float(r), C.float(g), C.float(b), &h, &s, &l)
	return float32(h), float32(s), float32(l)
}

func Hsv(h, s, v float32) *allegro.Color {
	r := C.al_color_hsv(C.float(h), C.float(s), C.float(v))
	return (*allegro.Color)(unsafe.Pointer(&r))
}

func HsvToRgb(h, s, v float32) (float32, float32, float32) {
	var r, g, b C.float
	C.al_color_hsv_to_rgb(C.float(h), C.float(s), C.float(v), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func RgbToHsv(r, g, b float32) (float32, float32, float32) {
	var h, s, v C.float
	C.al_color_rgb_to_hsv(C.float(r), C.float(g), C.float(b), &h, &s, &v)
	return float32(h), float32(s), float32(v)
}

func Html(s string) *allegro.Color {
	hs := C.CString(s)
	defer C.free(unsafe.Pointer(hs))
	r := C.al_color_html(hs)
	return (*allegro.Color)(unsafe.Pointer(&r))
}

func HtmlToRgb(s string) (float32, float32, float32) {
	var r, g, b C.float
	hs := C.CString(s)
	defer C.free(unsafe.Pointer(hs))
	C.al_color_html_to_rgb(hs, &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func RgbToHtml(r, g, b float32) string {
	var cstr *C.char
	defer C.free(unsafe.Pointer(cstr))
	C.al_color_rgb_to_html(C.float(r), C.float(g), C.float(b), cstr)
	return C.GoString(cstr)
}

func Name(name string) *allegro.Color {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	r := C.al_color_name(s)
	return (*allegro.Color)(unsafe.Pointer(&r))
}

func NameToRgb(name string) (float32, float32, float32) {
	var r, g, b C.float
	hs := C.CString(name)
	defer C.free(unsafe.Pointer(hs))
	C.al_color_name_to_rgb(hs, &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func RgbToName(r, g, b float32) string {
	var cstr *C.char
	defer C.free(unsafe.Pointer(cstr))
	C.al_color_rgb_to_html(C.float(r), C.float(g), C.float(b), cstr)
	return C.GoString(cstr)
}

func Yuv(y, u, v float32) *allegro.Color {
	r := C.al_color_yuv(C.float(y), C.float(u), C.float(v))
	return (*allegro.Color)(unsafe.Pointer(&r))
}

func YuvToRgb(y, u, v float32) (float32, float32, float32) {
	var r, g, b C.float
	C.al_color_yuv_to_rgb(C.float(y), C.float(u), C.float(v), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func RgbToYuv(r, g, b float32) (float32, float32, float32) {
	var y, u, v C.float
	C.al_color_rgb_to_yuv(C.float(r), C.float(g), C.float(b), &y, &u, &v)
	return float32(y), float32(u), float32(v)
}
