package allegro

// #cgo LDFLAGS: -lallegro
// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func Init() bool {
	return bool(C.al_install_system(C.ALLEGRO_VERSION_INT, nil))
}

func InstallSystem(version int32) bool {
	return bool(C.al_install_system(C.int(version), nil))
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

func GetStandardPath(id int32) *Path {
	return (*Path)(unsafe.Pointer(C.al_get_standard_path(C.int(id))))
}

func SetExeName(path string) {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	C.al_set_exe_name(p)
}

func GetAppName() string {
	return C.GoString(C.al_get_app_name())
}

func SetAppName(appName string) {
	a := C.CString(appName)
	defer C.free(unsafe.Pointer(a))
	C.al_set_app_name(a)
}

func GetOrgName() string {
	return C.GoString(C.al_get_org_name())
}

func SetOrgName(orgName string) {
	o := C.CString(orgName)
	defer C.free(unsafe.Pointer(o))
	C.al_set_org_name(o)
}

func GetSystemConfig() *Config {
	return (*Config)(unsafe.Pointer(C.al_get_system_config()))
}
