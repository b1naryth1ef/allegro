package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

func CreateShader(platform uint32) *Shader {
	return (*Shader)(unsafe.Pointer(C.al_create_shader(C.ALLEGRO_SHADER_PLATFORM(platform))))
}

func (s *Shader) Destroy() {
	C.al_destroy_shader((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)))
}

func (s *Shader) AttachSource(t uint32, source string) bool {
	so := C.CString(source)
	defer C.free(unsafe.Pointer(so))
	return bool(C.al_attach_shader_source((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), C.ALLEGRO_SHADER_TYPE(t), so))
}

func (s *Shader) AttachSourceFile(t uint32, filename string) bool {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))
	return bool(C.al_attach_shader_source_file((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), C.ALLEGRO_SHADER_TYPE(t), f))
}

func (s *Shader) Build() bool {
	return bool(C.al_build_shader((*C.ALLEGRO_SHADER)(unsafe.Pointer(s))))
}

func (s *Shader) GetLog() string {
	return C.GoString(C.al_get_shader_log((*C.ALLEGRO_SHADER)(unsafe.Pointer(s))))
}

func (s *Shader) GetPlatform() uint32 {
	return uint32(C.al_get_shader_platform((*C.ALLEGRO_SHADER)(unsafe.Pointer(s))))
}

func (s *Shader) Use() {
	C.al_use_shader((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)))
}

func SetShaderSampler(name string, bmp *Bitmap, unit int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_sampler(n, (*C.ALLEGRO_BITMAP)(unsafe.Pointer(bmp)), C.int(unit)))
}

func SetShaderMatrix(name string, matrix *Transform) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_matrix(n, (*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(matrix))))
}

func SetShaderInt(name string, i int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_int(n, C.int(i)))
}

func SetShaderFloat(name string, f float32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_float(n, C.float(f)))
}

func SetShaderBool(name string, b bool) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_bool(n, C.bool(b)))
}

func SetShaderIntVector(name string, elemSize int32, i *int32, numElems int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_int_vector(n, C.int(elemSize), (*C.int)(unsafe.Pointer(i)), C.int(numElems)))
}

func SetShaderFloatVector(name string, elemSize int32, f *float32, numElems int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_float_vector(n, C.int(elemSize), (*C.float)(unsafe.Pointer(f)), C.int(numElems)))
}

func GetDefaultShaderSource(platform, t uint32) string {
	return C.GoString(C.al_get_default_shader_source(C.ALLEGRO_SHADER_PLATFORM(platform), C.ALLEGRO_SHADER_TYPE(t)))
}
