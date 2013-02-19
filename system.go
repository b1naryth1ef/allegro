package allegro

// #cgo pkg-config: allegro-5.0
// #include <allegro5/allegro.h>
import "C"

import "unsafe"

const (
	ResourcesPath     = C.ALLEGRO_RESOURCES_PATH
	TempPath          = C.ALLEGRO_TEMP_PATH
	UserDataPath      = C.ALLEGRO_USER_DATA_PATH
	UserHomePath      = C.ALLEGRO_USER_HOME_PATH
	UserSettingsPath  = C.ALLEGRO_USER_SETTINGS_PATH
	UserDocumentsPath = C.ALLEGRO_USER_DOCUMENTS_PATH
	ExenamePath       = C.ALLEGRO_EXENAME_PATH
)

func Init() bool {
	return bool(C.al_install_system(C.ALLEGRO_VERSION_INT, nil))
}

func InstallSystem(version int) bool {
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

func GetStandardPath(id int) *Path {
	return &Path{C.al_get_standard_path(C.int(id))}
}

//func SetExeName(string path) {
//	p := C.CString(path)
//	defer C.free(unsafe.pointer(p))
//	C.al_set_exe_name(p)
//}

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
	return &Config{C.al_get_system_config()}
}
