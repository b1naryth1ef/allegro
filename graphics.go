package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

/**********/
/* Colors */
/**********/

type Color struct {
	allegroColor C.ALLEGRO_COLOR
}

func MapRgb(r uint8, g uint8, b uint8) *Color {
	return &Color{C.al_map_rgb(C.uchar(r), C.uchar(g), C.uchar(b))}
}

func MapRgbF(r float32, g float32, b float32) *Color {
	return &Color{C.al_map_rgb_f(C.float(r), C.float(g), C.float(b))}
}

func MapRgba(r uint8, g uint8, b uint8, a uint8) *Color {
	return &Color{C.al_map_rgba(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a))}
}

func MapRgbaF(r float32, g float32, b float32, a float32) *Color {
	return &Color{C.al_map_rgba_f(C.float(r), C.float(g), C.float(b), C.float(a))}
}

func (c *Color) UnmapRgb() (uint8, uint8, uint8) {
	var r, g, b C.uchar
	C.al_unmap_rgb(c.allegroColor, &r, &g, &b)
	return uint8(r), uint8(g), uint8(b)
}

func (c *Color) UnmapRgbF() (float32, float32, float32) {
	var r, g, b C.float
	C.al_unmap_rgb_f(c.allegroColor, &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func (c *Color) UnmapRgba() (uint8, uint8, uint8, uint8) {
	var r, g, b, a C.uchar
	C.al_unmap_rgba(c.allegroColor, &r, &g, &b, &a)
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func (c *Color) UnmapRgbaF() (float32, float32, float32, float32) {
	var r, g, b, a C.float
	C.al_unmap_rgba_f(c.allegroColor, &r, &g, &b, &a)
	return float32(r), float32(g), float32(b), float32(a)
}

/*****************************/
/* Locking and Pixel Formats */
/*****************************/

type LockedRegion struct {
	data      *byte
	format    int
	pitch     int
	pixelSize int
}

const (
	PixelFormatAny            = C.ALLEGRO_PIXEL_FORMAT_ANY
	PixelFormatAnyNoAlpha     = C.ALLEGRO_PIXEL_FORMAT_ANY_NO_ALPHA
	PixelFormatAnyWithAlpha   = C.ALLEGRO_PIXEL_FORMAT_ANY_WITH_ALPHA
	PixelFormatAny15NoAlpha   = C.ALLEGRO_PIXEL_FORMAT_ANY_15_NO_ALPHA
	PixelFormatAny16NoAlpha   = C.ALLEGRO_PIXEL_FORMAT_ANY_16_NO_ALPHA
	PixelFormatAny16WithAlpha = C.ALLEGRO_PIXEL_FORMAT_ANY_16_WITH_ALPHA
	PixelFormatAny24NoAlpha   = C.ALLEGRO_PIXEL_FORMAT_ANY_24_NO_ALPHA
	PixelFormatAny32NoAlpha   = C.ALLEGRO_PIXEL_FORMAT_ANY_32_NO_ALPHA
	PixelFormatAny32WithAlpha = C.ALLEGRO_PIXEL_FORMAT_ANY_32_WITH_ALPHA
	PixelFormatArgb8888       = C.ALLEGRO_PIXEL_FORMAT_ARGB_8888
	PixelFormatRgba8888       = C.ALLEGRO_PIXEL_FORMAT_RGBA_8888
	PixelFormatArgb4444       = C.ALLEGRO_PIXEL_FORMAT_ARGB_4444
	PixelFormatRgb888         = C.ALLEGRO_PIXEL_FORMAT_RGB_888
	PixelFormatRgb565         = C.ALLEGRO_PIXEL_FORMAT_RGB_565
	PixelFormatRgb555         = C.ALLEGRO_PIXEL_FORMAT_RGB_555
	PixelFormatRgba5551       = C.ALLEGRO_PIXEL_FORMAT_RGBA_5551
	PixelFormatArgb1555       = C.ALLEGRO_PIXEL_FORMAT_ARGB_1555
	PixelFormatAbgr8888       = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888
	PixelFormatXbgr8888       = C.ALLEGRO_PIXEL_FORMAT_XBGR_8888
	PixelFormatBgr888         = C.ALLEGRO_PIXEL_FORMAT_BGR_888
	PixelFormatBgr565         = C.ALLEGRO_PIXEL_FORMAT_BGR_565
	PixelFormatBgr555         = C.ALLEGRO_PIXEL_FORMAT_BGR_555
	PixelFormatRgbx8888       = C.ALLEGRO_PIXEL_FORMAT_RGBX_8888
	PixelFormatXrgb8888       = C.ALLEGRO_PIXEL_FORMAT_XRGB_8888
	PixelFormatAbgrF32        = C.ALLEGRO_PIXEL_FORMAT_ABGR_F32
	PixelFormatAbgr8888Le     = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888_LE
	PixelFormatRgba4444       = C.ALLEGRO_PIXEL_FORMAT_RGBA_4444
)

func GetPixelSize(format int) int {
	return int(C.al_get_pixel_size(C.int(format)))
}

func GetPixelFormatBits(format int) int {
	return int(C.al_get_pixel_format_bits(C.int(format)))
}

func (b *Bitmap) LockBitmap(format, flags int) *LockedRegion {
	l := C.al_lock_bitmap(b.allegroBitmap, C.int(format), C.int(flags))
	return (*LockedRegion)(unsafe.Pointer(l))
}

// flags
const (
	LockReadonly  = C.ALLEGRO_LOCK_READONLY
	LockWriteOnly = C.ALLEGRO_LOCK_WRITEONLY
	LockReadWrite = C.ALLEGRO_LOCK_READWRITE
)

func (b *Bitmap) LockBitmapRegion(x, y, width, height, format, flags int) *LockedRegion {
	l := C.al_lock_bitmap_region(b.allegroBitmap, C.int(x), C.int(y), C.int(width), C.int(height), C.int(format), C.int(flags))
	return (*LockedRegion)(unsafe.Pointer(l))
}

func (b *Bitmap) UnlockBitmap() {
	C.al_unlock_bitmap(b.allegroBitmap)
}

/*******************/
/* Bitmap Creation */
/*******************/

type Bitmap struct {
	allegroBitmap *C.ALLEGRO_BITMAP
}

func CreateBitmap(w, h int) *Bitmap {
	return &Bitmap{C.al_create_bitmap(C.int(w), C.int(h))}
}

func (b *Bitmap) CreateSubBitmap(x, y, w, h int) *Bitmap {
	return &Bitmap{C.al_create_sub_bitmap(b.allegroBitmap, C.int(x), C.int(y), C.int(w), C.int(h))}
}

func (b *Bitmap) CloneBitmap() *Bitmap {
	return &Bitmap{C.al_clone_bitmap(b.allegroBitmap)}
}

func (b *Bitmap) DestroyBitmap() {
	C.al_destroy_bitmap(b.allegroBitmap)
}

func GetNewBitmapFlags() int {
	return int(C.al_get_new_bitmap_flags())
}

func GetNewBitmapFormat() int {
	return int(C.al_get_new_bitmap_format())
}

// flags
const (
	VideoBitmap          = C.ALLEGRO_VIDEO_BITMAP
	MemoryBitmap         = C.ALLEGRO_MEMORY_BITMAP
	KeepBitmapFormat     = C.ALLEGRO_KEEP_BITMAP_FORMAT
	ForceLocking         = C.ALLEGRO_FORCE_LOCKING
	NoPreserveTexture    = C.ALLEGRO_NO_PRESERVE_TEXTURE
	AlphaTest            = C.ALLEGRO_ALPHA_TEST
	MinLinear            = C.ALLEGRO_MIN_LINEAR
	MagLinear            = C.ALLEGRO_MAG_LINEAR
	Mipmap               = C.ALLEGRO_MIPMAP
	NoPremultipliedAlpha = C.ALLEGRO_NO_PREMULTIPLIED_ALPHA
)

func SetNewBitmapFlags(flags int) {
	C.al_set_new_bitmap_flags(C.int(flags))
}

func AddNewBitmapFlag(flag int) {
	C.al_add_new_bitmap_flag(C.int(flag))
}

func SetNewBitmapFormat(format int) {
	C.al_set_new_bitmap_format(C.int(format))
}

/*********************/
/* Bitmap Properties */
/*********************/

func (b *Bitmap) GetBitmapFlags() int {
	return int(C.al_get_bitmap_flags(b.allegroBitmap))
}

func (b *Bitmap) GetBitmapFormat() int {
	return int(C.al_get_bitmap_format(b.allegroBitmap))
}

func (b *Bitmap) GetBitmapHeight() int {
	return int(C.al_get_bitmap_height(b.allegroBitmap))
}

func (b *Bitmap) GetBitmapWidth() int {
	return int(C.al_get_bitmap_width(b.allegroBitmap))
}

func (b *Bitmap) GetPixel(x, y int) *Color {
	return &Color{C.al_get_pixel(b.allegroBitmap, C.int(x), C.int(y))}
}

func (b *Bitmap) IsBitmapLocked() bool {
	return bool(C.al_is_bitmap_locked(b.allegroBitmap))
}

func (b *Bitmap) IsCompatibleBitmap() bool {
	return bool(C.al_is_compatible_bitmap(b.allegroBitmap))
}

func (b *Bitmap) IsSubBitmap() bool {
	return bool(C.al_is_sub_bitmap(b.allegroBitmap))
}

//func (b *Bitmap) GetParentBitmap() *Bitmap {
//	return &Bitmap{C.al_get_parent_bitmap(b.allegroBitmap)}
//}

/**********************/
/* Drawing Operations */
/**********************/
func ClearToColor(c *Color) {
	C.al_clear_to_color(c.allegroColor)
}

// flags
const (
	FlipHorizontal = C.ALLEGRO_FLIP_HORIZONTAL
	FlipVertical   = C.ALLEGRO_FLIP_VERTICAL
)

func (b *Bitmap) DrawBitmap(dx, dy float32, flags int) {
	C.al_draw_bitmap(b.allegroBitmap, C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawTintedBitmap(tint *Color, dx, dy float32, flags int) {
	C.al_draw_tinted_bitmap(b.allegroBitmap, tint.allegroColor, C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawBitmapRegion(sx, sy, sw, sh, dx, dy float32, flags int) {
	C.al_draw_bitmap_region(b.allegroBitmap, C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawTintedBitmapRegion(tint *Color, sx, sy, sw, sh, dx, dy float32, flags int) {
	C.al_draw_tinted_bitmap_region(b.allegroBitmap, tint.allegroColor, C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawPixel(x, y float32, c *Color) {
	C.al_draw_pixel(C.float(x), C.float(y), c.allegroColor)
}

func (b *Bitmap) DrawRotatedBitmap(cx, cy, dx, dy, angle float32, flags int) {
	C.al_draw_rotated_bitmap(b.allegroBitmap, C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawTintedRotatedBitmap(c *Color, cx, cy, dx, dy, angle float32, flags int) {
	C.al_draw_tinted_rotated_bitmap(b.allegroBitmap, c.allegroColor, C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawScaledRotatedBitmap(cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
	C.al_draw_scaled_rotated_bitmap(b.allegroBitmap, C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawTintedScaledRotatedBitmap(c *Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
	C.al_draw_tinted_scaled_rotated_bitmap(b.allegroBitmap, c.allegroColor, C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

//func (b *Bitmap) DrawTintedScaledRotatedBitmapRegion(sx, sy, sw, sh float32, c *Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
//	C.al_draw_tinted_scaled_rotated_bitmap_region(b.allegroBitmap, C.float(sx), C.float(sy), C.float(sw), C.float(sh), c.allegroColor, C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
//}

func (b *Bitmap) DrawScaledBitmap(sx, sy, sw, sh, dx, dy, dw, dh float32, flags int) {
	C.al_draw_scaled_bitmap(b.allegroBitmap, C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.float(dw), C.float(dh), C.int(flags))
}

func (b *Bitmap) DrawTintedScaledBitmap(c *Color, sx, sy, sw, sh, dx, dy, dw, dh float32, flags int) {
	C.al_draw_tinted_scaled_bitmap(b.allegroBitmap, c.allegroColor, C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.float(dw), C.float(dh), C.int(flags))
}

func GetTargetBitmap() *Bitmap {
	return &Bitmap{C.al_get_target_bitmap()}
}

func PutPixel(x, y int, c *Color) {
	C.al_put_pixel(C.int(x), C.int(y), c.allegroColor)
}

func PutBlendedPixel(x, y int, c *Color) {
	C.al_put_blended_pixel(C.int(x), C.int(y), c.allegroColor)
}

func (b *Bitmap) SetTargetBitmap() {
	C.al_set_target_bitmap(b.allegroBitmap)
}

func (d *Display) SetTargetBackbuffer() {
	C.al_set_target_backbuffer(d.allegroDisplay)
}

func GetCurrentDisplay() *Display {
	return &Display{C.al_get_current_display()}
}

/******************/
/* Blendind modes */
/******************/

func GetBlender() (int, int, int) {
	var op, src, dst C.int
	C.al_get_blender(&op, &src, &dst)
	return int(op), int(src), int(dst)
}

func GetSeparateBlender() (int, int, int, int, int, int) {
	var op, src, dst, alpha_op, alpha_src, alpha_dst C.int
	C.al_get_separate_blender(&op, &src, &dst, &alpha_op, &alpha_src, &alpha_dst)
	return int(op), int(src), int(dst), int(alpha_op), int(alpha_src), int(alpha_dst)
}

func SetBlender(op, src, dst int) {
	C.al_set_blender(C.int(op), C.int(src), C.int(dst))
}

func SetSeparateBlender(op, src, dst, alpha_op, alpha_src, alpha_dst int) {
	C.al_set_separate_blender(C.int(op), C.int(src), C.int(dst), C.int(alpha_op), C.int(alpha_src), C.int(alpha_dst))
}

/************/
/* Clipping */
/************/

func GetClippingRectangle() (int, int, int, int) {
	var x, y, w, h C.int
	C.al_get_clipping_rectangle(&x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func SetClippingRectangle(x, y, w, h int) {
	C.al_set_clipping_rectangle(C.int(x), C.int(y), C.int(w), C.int(h))
}

//func ResetClippingRectangle() {
//	C.al_reset_clipping_rectangle()
//}

/******************************/
/* Graphics utility functions */
/******************************/

func (b *Bitmap) ConvertMaskToAlpha(maskColor *Color) {
	C.al_convert_mask_to_alpha(b.allegroBitmap, maskColor.allegroColor)
}

func HoldBitmapDrawing(hold bool) {
	C.al_hold_bitmap_drawing(C.bool(hold))
}

func IsBitmapDrawingHeld() bool {
	return bool(C.al_is_bitmap_drawing_held())
}

/************/
/* Image IO */
/************/

func LoadBitmap(fileName string) *Bitmap {
	f := C.CString(fileName)
	defer C.free(unsafe.Pointer(f))
	return &Bitmap{C.al_load_bitmap(f)}
}
