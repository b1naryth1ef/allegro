package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

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

//func EmitUserEvent()

func (u *UserEvent) Unref() {
	C.al_unref_user_event((*C.ALLEGRO_USER_EVENT)(unsafe.Pointer(u)))
}

func (e *EventSource) GetData() unsafe.Pointer {
	return unsafe.Pointer(C.al_get_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(e))))
}

func (e *EventSource) SetData(data unsafe.Pointer) {
	C.al_set_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(unsafe.Pointer(e)), C.intptr_t(data))
}
