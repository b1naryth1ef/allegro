package allegro

// #cgo pkg-config: allegro-5.0
// #include <allegro5/allegro.h>
import "C"

import (
	"errors"
	"unsafe"
)

func Init() error {
	b := bool(C.al_install_system(C.ALLEGRO_VERSION_INT, nil))
	if b == false {
		return errors.New("Allegro can't be initialized")
	}
	return nil
}

func UninstallSystem() {
	C.al_uninstall_system()
}

func IsSystemInstalled() bool {
	return bool(C.al_is_system_installed())
}

func GetAllegroVersion() uint32 {
	return uint32(C.al_get_allegro_version())
}

//func SetExeName(string path) {
//	p := C.CString(path)
//	defer C.free(unsafe.pointer(p))
//	C.al_set_exe_name(p)
//}

func SetAppName(appName string) {
	a := C.CString(appName)
	defer C.free(unsafe.Pointer(a))
	C.al_set_app_name(a)
}

func SetOrgName(orgName string) {
	o := C.CString(orgName)
	defer C.free(unsafe.Pointer(o))
	C.al_set_org_name(o)
}

func GetAppName() string {
	return C.GoString(C.al_get_app_name())
}

func GetOrgName() string {
	return C.GoString(C.al_get_org_name())
}

/*
	SKIPPED:
	
	al_install_system
	al_get_standard_path
	al_get_system_config
	al_register_assert_handler
*/
