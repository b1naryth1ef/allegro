package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func InstallKeyboard() bool {
	return bool(C.al_install_keyboard())
}

func IsKeyboardInstalled() bool {
	return bool(C.al_is_keyboard_installed())
}

func UninstallKeyboard() {
	C.al_uninstall_keyboard()
}

func GetKeyboardState() *KeyboardState {
	ks := new(C.ALLEGRO_KEYBOARD_STATE)
	C.al_get_keyboard_state(ks)
	return (*KeyboardState)(unsafe.Pointer(ks))
}

func (k *KeyboardState) KeyDown(keycode int32) bool {
	return bool(C.al_key_down((*C.ALLEGRO_KEYBOARD_STATE)(unsafe.Pointer(k)), C.int(keycode)))
}

func KeycodeToName(keycode int32) string {
	return C.GoString(C.al_keycode_to_name(C.int(keycode)))
}

func SetKeyboardLeds(leds int32) bool {
	return bool(C.al_set_keyboard_leds(C.int(leds)))
}

func GetKeyboardEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_keyboard_event_source()))
}
