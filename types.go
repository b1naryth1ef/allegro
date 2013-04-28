package allegro

// #include <allegro5/allegro.h>
import "C"

// config.go
type Config C.ALLEGRO_CONFIG
type ConfigEntry C.ALLEGRO_CONFIG_ENTRY
type ConfigSection C.ALLEGRO_CONFIG_SECTION

// display.go
type Display C.ALLEGRO_DISPLAY

// events.go
type EventSource C.ALLEGRO_EVENT_SOURCE
type EventQueue C.ALLEGRO_EVENT_QUEUE
type Event C.ALLEGRO_EVENT

// file.go
type File C.ALLEGRO_FILE

// filesystem.go
type FsEntry C.ALLEGRO_FS_ENTRY

// fullscreen.go
type DisplayMode C.ALLEGRO_DISPLAY_MODE

// graphics.go
type Color C.ALLEGRO_COLOR

type Bitmap C.ALLEGRO_BITMAP

type LockedRegion C.ALLEGRO_LOCKED_REGION

// joystick.go
type Joystick C.ALLEGRO_JOYSTICK

type JoystickState C.ALLEGRO_JOYSTICK_STATE

// keyboard.go
type Keyboard C.ALLEGRO_KEYBOARD

type KeyboardState C.ALLEGRO_KEYBOARD_STATE

// monitor.go
type MonitorInfo C.ALLEGRO_MONITOR_INFO

// mouse.go
type Mouse C.ALLEGRO_MOUSE

type MouseCursor C.ALLEGRO_MOUSE_CURSOR

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

// shader.go
type Shader C.ALLEGRO_SHADER

// state.go
type State C.ALLEGRO_STATE

// timer.go
type Timer C.ALLEGRO_TIMER

// touchinput.go
type TouchInput C.ALLEGRO_TOUCH_INPUT

type TouchState C.ALLEGRO_TOUCH_STATE

type TouchInputState C.ALLEGRO_TOUCH_INPUT_STATE

//transformation.go
type Transform struct {
	M [4][4]float32
}
