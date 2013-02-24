package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

type Keyboard C.ALLEGRO_KEYBOARD

type KeyboardState C.ALLEGRO_KEYBOARD_STATE

const (
	KeyA             = C.ALLEGRO_KEY_A
	KeyB             = C.ALLEGRO_KEY_B
	KeyC             = C.ALLEGRO_KEY_C
	KeyD             = C.ALLEGRO_KEY_D
	KeyE             = C.ALLEGRO_KEY_E
	KeyF             = C.ALLEGRO_KEY_F
	KeyG             = C.ALLEGRO_KEY_G
	KeyH             = C.ALLEGRO_KEY_H
	KeyI             = C.ALLEGRO_KEY_I
	KeyJ             = C.ALLEGRO_KEY_J
	KeyK             = C.ALLEGRO_KEY_K
	KeyL             = C.ALLEGRO_KEY_L
	KeyM             = C.ALLEGRO_KEY_M
	KeyN             = C.ALLEGRO_KEY_N
	KeyO             = C.ALLEGRO_KEY_O
	KeyP             = C.ALLEGRO_KEY_P
	KeyQ             = C.ALLEGRO_KEY_Q
	KeyR             = C.ALLEGRO_KEY_R
	KeyS             = C.ALLEGRO_KEY_S
	KeyT             = C.ALLEGRO_KEY_T
	KeyU             = C.ALLEGRO_KEY_U
	KeyV             = C.ALLEGRO_KEY_V
	KeyW             = C.ALLEGRO_KEY_W
	KeyX             = C.ALLEGRO_KEY_X
	KeyY             = C.ALLEGRO_KEY_Y
	KeyZ             = C.ALLEGRO_KEY_Z
	Key0             = C.ALLEGRO_KEY_0
	Key1             = C.ALLEGRO_KEY_1
	Key2             = C.ALLEGRO_KEY_2
	Key3             = C.ALLEGRO_KEY_3
	Key4             = C.ALLEGRO_KEY_4
	Key5             = C.ALLEGRO_KEY_5
	Key6             = C.ALLEGRO_KEY_6
	Key7             = C.ALLEGRO_KEY_7
	Key8             = C.ALLEGRO_KEY_8
	Key9             = C.ALLEGRO_KEY_9
	KeyPad0          = C.ALLEGRO_KEY_PAD_0
	KeyPad1          = C.ALLEGRO_KEY_PAD_1
	KeyPad2          = C.ALLEGRO_KEY_PAD_2
	KeyPad3          = C.ALLEGRO_KEY_PAD_3
	KeyPad4          = C.ALLEGRO_KEY_PAD_4
	KeyPad5          = C.ALLEGRO_KEY_PAD_5
	KeyPad6          = C.ALLEGRO_KEY_PAD_6
	KeyPad7          = C.ALLEGRO_KEY_PAD_7
	KeyPad8          = C.ALLEGRO_KEY_PAD_8
	KeyPad9          = C.ALLEGRO_KEY_PAD_9
	KeyF1            = C.ALLEGRO_KEY_F1
	KeyF2            = C.ALLEGRO_KEY_F2
	KeyF3            = C.ALLEGRO_KEY_F3
	KeyF4            = C.ALLEGRO_KEY_F4
	KeyF5            = C.ALLEGRO_KEY_F5
	KeyF6            = C.ALLEGRO_KEY_F6
	KeyF7            = C.ALLEGRO_KEY_F7
	KeyF8            = C.ALLEGRO_KEY_F8
	KeyF9            = C.ALLEGRO_KEY_F9
	KeyF10           = C.ALLEGRO_KEY_F10
	KeyF11           = C.ALLEGRO_KEY_F11
	KeyF12           = C.ALLEGRO_KEY_F12
	KeyEscape        = C.ALLEGRO_KEY_ESCAPE
	KeyTilde         = C.ALLEGRO_KEY_TILDE
	KeyMinus         = C.ALLEGRO_KEY_MINUS
	KeyEquals        = C.ALLEGRO_KEY_EQUALS
	KeyBackspace     = C.ALLEGRO_KEY_BACKSPACE
	KeyTab           = C.ALLEGRO_KEY_TAB
	KeyOpenbrace     = C.ALLEGRO_KEY_OPENBRACE
	KeyClosebrace    = C.ALLEGRO_KEY_CLOSEBRACE
	KeyEnter         = C.ALLEGRO_KEY_ENTER
	KeySemicolon     = C.ALLEGRO_KEY_SEMICOLON
	KeyQuote         = C.ALLEGRO_KEY_QUOTE
	KeyBackslash     = C.ALLEGRO_KEY_BACKSLASH
	KeyBackslash2    = C.ALLEGRO_KEY_BACKSLASH2
	KeyComma         = C.ALLEGRO_KEY_COMMA
	KeyFullstop      = C.ALLEGRO_KEY_FULLSTOP
	KeySlash         = C.ALLEGRO_KEY_SLASH
	KeySpace         = C.ALLEGRO_KEY_SPACE
	KeyInsert        = C.ALLEGRO_KEY_INSERT
	KeyDelete        = C.ALLEGRO_KEY_DELETE
	KeyHome          = C.ALLEGRO_KEY_HOME
	KeyEnd           = C.ALLEGRO_KEY_END
	KeyPgup          = C.ALLEGRO_KEY_PGUP
	KeyPgdn          = C.ALLEGRO_KEY_PGDN
	KeyLeft          = C.ALLEGRO_KEY_LEFT
	KeyRight         = C.ALLEGRO_KEY_RIGHT
	KeyUp            = C.ALLEGRO_KEY_UP
	KeyDown          = C.ALLEGRO_KEY_DOWN
	KeyPadSlash      = C.ALLEGRO_KEY_PAD_SLASH
	KeyPadAsterisk   = C.ALLEGRO_KEY_PAD_ASTERISK
	KeyPadMinus      = C.ALLEGRO_KEY_PAD_MINUS
	KeyPadPlus       = C.ALLEGRO_KEY_PAD_PLUS
	KeyPadDelete     = C.ALLEGRO_KEY_PAD_DELETE
	KeyPadEnter      = C.ALLEGRO_KEY_PAD_ENTER
	KeyPrintscreen   = C.ALLEGRO_KEY_PRINTSCREEN
	KeyPause         = C.ALLEGRO_KEY_PAUSE
	KeyAbntC1        = C.ALLEGRO_KEY_ABNT_C1
	KeyYen           = C.ALLEGRO_KEY_YEN
	KeyKana          = C.ALLEGRO_KEY_KANA
	KeyConvert       = C.ALLEGRO_KEY_CONVERT
	KeyNoconvert     = C.ALLEGRO_KEY_NOCONVERT
	KeyAt            = C.ALLEGRO_KEY_AT
	KeyCircumflex    = C.ALLEGRO_KEY_CIRCUMFLEX
	KeyColon2        = C.ALLEGRO_KEY_COLON2
	KeyKanji         = C.ALLEGRO_KEY_KANJI
	KeyPadEquals     = C.ALLEGRO_KEY_PAD_EQUALS
	KeyBackquote     = C.ALLEGRO_KEY_BACKQUOTE
	KeySemicolon2    = C.ALLEGRO_KEY_SEMICOLON2
	KeyCommand       = C.ALLEGRO_KEY_COMMAND
	KeyLshift        = C.ALLEGRO_KEY_LSHIFT
	KeyRshift        = C.ALLEGRO_KEY_RSHIFT
	KeyLctrl         = C.ALLEGRO_KEY_LCTRL
	KeyRctrl         = C.ALLEGRO_KEY_RCTRL
	KeyAlt           = C.ALLEGRO_KEY_ALT
	KeyAltgr         = C.ALLEGRO_KEY_ALTGR
	KeyLwin          = C.ALLEGRO_KEY_LWIN
	KeyRwin          = C.ALLEGRO_KEY_RWIN
	KeyMenu          = C.ALLEGRO_KEY_MENU
	KeyScrolllock    = C.ALLEGRO_KEY_SCROLLLOCK
	KeyNumlock       = C.ALLEGRO_KEY_NUMLOCK
	KeyCapslock      = C.ALLEGRO_KEY_CAPSLOCK
	KeyMax           = C.ALLEGRO_KEY_MAX
	KeymodShift      = C.ALLEGRO_KEYMOD_SHIFT
	KeymodCtrl       = C.ALLEGRO_KEYMOD_CTRL
	KeymodAlt        = C.ALLEGRO_KEYMOD_ALT
	KeymodLwin       = C.ALLEGRO_KEYMOD_LWIN
	KeymodRwin       = C.ALLEGRO_KEYMOD_RWIN
	KeymodMenu       = C.ALLEGRO_KEYMOD_MENU
	KeymodAltgr      = C.ALLEGRO_KEYMOD_ALTGR
	KeymodCommand    = C.ALLEGRO_KEYMOD_COMMAND
	KeymodScrolllock = C.ALLEGRO_KEYMOD_SCROLLLOCK
	KeymodNumlock    = C.ALLEGRO_KEYMOD_NUMLOCK
	KeymodCapslock   = C.ALLEGRO_KEYMOD_CAPSLOCK
	KeymodInaltseq   = C.ALLEGRO_KEYMOD_INALTSEQ
	KeymodAccent1    = C.ALLEGRO_KEYMOD_ACCENT1
	KeymodAccent2    = C.ALLEGRO_KEYMOD_ACCENT2
	KeymodAccent3    = C.ALLEGRO_KEYMOD_ACCENT3
	KeymodAccent4    = C.ALLEGRO_KEYMOD_ACCENT4
)

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
