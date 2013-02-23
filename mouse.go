package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

const (
	SystemMouseCursorNone        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_NONE
	SystemMouseCursorDefault     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_DEFAULT
	SystemMouseCursorArrow       = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ARROW
	SystemMouseCursorBusy        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_BUSY
	SystemMouseCursorQuestion    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_QUESTION
	SystemMouseCursorEdit        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_EDIT
	SystemMouseCursorMove        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_MOVE
	SystemMouseCursorResizeN     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_N
	SystemMouseCursorResizeW     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_W
	SystemMouseCursorResizeS     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_S
	SystemMouseCursorResizeE     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_E
	SystemMouseCursorResizeNW    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NW
	SystemMouseCursorResizeSW    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SW
	SystemMouseCursorResizeSE    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SE
	SystemMouseCursorResizeNE    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NE
	SystemMouseCursorProgress    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PROGRESS
	SystemMouseCursorPrecision   = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PRECISION
	SystemMouseCursorLink        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_LINK
	SystemMouseCursorSelect      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ALT_SELECT
	SystemMouseCursorUnavailable = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_UNAVAILABLE
)

type Mouse C.ALLEGRO_MOUSE

type MouseState struct {
	X        int32
	Y        int32
	Z        int32
	W        int32
	moreAxes [C.ALLEGRO_MOUSE_MAX_EXTRA_AXES]int32
	Buttons  int32
	pressure int32
	disp     *Display
}

func InstallMouse() bool {
	return bool(C.al_install_mouse())
}

func IsMouseInstalled() bool {
	return bool(C.al_is_mouse_installed())
}

func UninstallMouse() {
	C.al_uninstall_mouse()
}

func GetMouseNumAxes() uint32 {
	return uint32(C.al_get_mouse_num_axes())
}

func GetMouseNumButtons() uint32 {
	return uint32(C.al_get_mouse_num_buttons())
}

func GetMouseState() *MouseState {
	ms := new(C.ALLEGRO_MOUSE_STATE)
	C.al_get_mouse_state(ms)
	return (*MouseState)(unsafe.Pointer(ms))
}

func (m *MouseState) GetAxis(axis int32) int32 {
	return int32(C.al_get_mouse_state_axis((*C.ALLEGRO_MOUSE_STATE)(unsafe.Pointer(m)), C.int(axis)))
}

func (m *MouseState) ButtonDown(button int32) bool {
	return bool(C.al_mouse_button_down((*C.ALLEGRO_MOUSE_STATE)(unsafe.Pointer(m)), C.int(button)))
}

func (d *Display) SetMouseXY(x, y int32) bool {
	return bool(C.al_set_mouse_xy((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(x), C.int(y)))
}

func SetMouseZ(z int32) bool {
	return bool(C.al_set_mouse_z(C.int(z)))
}

func SetMouseW(w int32) bool {
	return bool(C.al_set_mouse_w(C.int(w)))
}

func SetMouseAxis(which, value int32) bool {
	return bool(C.al_set_mouse_axis(C.int(which), C.int(value)))
}

func GetMouseEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_mouse_event_source()))
}

/*****************/
/* Mouse Cursors */
/*****************/

type MouseCursor C.ALLEGRO_MOUSE_CURSOR

func CreateMouseCursor(bmp *Bitmap, xFocus, yFocus int32) *MouseCursor {
	return (*MouseCursor)(unsafe.Pointer(C.al_create_mouse_cursor((*C.ALLEGRO_BITMAP)(unsafe.Pointer(bmp)), C.int(xFocus), C.int(yFocus))))
}

func (m *MouseCursor) Destroy() {
	C.al_destroy_mouse_cursor((*C.ALLEGRO_MOUSE_CURSOR)(unsafe.Pointer(m)))
}

func (m *MouseCursor) Set(display *Display) bool {
	return bool(C.al_set_mouse_cursor((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)), (*C.ALLEGRO_MOUSE_CURSOR)(unsafe.Pointer(m))))
}

func SetSystemMouseCursor(display *Display, cursorId uint32) bool {
	return bool(C.al_set_system_mouse_cursor((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)), (C.ALLEGRO_SYSTEM_MOUSE_CURSOR)(cursorId)))
}

func GetMouseCursorPosition() (bool, int32, int32) {
	var x, y C.int
	r := bool(C.al_get_mouse_cursor_position(&x, &y))
	return r, int32(x), int32(y)
}

func HideMouseCursor(display *Display) bool {
	return bool(C.al_hide_mouse_cursor((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display))))
}

func ShowMouseCursor(display *Display) bool {
	return bool(C.al_show_mouse_cursor((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display))))
}

func GrabMouse(display *Display) bool {
	return bool(C.al_grab_mouse((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display))))
}

func UngrabMouse() bool {
	return bool(C.al_ungrab_mouse())
}
