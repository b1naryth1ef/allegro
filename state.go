package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

const (
	StateNewDisplayParameters = C.ALLEGRO_STATE_NEW_DISPLAY_PARAMETERS
	StateNewBitmapParameters  = C.ALLEGRO_STATE_NEW_BITMAP_PARAMETERS
	StateDisplay              = C.ALLEGRO_STATE_DISPLAY
	StateBlender              = C.ALLEGRO_STATE_BLENDER
	StteTransform             = C.ALLEGRO_STATE_TRANSFORM
	StateNewFileInterface     = C.ALLEGRO_STATE_NEW_FILE_INTERFACE
	StateBitmap               = C.ALLEGRO_STATE_BITMAP
	StateTargetBitmap         = C.ALLEGRO_STATE_TARGET_BITMAP
	StateAll                  = C.ALLEGRO_STATE_ALL
)

type State C.ALLEGRO_STATE

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
