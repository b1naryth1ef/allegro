package allegro

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_opengl.h>
import "C"

import "unsafe"

func (b *Bitmap) GetOpenglTexture() uint32 {
	return uint32(C.al_get_opengl_texture((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) GetOpenglTextureSize() (int32, int32) {
	var w, h C.int
	C.al_get_opengl_texture_size((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), &w, &h)
	return int32(w), int32(h)
}

func (b *Bitmap) GetOpenglTexturePosition() (int32, int32) {
	var u, v C.int
	C.al_get_opengl_texture_position((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), &u, &v)
	return int32(u), int32(v)
}

func (b *Bitmap) GetOpenglFbo() uint32 {
	return uint32(C.al_get_opengl_fbo((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) RemoveOpenglFbo() {
	C.al_remove_opengl_fbo((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)))
}

func HaveOpenglExtension(extension string) bool {
	e := C.CString(extension)
	defer C.free(unsafe.Pointer(e))
	return bool(C.al_have_opengl_extension(e))
}

func GetOpenglVersion() uint32 {
	return uint32(C.al_get_opengl_version())
}

func GetOpenglVariant() int32 {
	return int32(C.al_get_opengl_variant())
}

func SetCurrentOpenglContext(display *Display) {
	C.al_set_current_opengl_context((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)))
}
