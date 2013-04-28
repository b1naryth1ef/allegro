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

/***********************/
/* DisplayMode getters */
/***********************/

func (d *DisplayMode) Width() int32{
	dc := (*C.ALLEGRO_DISPLAY_MODE)(unsafe.Pointer(d))
	return int32(dc.width)
}

func (d *DisplayMode) Height() int32{
	dc := (*C.ALLEGRO_DISPLAY_MODE)(unsafe.Pointer(d))
	return int32(dc.height)
}

func (d *DisplayMode) Format() int32{
	dc := (*C.ALLEGRO_DISPLAY_MODE)(unsafe.Pointer(d))
	return int32(dc.format)
}

func (d *DisplayMode) RefreshRate() int32{
	dc := (*C.ALLEGRO_DISPLAY_MODE)(unsafe.Pointer(d))
	return int32(dc.refresh_rate)
}
