package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func GetTime() float64 {
	return float64(C.al_get_time())
}

func CurrentTime() float64 {
	return float64(C.al_get_time())
}

func InitTimeout(seconds float64) *Timeout {
	var t C.ALLEGRO_TIMEOUT
	C.al_init_timeout(&t, C.double(seconds))
	return (*Timeout)(unsafe.Pointer(&t))
}

func Rest(seconds float64) {
	C.al_rest(C.double(seconds))
}
