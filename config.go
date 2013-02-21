package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

type Config struct {
	allegroConfig *C.ALLEGRO_CONFIG
}

type ConfigSection struct {
	allegroConfigSection **C.ALLEGRO_CONFIG_SECTION
}

type ConfigEntry struct {
	allegroConfigEntry **C.ALLEGRO_CONFIG_ENTRY
}

func CreateConfig() *Config {
	return &Config{C.al_create_config()}
}

func (c *Config) Destroy() {
	C.al_destroy_config(c.allegroConfig)
}

func LoadConfigFile(fileName string) *Config {
	f := C.CString(fileName)
	defer C.free(unsafe.Pointer(f))
	return &Config{C.al_load_config_file(f)}
}

func LoadConfigFileF(f *File) *Config {
	return &Config{C.al_load_config_file_f(f.allegroFile)}
}

func (c *Config) SaveFile(fileName string) bool {
	f := C.CString(fileName)
	defer C.free(unsafe.Pointer(f))
	return bool(C.al_save_config_file(f, c.allegroConfig))
}

func (c *Config) SaveFileF(f *File) bool {
	return bool(C.al_save_config_file_f(f.allegroFile, c.allegroConfig))
}

func (c *Config) AddSection(name string) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.al_add_config_section(c.allegroConfig, n)
}

func (c *Config) AddComment(section, comment string) {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	co := C.CString(comment)
	defer C.free(unsafe.Pointer(co))
	C.al_add_config_comment(c.allegroConfig, s, co)
}

func (c *Config) GetValue(section, key string) string {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	return C.GoString(C.al_get_config_value(c.allegroConfig, s, k))
}

func (c *Config) SetValue(section, key, value string) {
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	v := C.CString(value)
	defer C.free(unsafe.Pointer(v))
	C.al_set_config_value(c.allegroConfig, s, k, v)
}

func (c *Config) GetFirstSection() (string, *ConfigSection) {
	cs := new((*C.ALLEGRO_CONFIG_SECTION))
	s := C.al_get_first_config_section(c.allegroConfig, cs)
	return C.GoString(s), &ConfigSection{cs}
}

func GetNextSection(iterator *ConfigSection) string {
	return C.GoString(C.al_get_next_config_section(iterator.allegroConfigSection))
}

func (c *Config) GetFirstEntry(section string) (string, *ConfigEntry) {
	ce := new((*C.ALLEGRO_CONFIG_ENTRY))
	s := C.CString(section)
	defer C.free(unsafe.Pointer(s))
	r := C.al_get_first_config_entry(c.allegroConfig, s, ce)
	return C.GoString(r), &ConfigEntry{ce}
}

func GetNextEntry(iterator *ConfigEntry) string {
	return C.GoString(C.al_get_next_config_entry(iterator.allegroConfigEntry))
}

func (c *Config) Merge(cfg *Config) *Config {
	return &Config{C.al_merge_config(c.allegroConfig, cfg.allegroConfig)}
}

func (c *Config) MergeInto(add *Config) {
	C.al_merge_config_into(c.allegroConfig, add.allegroConfig)
}
