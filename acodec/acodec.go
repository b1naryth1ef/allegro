package acodec

// #cgo LDFLAGS: -lallegro_acodec
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_acodec.h>
import "C"

func InitAcodecAddon() bool {
	return bool(C.al_init_acodec_addon())
}

func GetAllegroCodecVersion() uint32 {
	return uint32(C.al_get_allegro_acodec_version())
}
