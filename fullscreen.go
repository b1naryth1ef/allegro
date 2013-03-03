package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func GetDisplayMode(index int32) *DisplayMode {
	var d C.ALLEGRO_DISPLAY_MODE
	r := C.al_get_display_mode(C.int(index), &d)
	return (*DisplayMode)(unsafe.Pointer(r))
}

func GetNumDisplayModes() int32 {
	return int32(C.al_get_num_display_modes())
}
