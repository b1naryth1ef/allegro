package allegro

// #cgo LDFLAGS: -lallegro_shader
// #include <allegro5/allegro.h>
// #include <allegro5/allegro_shader.h>
// #include <allegro5/allegro_shader_glsl.h>
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

func (s *Shader) Link() bool {
	return bool(C.al_link_shader((*C.ALLEGRO_SHADER)(unsafe.Pointer(s))))
}

func (s *Shader) GetLog() string {
	return C.GoString(C.al_get_shader_log((*C.ALLEGRO_SHADER)(unsafe.Pointer(s))))
}

//func (s *Shader) GetPlatform() uint32 {
//	return uint32(C.al_get_shader_platform((*C.ALLEGRO_SHADER)(unsafe.Pointer(s))))
//}

func (s *Shader) Use(use bool) {
	C.al_use_shader((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), C.bool(use))
}

func (s *Shader) SetSampler(name string, bmp *Bitmap, unit int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_sampler((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), n, (*C.ALLEGRO_BITMAP)(unsafe.Pointer(bmp)), C.int(unit)))
}

func (s *Shader) SetMatrix(name string, matrix *Transform) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_matrix((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), n, (*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(matrix))))
}

func (s *Shader) SetInt(name string, i int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_int((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), n, C.int(i)))
}

func (s *Shader) SetFloat(name string, f float32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_float((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), n, C.float(f)))
}

func (s *Shader) SetIntVector(name string, elemSize int32, i *int32, numElems int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_int_vector((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), n, C.int(elemSize), (*C.int)(unsafe.Pointer(i)), C.int(numElems)))
}

func (s *Shader) SetFloatVector(name string, elemSize int32, f *float32, numElems int32) bool {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return bool(C.al_set_shader_float_vector((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), n, C.int(elemSize), (*C.float)(unsafe.Pointer(f)), C.int(numElems)))
}

func (s *Shader) SetVertexArray(v *float32, stride int32) bool {
	return bool(C.al_set_shader_vertex_array((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), (*C.float)(unsafe.Pointer(v)), C.int(stride)))
}

func (s *Shader) SetColorArray(c *uint8, stride int32) bool {
	return bool(C.al_set_shader_color_array((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), (*C.uchar)(unsafe.Pointer(c)), C.int(stride)))
}

func (s *Shader) SetTexCoordArray(u *float32, stride int32) bool {
	return bool(C.al_set_shader_texcoord_array((*C.ALLEGRO_SHADER)(unsafe.Pointer(s)), (*C.float)(unsafe.Pointer(u)), C.int(stride)))
}

func (s *Shader) Set(display *Display) {
	C.al_set_shader((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)), (*C.ALLEGRO_SHADER)(unsafe.Pointer(s)))
}

func GetDefaultGlslVertexShader() string {
	return C.GoString(C.al_get_default_glsl_vertex_shader())
}

func GetDefaultGlslPixelShader() string {
	return C.GoString(C.al_get_default_glsl_pixel_shader())
}
