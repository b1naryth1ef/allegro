package physicsfs

// #cgo LDFLAGS: -lallegro_physfs
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_physfs.h>
import "C"

func SetFileInterface() {
	C.al_set_physfs_file_interface()
}

func GetVersion() uint32 {
	return uint32(C.al_get_allegro_physfs_version())
}
