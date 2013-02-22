package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

type Timer C.ALLEGRO_TIMER

func CreateTimer(speedSecs float64) *Timer {
	return (*Timer)(unsafe.Pointer(C.al_create_timer(C.double(speedSecs))))
}

func (t *Timer) Destroy() {
	C.al_destroy_timer((*C.ALLEGRO_TIMER)(unsafe.Pointer(t)))
}

func (t *Timer) Start() {
	C.al_start_timer((*C.ALLEGRO_TIMER)(unsafe.Pointer(t)))
}

func (t *Timer) Stop() {
	C.al_stop_timer((*C.ALLEGRO_TIMER)(unsafe.Pointer(t)))
}

func (t *Timer) IsStarted() bool {
	return bool(C.al_get_timer_started((*C.ALLEGRO_TIMER)(unsafe.Pointer(t))))
}

func (t *Timer) GetCount() int64 {
	return int64(C.al_get_timer_count((*C.ALLEGRO_TIMER)(unsafe.Pointer(t))))
}

func (t *Timer) SetCount(newCount int64) {
	C.al_set_timer_count((*C.ALLEGRO_TIMER)(unsafe.Pointer(t)), C.int64_t(newCount))
}

func (t *Timer) AddCount(diff int64) {
	C.al_add_timer_count((*C.ALLEGRO_TIMER)(unsafe.Pointer(t)), C.int64_t(diff))
}

func (t *Timer) GetSpeed() float64 {
	return float64(C.al_get_timer_speed((*C.ALLEGRO_TIMER)(unsafe.Pointer(t))))
}

func (t *Timer) SetSpeed(newSpeedSecs float64) {
	C.al_set_timer_speed((*C.ALLEGRO_TIMER)(unsafe.Pointer(t)), C.double(newSpeedSecs))
}

func (t *Timer) GetEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_timer_event_source((*C.ALLEGRO_TIMER)(unsafe.Pointer(t)))))
}
