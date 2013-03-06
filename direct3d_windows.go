package allegro

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_direct3d.h>
import "C"

import "unsafe"

func HaveD3dNonPow2TextureSupport() bool {
	return bool(C.al_have_d3d_non_pow2_texture_support())
}

func HaveD3dNonSquareTextureSupport() bool {
	return bool(C.al_have_d3d_non_square_texture_support())
}

func (b *Bitmap) GetD3dTextureSize() (int32, int32) {
	var w, h C.int
	C.al_get_d3d_texture_size((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), &w, &h)
	return int32(w), int32(h)
}

func (b *Bitmap) GetD3dTexturePosition() (int32, int32) {
	var u, v C.int
	C.al_get_d3d_texture_position((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), &u, &v)
	return int32(u), int32(v)
}

func (d *Display) IsD3dDeviceLost() bool {
	return bool(C.al_is_d3d_device_lost((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}
