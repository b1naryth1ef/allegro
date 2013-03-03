package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

/********************/
/* Display Creation */
/********************/

func CreateDisplay(w, h int32) *Display {
	return (*Display)(unsafe.Pointer(C.al_create_display(C.int(w), C.int(h))))
}

func (d *Display) Destroy() {
	C.al_destroy_display((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))
}

func GetNewDisplayFlags() int32 {
	return int32(C.al_get_new_display_flags())
}

func SetNewDisplayFlags(flags int32) {
	C.al_set_new_display_flags(C.int(flags))
}

func GetNewDisplayOption(option int32) (int32, int32) {
	var importance C.int
	v := int32(C.al_get_new_display_option(C.int(option), &importance))
	return v, int32(importance)
}

func SetNewDisplayOption(option, value, importance int32) {
	C.al_set_new_display_option(C.int(option), C.int(value), C.int(importance))
}

func ResetNewDisplayOptions() {
	C.al_reset_new_display_options()
}

func GetNewWindowPosition() (int32, int32) {
	var x, y C.int
	C.al_get_new_window_position(&x, &y)
	return int32(x), int32(y)
}

func SetNewWindowPosition(x, y int32) {
	C.al_set_new_window_position(C.int(x), C.int(y))
}

func GetNewDisplayRefreshRate() int32 {
	return int32(C.al_get_new_display_refresh_rate())
}

func SetNewDisplayRefreshRate(refreshRate int32) {
	C.al_set_new_display_refresh_rate(C.int(refreshRate))
}

/**********************/
/* Display Operations */
/**********************/

func (d *Display) GetEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_display_event_source((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))))
}

func (d *Display) GetBackbuffer() *Bitmap {
	return (*Bitmap)(unsafe.Pointer(C.al_get_backbuffer((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))))
}

func FlipDisplay() {
	C.al_flip_display()
}

func UpdateDisplayRegion(x, y, width, height int32) {
	C.al_update_display_region(C.int(x), C.int(y), C.int(width), C.int(height))
}

func WaitForVsync() bool {
	return bool(C.al_wait_for_vsync())
}

/*****************************/
/* Display Size and Position */
/*****************************/

func (d *Display) GetWidth() int32 {
	return int32(C.al_get_display_width((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) GetHeight() int32 {
	return int32(C.al_get_display_height((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) Resize(width, height int32) bool {
	return bool(C.al_resize_display((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(width), C.int(height)))
}

func (d *Display) AcknowledgeResize() bool {
	return bool(C.al_acknowledge_resize((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) GetWindowPosition() (int32, int32) {
	var x, y C.int
	C.al_get_window_position((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), &x, &y)
	return int32(x), int32(y)
}

func (d *Display) SetWindowPosition(x, y int32) {
	C.al_set_window_position((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(x), C.int(y))
}

func (d *Display) GetWindowConstraints() (bool, int32, int32, int32, int32) {
	var min_w, min_h, max_w, max_h C.int
	r := C.al_get_window_constraints((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), &min_w, &min_h, &max_w, &max_h)
	return bool(r), int32(min_w), int32(min_h), int32(max_w), int32(max_h)
}

func (d *Display) SetWindowConstraints(min_w, min_h, max_w, max_h int32) bool {
	return bool(C.al_set_window_constraints((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(min_w), C.int(min_h), C.int(max_w), C.int(max_h)))
}

/********************/
/* Display Settings */
/********************/

func (d *Display) GetFlags() int32 {
	return int32(C.al_get_display_flags((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) SetFlag(flag int, onoff bool) bool {
	return bool(C.al_set_display_flag((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(flag), C.bool(onoff)))
}

func (d *Display) GetOption(option int32) int32 {
	return int32(C.al_get_display_option((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(option)))
}

func (d *Display) SetOption(option, value int32) {
	C.al_set_display_option((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(option), C.int(value))
}

func (d *Display) GetFormat() int32 {
	return int32(C.al_get_display_format((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) GetOrientation() int32 {
	return int32(C.al_get_display_orientation((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) GetRefreshRate() int32 {
	return int32(C.al_get_display_refresh_rate((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) SetWindowTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.al_set_window_title((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), t)
}

func (d *Display) SetIcon(icon *Bitmap) {
	C.al_set_display_icon((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), (*C.ALLEGRO_BITMAP)(unsafe.Pointer(icon)))
}

//func (d *Display) SetIcons(numIcons int32, icons [](*Bitmap)) {
//	C.al_set_display_icons(C.int(numIcons), (**C.ALLEGRO_BITMAP)(unsafe.Pointer(&icons[0])))
//}

/*****************/
/* Drawing Halts */
/*****************/

func (d *Display) AcknowledgeDrawingHalt() {
	C.al_acknowledge_drawing_halt((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))
}

func (d *Display) AcknowledgeDrawingResume() {
	C.al_acknowledge_drawing_resume((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), nil)
}

/***************/
/* Screensaver */
/***************/

func InhibitScreensaver(inhibit bool) bool {
	return bool(C.al_inhibit_screensaver(C.bool(inhibit)))
}
