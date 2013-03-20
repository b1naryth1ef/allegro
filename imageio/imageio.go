package imageio

// #cgo LDFLAGS: -lallegro_image
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_image.h>
import "C"

func Init() bool {
	return bool(C.al_init_image_addon())
}

func Shutdown() {
	C.al_shutdown_image_addon()
}

func GetVersion() uint32 {
	return uint32(C.al_get_allegro_image_version())
}
