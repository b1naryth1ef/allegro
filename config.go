package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func CreateConfig() *Config {
	return (*Config)(unsafe.Pointer(C.al_create_config()))
}

func (c *Config) Destroy() {
	C.al_destroy_config((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)))
}

func LoadConfigFile(filename string) *Config {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return (*Config)(unsafe.Pointer(C.al_load_config_file(f)))
}

func LoadConfigFileF(file *File) *Config {
	return (*Config)(unsafe.Pointer(C.al_load_config_file_f((*C.ALLEGRO_FILE)(unsafe.Pointer(file)))))
}

func (c *Config) SaveFile(filename string) bool {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return bool(C.al_save_config_file(f, (*C.ALLEGRO_CONFIG)(unsafe.Pointer(c))))
}

func (c *Config) SaveFileF(file *File) bool {
	return bool(C.al_save_config_file_f((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), (*C.ALLEGRO_FILE)(unsafe.Pointer(file))))
}

func (c *Config) AddSection(name string) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.al_add_config_section((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), n)
}

func (c *Config) RemoveSection(section string) bool {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	return bool(C.al_remove_config_section((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), s))
}

func (c *Config) AddComment(section, comment string) {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	co := C.CString(comment)
	defer C.free(unsafe.Pointer(co))
	C.al_add_config_comment((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), s, co)
}

func (c *Config) GetValue(section, key string) string {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	return C.GoString(C.al_get_config_value((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), s, k))
}

func (c *Config) SetValue(section, key, value string) {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	v := C.CString(value)
	defer C.free(unsafe.Pointer(v))
	C.al_set_config_value((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), s, k, v)
}

func (c *Config) RemoveKey(section, key string) bool {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	return bool(C.al_remove_config_key((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), s, k))
}

func (c *Config) GetFirstSection() (string, **ConfigSection) {
	cs := new((*C.ALLEGRO_CONFIG_SECTION))
	s := C.al_get_first_config_section((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), cs)
	return C.GoString(s), (**ConfigSection)(unsafe.Pointer(cs))
}

func GetNextSection(iterator **ConfigSection) string {
	return C.GoString(C.al_get_next_config_section((**C.ALLEGRO_CONFIG_SECTION)(unsafe.Pointer(iterator))))
}

func (c *Config) GetFirstEntry(section string) (string, **ConfigEntry) {
	ce := new((*C.ALLEGRO_CONFIG_ENTRY))
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	r := C.al_get_first_config_entry((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), s, ce)
	return C.GoString(r), (**ConfigEntry)(unsafe.Pointer(ce))
}

func GetNextEntry(iterator **ConfigEntry) string {
	return C.GoString(C.al_get_next_config_entry((**C.ALLEGRO_CONFIG_ENTRY)(unsafe.Pointer(iterator))))
}

func (c *Config) Merge(cfg *Config) *Config {
	return (*Config)(unsafe.Pointer(C.al_merge_config((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), (*C.ALLEGRO_CONFIG)(unsafe.Pointer(cfg)))))
}

func (c *Config) MergeInto(add *Config) {
	C.al_merge_config_into((*C.ALLEGRO_CONFIG)(unsafe.Pointer(c)), (*C.ALLEGRO_CONFIG)(unsafe.Pointer(add)))
}
