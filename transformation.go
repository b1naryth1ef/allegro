package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

type Transform struct {
	M [4][4]float32
}

func (t *Transform) Copy() *Transform {
	ts := new(C.ALLEGRO_TRANSFORM)
	C.al_copy_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), ts)
	return (*Transform)(unsafe.Pointer(ts))
}

func (t *Transform) Use() {
	C.al_use_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)))
}

func GetCurrentTransform() *Transform {
	return (*Transform)(unsafe.Pointer(C.al_get_current_transform()))
}

func GetCurrentInverseTransform() *Transform {
	return (*Transform)(unsafe.Pointer(C.al_get_current_inverse_transform()))
}

func (t *Transform) Invert() {
	C.al_invert_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)))
}

func (t *Transform) CheckInverse(tol float32) int32 {
	return int32(C.al_check_inverse((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(tol)))
}

func (t *Transform) Identity() {
	C.al_identity_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)))
}

func (t *Transform) Build(x, y, sx, sy, theta float32) {
	C.al_build_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(x), C.float(y), C.float(sx), C.float(sy), C.float(theta))
}

func (t *Transform) Translate(x, y float32) {
	C.al_translate_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(x), C.float(y))
}

func (t *Transform) Rotate(theta float32) {
	C.al_rotate_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(theta))
}

func (t *Transform) Scale(sx, sy float32) {
	C.al_scale_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(sx), C.float(sy))
}

func (t *Transform) TransformCoordinates(x, y float32) (float32, float32) {
	k := x
	l := y
	C.al_transform_coordinates((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), (*C.float)(unsafe.Pointer(&k)), (*C.float)(unsafe.Pointer(&l)))
	return k, l
}

func (t *Transform) Compose(other *Transform) {
	C.al_compose_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), (*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(other)))
}

func (t *Transform) Orthographic(left, top, n, right, bottom, f float32) {
	C.al_orthographic_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(left), C.float(top), C.float(n), C.float(right), C.float(bottom), C.float(f))
}

func (t *Transform) Perspective(left, top, n, right, bottom, f float32) {
	C.al_perspective_transform((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(left), C.float(top), C.float(n), C.float(right), C.float(bottom), C.float(f))
}

func (t *Transform) Translate3d(x, y, z float32) {
	C.al_translate_transform_3d((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(x), C.float(y), C.float(z))
}

func (t *Transform) Scale3d(sx, sy, sz float32) {
	C.al_scale_transform_3d((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(sx), C.float(sy), C.float(sz))
}

func (t *Transform) Rotate3d(x, y, z, angle float32) {
	C.al_rotate_transform_3d((*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)), C.float(x), C.float(y), C.float(z), C.float(angle))
}

func (t *Transform) SetProjection(display *Display) {
	C.al_set_projection_transform((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)), (*C.ALLEGRO_TRANSFORM)(unsafe.Pointer(t)))
}
