package allegro

// #include <allegro5/allegro.h>
import "C"

type Timer struct {
	allegroTimer *C.ALLEGRO_TIMER
}

func CreateTimer(speedSecs float64) *Timer {
	return &Timer{C.al_create_timer(C.double(speedSecs))}
}

func (t *Timer) StartTimer() {
	C.al_start_timer(t.allegroTimer)
}

func (t *Timer) StopTimer() {
	C.al_stop_timer(t.allegroTimer)
}

func (t *Timer) GetTimerStarted() bool {
	return bool(C.al_get_timer_started(t.allegroTimer))
}

func (t *Timer) DestroyTimer() {
	C.al_destroy_timer(t.allegroTimer)
}

func (t *Timer) GetTimerCount() int64 {
	return int64(C.al_get_timer_count(t.allegroTimer))
}

func (t *Timer) SetTimerCount(newCount int64) {
	C.al_set_timer_count(t.allegroTimer, C.int64_t(newCount))
}

func (t *Timer) AddTimerCount(diff int64) {
	C.al_add_timer_count(t.allegroTimer, C.int64_t(diff))
}

func (t *Timer) GetTimerSpeed() float64 {
	return float64(C.al_get_timer_speed(t.allegroTimer))
}

func (t *Timer) SetTimerSpeed(newSpeedSecs float64) {
	C.al_set_timer_speed(t.allegroTimer, C.double(newSpeedSecs))
}
