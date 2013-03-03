package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func StoreAllegroState(flags int32) *State {
	s := new(C.ALLEGRO_STATE)
	C.al_store_state(s, C.int(flags))
	return (*State)(unsafe.Pointer(s))
}

func (s *State) Restore() {
	C.al_restore_state((*C.ALLEGRO_STATE)(unsafe.Pointer(s)))
}

func GetErrno() int32 {
	return int32(C.al_get_errno())
}

func SetErrno(errnum int32) {
	C.al_set_errno(C.int(errnum))
}
