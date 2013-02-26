package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

const TouchInputMaxTouchCount = C.ALLEGRO_TOUCH_INPUT_MAX_TOUCH_COUNT

const (
	MouseEmulationNone        = C.ALLEGRO_MOUSE_EMULATION_NONE
	MouseEmulationTransparent = C.ALLEGRO_MOUSE_EMULATION_TRANSPARENT
	MouseEmulationInclusive   = C.ALLEGRO_MOUSE_EMULATION_INCLUSIVE
	MouseEmulationExclusive   = C.ALLEGRO_MOUSE_EMULATION_EXCLUSIVE
	MouseEmulation50X         = C.ALLEGRO_MOUSE_EMULATION_5_0_x
)

type TouchInput C.ALLEGRO_TOUCH_INPUT

type TouchState C.ALLEGRO_TOUCH_STATE

type TouchInputState C.ALLEGRO_TOUCH_INPUT_STATE

func InstallTouchInput() bool {
	return bool(C.al_install_touch_input())
}

func UninstallTouchInput() {
	C.al_uninstall_touch_input()
}

func IsTouchInputInstalled() bool {
	return bool(C.al_is_touch_input_installed())
}

func GetTouchInputState() *TouchInputState {
	tis := new(C.ALLEGRO_TOUCH_INPUT_STATE)
	C.al_get_touch_input_state(tis)
	return (*TouchInputState)(unsafe.Pointer(tis))
}

func SetMouseEmulationMode(mode int32) {
	C.al_set_mouse_emulation_mode(C.int(mode))
}

func GetMouseEmulationMode() int32 {
	return int32(C.al_get_mouse_emulation_mode())
}

func GetTouchInputEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_touch_input_event_source()))
}

func GetTouchInputMouseEmulationEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_touch_input_mouse_emulation_event_source()))
}
