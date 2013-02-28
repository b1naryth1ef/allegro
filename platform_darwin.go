package allegro

// #include <allegro5/allegro.h>
import "C"

func (d *Display) OsxGetWindow *C.NSWindow {
	return C.al_osx_get_window((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))
}
