package allegro

// #include <allegro5/allegro.h>
import "C"

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

type Event struct {
}

type UserEvent struct {
	allegroUserEvent *C.ALLEGRO_USER_EVENT
}

type EventQueue struct {
	allegroEventQueue *C.ALLEGRO_EVENT_QUEUE
}

type EventSource struct {
	allegroEventSource *C.ALLEGRO_EVENT_SOURCE
}

func CreateEventQueue() *EventQueue {
	return &EventQueue{C.al_create_event_queue()}
}

func (e *EventQueue) Destroy() {
	C.al_destroy_event_queue(e.allegroEventQueue)
}

func (e *EventQueue) RegisterEventSource(source *EventSource) {
	C.al_register_event_source(e.allegroEventQueue, source.allegroEventSource)
}

func (e *EventQueue) UnregisterEventSource(source *EventSource) {
	C.al_unregister_event_source(e.allegroEventQueue, source.allegroEventSource)
}

func (e *EventQueue) IsEmpty() bool {
	return bool(C.al_is_event_queue_empty(e.allegroEventQueue))
}

func (e *EventQueue) GetNextEvent() bool, *Event {
	// TODO
}

func (e *EventQueue) PeekNextEvent() bool, *Event {
	// TODO
}

func (e *EventQueue) DropNextEvent() bool {
	return bool(C.al_drop_next_event(e.allegroEventQueue))
}


