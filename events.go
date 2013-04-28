package allegro

// #include <allegro5/allegro.h>
// #include "events.h"
import "C"

import "unsafe"

/***********************************/
/* Getters for the Event structure */
/***********************************/

// Event

func (e *Event) Type() uint32 {
	return uint32(C.GetEventType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// Any

func (e *Event) AnyType() uint32 {
	return uint32(C.GetAnyType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) AnySource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.GetAnySource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) AnyTimestamp() float64 {
	return float64(C.GetAnyTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// Display

func (e *Event) DisplayType() uint32 {
	return uint32(C.GetDisplayType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) DisplaySource() *Display {
	return (*Display)(unsafe.Pointer(C.GetDisplaySource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) DisplayTimestamp() float64 {
	return float64(C.GetDisplayTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) DisplayX() int32 {
	return int32(C.GetDisplayX((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) DisplayY() int32 {
	return int32(C.GetDisplayY((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) DisplayWidth() int32 {
	return int32(C.GetDisplayWidth((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) DisplayHeight() int32 {
	return int32(C.GetDisplayHeight((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) DisplayOrientation() int32 {
	return int32(C.GetDisplayOrientation((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// Joystick

func (e *Event) JoystickType() uint32 {
	return uint32(C.GetJoystickType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) JoystickSource() *Joystick {
	return (*Joystick)(unsafe.Pointer(C.GetJoystickSource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) JoystickTimestamp() float64 {
	return float64(C.GetJoystickTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) JoystickId() *Joystick {
	return (*Joystick)(unsafe.Pointer(C.GetJoystickId((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) JoystickStick() int32 {
	return int32(C.GetJoystickStick((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) JoystickAxis() int32 {
	return int32(C.GetJoystickAxis((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) JoystickPos() float32 {
	return float32(C.GetJoystickPos((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) JoystickButton() int32 {
	return int32(C.GetJoystickButton((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// Keyboard

func (e *Event) KeyboardType() uint32 {
	return uint32(C.GetKeyboardType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) KeyboardSource() *Keyboard {
	return (*Keyboard)(unsafe.Pointer(C.GetKeyboardSource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) KeyboardTimestamp() float64 {
	return float64(C.GetKeyboardTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) KeyboardDisplay() *Display {
	return (*Display)(unsafe.Pointer(C.GetKeyboardDisplay((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) KeyboardKeycode() int32 {
	return int32(C.GetKeyboardKeycode((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) KeyboardUnichar() int32 {
	return int32(C.GetKeyboardUnichar((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) KeyboardModifiers() uint32 {
	return uint32(C.GetKeyboardModifiers((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) KeyboardRepeat() bool {
	return bool(C.GetKeyboardRepeat((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// Mouse

func (e *Event) MouseType() uint32 {
	return uint32(C.GetMouseType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseSource() *Mouse {
	return (*Mouse)(unsafe.Pointer(C.GetMouseSource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) MouseTimestamp() float64 {
	return float64(C.GetMouseTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseDisplay() *Display {
	return (*Display)(unsafe.Pointer(C.GetMouseDisplay((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) MouseX() int32 {
	return int32(C.GetMouseX((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseY() int32 {
	return int32(C.GetMouseY((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseZ() int32 {
	return int32(C.GetMouseZ((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseW() int32 {
	return int32(C.GetMouseW((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseDx() int32 {
	return int32(C.GetMouseDx((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseDy() int32 {
	return int32(C.GetMouseDy((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseDz() int32 {
	return int32(C.GetMouseDz((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseDw() int32 {
	return int32(C.GetMouseDw((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MouseButton() uint32 {
	return uint32(C.GetMouseButton((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) MousePressure() float32 {
	return float32(C.GetMousePressure((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// Timer

func (e *Event) TimerType() uint32 {
	return uint32(C.GetTimerType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TimerSource() *Timer {
	return (*Timer)(unsafe.Pointer(C.GetTimerSource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) TimerTimestamp() float64 {
	return float64(C.GetTimerTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TimerCount() int64 {
	return int64(C.GetTimerCount((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TimerError() float64 {
	return float64(C.GetTimerError((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// Touch

func (e *Event) TouchType() uint32 {
	return uint32(C.GetTouchType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TouchSource() *TouchInput {
	return (*TouchInput)(unsafe.Pointer(C.GetTouchSource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) TouchTimestamp() float64 {
	return float64(C.GetTouchTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TouchDisplay() *Display {
	return (*Display)(unsafe.Pointer(C.GetTouchDisplay((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) TouchId() int32 {
	return int32(C.GetTouchId((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TouchX() float32 {
	return float32(C.GetTouchX((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TouchY() float32 {
	return float32(C.GetTouchY((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TouchDx() float32 {
	return float32(C.GetTouchDx((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TouchDy() float32 {
	return float32(C.GetTouchDy((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) TouchPrimary() bool {
	return bool(C.GetTouchPrimary((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

// UserEvent

func (e *Event) UserType() uint32 {
	return uint32(C.GetUserType((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) UserSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.GetUserSource((*C.ALLEGRO_EVENT)(unsafe.Pointer(e)))))
}

func (e *Event) UserTimestamp() float64 {
	return float64(C.GetUserTimestamp((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) UserData1() uintptr {
	return uintptr(C.GetUserData1((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) UserData2() uintptr {
	return uintptr(C.GetUserData2((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) UserData3() uintptr {
	return uintptr(C.GetUserData3((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

func (e *Event) UserData4() uintptr {
	return uintptr(C.GetUserData4((*C.ALLEGRO_EVENT)(unsafe.Pointer(e))))
}

/****************/
/* Event module */
/****************/

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

func (e *EventQueue) Pause(pause bool) {
	C.al_pause_event_queue((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)), C.bool(pause))
}

func (e *EventQueue) IsPaused() bool {
	return bool(C.al_is_event_queue_paused((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e))))
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

func (e *EventQueue) Flush() {
	C.al_flush_event_queue((*C.ALLEGRO_EVENT_QUEUE)(unsafe.Pointer(e)))
}

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

// TODO: Implement emitting user event
//func EmitUserEvent() bool {}
//func (u *UserEvent) Unref() {}

func (e *EventSource) GetData() uintptr {
	return uintptr(C.al_get_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(e))))
}

func (e *EventSource) SetData(data uintptr) {
	C.al_set_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(e)), C.intptr_t(data))
}
