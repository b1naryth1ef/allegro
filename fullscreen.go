package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

type DisplayMode struct {
	Width       int
	Height      int
	Format      int
	RefreshRate int
}

func GetDisplayMode(index int) *DisplayMode {
	var d C.ALLEGRO_DISPLAY_MODE
	r := C.al_get_display_mode(C.int(index), &d)
	return (*DisplayMode)(unsafe.Pointer(r))
}

func GetNumDisplayModes() int {
	return int(C.al_get_num_display_modes())
}
