package allegro

// #include <allegro5/allegro.h>
// #include "events.h"
import "C"

import "unsafe"

const (
	EventJoystickAxis          = C.ALLEGRO_EVENT_JOYSTICK_AXIS
	EventJoystickButtonDown    = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN
	EventJoystickButtonUp      = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP
	EventJoystickConfiguration = C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION
	EventKeyDown               = C.ALLEGRO_EVENT_KEY_DOWN
	EventKeyChar               = C.ALLEGRO_EVENT_KEY_CHAR
	EventKeyUp                 = C.ALLEGRO_EVENT_KEY_UP
	EventMouseAxis             = C.ALLEGRO_EVENT_MOUSE_AXES
	EventMouseButtonDown       = C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN
	EventMouseButtonUp         = C.ALLEGRO_EVENT_MOUSE_BUTTON_UP
	EventMouseEnterDisplay     = C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY
	EventMouseLeaveDisplay     = C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY
	EventMouseWarped           = C.ALLEGRO_EVENT_MOUSE_WARPED
	EventTimer                 = C.ALLEGRO_EVENT_TIMER
	EventDisplayExpose         = C.ALLEGRO_EVENT_DISPLAY_EXPOSE
	EventDisplayResize         = C.ALLEGRO_EVENT_DISPLAY_RESIZE
	EventDisplayClose          = C.ALLEGRO_EVENT_DISPLAY_CLOSE
	EventDisplayLost           = C.ALLEGRO_EVENT_DISPLAY_LOST
	EventDisplayFound          = C.ALLEGRO_EVENT_DISPLAY_FOUND
	EventDisplaySwitchIn       = C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN
	EventDisplaySwitchOut      = C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT
	EventDisplayOrientation    = C.ALLEGRO_EVENT_DISPLAY_ORIENTATION
)

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
	Stick     int32
	Axis      int32
	Pos       float32
	Button    int32
}

type KeyboardEvent struct {
	Type      uint32
	Source    *Keyboard
	TimeStamp float64
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

type UserEvent struct {
	Type          uint32
	Source        *EventSource
	TimeStamp     float64
	internalDescr *C.ALLEGRO_USER_EVENT_DESCRIPTOR
	Data1         uintptr
	Data2         uintptr
	Data3         uintptr
	Data4         uintptr
}

// TODO fix the var names (they are types)
type Event struct {
	Type     uint32
	Any      AnyEvent
	Display  DisplayEvent
	Joystick JoystickEvent
	Keyboard KeyboardEvent
	Mouse    MouseEvent
	Timer    TimerEvent
	User     UserEvent
}

type EventSource C.ALLEGRO_EVENT_SOURCE

type EventQueue C.ALLEGRO_EVENT_QUEUE

func CreateEventQueue() *EventQueue {
	return (*EventQueue)(unsafe.Pointer(C.al_create_event_queue()))
}

func (e *EventQueue) Destroy() {
	C.al_destroy_event_queue((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)))
}

func (e *EventQueue) RegisterEventSource(source *EventSource) {
	C.al_register_event_source((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), (*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(source)))
}

func (e *EventQueue) UnregisterEventSource(source *EventSource) {
	C.al_unregister_event_source((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), (*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(source)))
}

func (e *EventQueue) IsEmpty() bool {
	return bool(C.al_is_event_queue_empty((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e))))
}

func (e *EventQueue) GetNextEvent() (bool, *Event) {
	ev := new(Event)
	r := bool(C.al_get_next_event((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), (*C.ALLEGRO_EVENT)(unsafe.Pointer(ev))))
	return r, ev
}

func (e *EventQueue) PeekNextEvent() (bool, *Event) {
	ev := new(Event)
	r := bool(C.al_peek_next_event((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), (*C.ALLEGRO_EVENT)(unsafe.Pointer(ev))))
	return r, ev
}

func (e *EventQueue) DropNextEvent() bool {
	return bool(C.al_drop_next_event((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e))))
}

//func (e *EventQueue) Flush() bool {
//	return bool(C.al_flush_event_queue((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e))))
//}

func (e *EventQueue) WaitForEvent() *Event {
	ev := new(Event)
	C.al_wait_for_event((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), (*C.ALLEGRO_EVENT)(unsafe.Pointer(ev)))
	return ev
}

func (e *EventQueue) WaitForEventTimed(secs float32) (bool, *Event) {
	ev := new(Event)
	r := bool(C.al_wait_for_event_timed((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), (*C.ALLEGRO_EVENT)(unsafe.Pointer(ev)), C.float(secs)))
	return r, ev
}

func (e *EventQueue) WaitForEventUntil(timeout float64) (bool, *Event) {
	var t C.ALLEGRO_TIMEOUT
	C.al_init_timeout(&t, C.double(timeout))
	ev := new(Event)
	r := bool(C.al_wait_for_event_until((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), (*C.ALLEGRO_EVENT)(unsafe.Pointer(ev)), &t))
	return r, ev
}

func InitUserEventSource() *EventSource {
	es := new(C.ALLEGRO_EVENT_SOURCE)
	C.al_init_user_event_source(es)
	return (*EventSource)(unsafe.Pointer(es))
}

func (e *EventSource) Destroy() {
	C.al_destroy_user_event_source((*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(e)))
}

func (e *EventSource) GetData() uintptr {
	return uintptr(C.al_get_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(e))))
}

func (e *EventSource) SetData(data uintptr) {
	C.al_set_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(e)), C.intptr_t(data))
}
