package allegro

// #include <allegro5/allegro.h>
// #include "events.h"
import "C"

import "unsafe"

// config.go
type Config C.ALLEGRO_CONFIG

type ConfigEntry C.ALLEGRO_CONFIG_ENTRY

type ConfigSection C.ALLEGRO_CONFIG_SECTION

// display.go
type Display C.ALLEGRO_DISPLAY

// events.go
type EventSource C.ALLEGRO_EVENT_SOURCE

type EventQueue C.ALLEGRO_EVENT_QUEUE

type Event struct {
	Type      uint32
	Any       AnyEvent
	DisplayE  DisplayEvent
	JoystickE JoystickEvent
	KeyboardE KeyboardEvent
	MouseE    MouseEvent
	TimerE    TimerEvent
	TouchE    TouchEvent
	UserE     UserEvent
}

type AnyEvent struct {
	Type      uint32
	Source    *EventSource
	TimeStamp float64
}

type DisplayEvent struct {
	Type          uint32
	Source        *Display
	TimeStamp     float64
	X, Y          int32
	Width, Height int32
	Orientation   int32
}

type JoystickEvent struct {
	Type      uint32
	Source    *Joystick
	TimeStamp float64
	Id        *Joystick
	Stick     int32
	Axis      int32
	Pos       float32
	Button    int32
}

type KeyboardEvent struct {
	Type      uint32
	Source    *Keyboard
	TimeStamp float64
	Disp      *Display
	Keycode   int32
	UniChar   int32
	Modifiers uint32
	Repeat    bool
}

type MouseEvent struct {
	Type           uint32
	Source         *Mouse
	TimeStamp      float64
	X, Y, Z, W     int32
	Dx, Dy, Dz, Dw int32
	Button         uint32
	Pressure       float32
}

type TimerEvent struct {
	Type      uint32
	Source    *Timer
	TimeStamp float64
	Count     int64
	Error     float64
}

type TouchEvent struct {
	Type      uint32
	Source    *TouchInput
	TimeStamp float64
	Disp      *Display
	Id        int32
	X, Y      float32
	Dx, Dy    float32
	Primary   bool
}

type UserEvent struct {
	Type          uint32
	Source        *EventSource
	TimeStamp     float64
	internalDescr *C.ALLEGRO_USER_EVENT_DESCRIPTOR
	Data1         unsafe.Pointer
	Data2         unsafe.Pointer
	Data3         unsafe.Pointer
	Data4         unsafe.Pointer
}

// file.go
type File C.ALLEGRO_FILE

// filesystem.go
type FsEntry C.ALLEGRO_FS_ENTRY

// fullscreen.go
type DisplayMode struct {
	Width       int32
	Height      int32
	Format      int32
	RefreshRate int32
}

// graphics.go
type Color C.ALLEGRO_COLOR

type Bitmap C.ALLEGRO_BITMAP

type LockedRegion struct {
	Data      unsafe.Pointer
	Format    int32
	Pitch     int32
	PixelSize int32
}

// joystick.go
type Joystick C.ALLEGRO_JOYSTICK

type JoystickState C.ALLEGRO_JOYSTICK_STATE

// keyboard.go
type Keyboard C.ALLEGRO_KEYBOARD

type KeyboardState C.ALLEGRO_KEYBOARD_STATE

// monitor.go
type MonitorInfo struct {
	X1 int32
	Y1 int32
	X2 int32
	Y2 int32
}

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
