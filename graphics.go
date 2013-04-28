package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

/**********/
/* Colors */
/**********/

func MapRgb(r uint8, g uint8, b uint8) *Color {
	c := C.al_map_rgb(C.uchar(r), C.uchar(g), C.uchar(b))
	return (*Color)(unsafe.Pointer(&c))
}

func MapRgbF(r float32, g float32, b float32) *Color {
	c := C.al_map_rgb_f(C.float(r), C.float(g), C.float(b))
	return (*Color)(unsafe.Pointer(&c))
}

func MapRgba(r uint8, g uint8, b uint8, a uint8) *Color {
	c := C.al_map_rgba(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a))
	return (*Color)(unsafe.Pointer(&c))
}

func MapRgbaF(r float32, g float32, b float32, a float32) *Color {
	c := C.al_map_rgba_f(C.float(r), C.float(g), C.float(b), C.float(a))
	return (*Color)(unsafe.Pointer(&c))
}

func (c *Color) UnmapRgb() (uint8, uint8, uint8) {
	var r, g, b C.uchar
	C.al_unmap_rgb((C.ALLEGRO_COLOR)(*c), &r, &g, &b)
	return uint8(r), uint8(g), uint8(b)
}

func (c *Color) UnmapRgbF() (float32, float32, float32) {
	var r, g, b C.float
	C.al_unmap_rgb_f((C.ALLEGRO_COLOR)(*c), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func (c *Color) UnmapRgba() (uint8, uint8, uint8, uint8) {
	var r, g, b, a C.uchar
	C.al_unmap_rgba((C.ALLEGRO_COLOR)(*c), &r, &g, &b, &a)
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func (c *Color) UnmapRgbaF() (float32, float32, float32, float32) {
	var r, g, b, a C.float
	C.al_unmap_rgba_f((C.ALLEGRO_COLOR)(*c), &r, &g, &b, &a)
	return float32(r), float32(g), float32(b), float32(a)
}

/*****************************/
/* Locking and Pixel Formats */
/*****************************/

func GetPixelSize(format int32) int32 {
	return int32(C.al_get_pixel_size(C.int(format)))
}

func GetPixelFormatBits(format int32) int32 {
	return int32(C.al_get_pixel_format_bits(C.int(format)))
}

func (b *Bitmap) Lock(format, flags int32) *LockedRegion {
	l := C.al_lock_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.int(format), C.int(flags))
	return (*LockedRegion)(unsafe.Pointer(l))
}

func (b *Bitmap) LockRegion(x, y, width, height, format, flags int32) *LockedRegion {
	l := C.al_lock_bitmap_region((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.int(x), C.int(y), C.int(width), C.int(height), C.int(format), C.int(flags))
	return (*LockedRegion)(unsafe.Pointer(l))
}

func (b *Bitmap) Unlock() {
	C.al_unlock_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)))
}

/*******************/
/* Bitmap Creation */
/*******************/

func CreateBitmap(w, h int32) *Bitmap {
	return (*Bitmap)(unsafe.Pointer(C.al_create_bitmap(C.int(w), C.int(h))))
}

func (b *Bitmap) Destroy() {
	C.al_destroy_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)))
}

func (b *Bitmap) CreateSub(x, y, w, h int32) *Bitmap {
	return (*Bitmap)(unsafe.Pointer(C.al_create_sub_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.int(x), C.int(y), C.int(w), C.int(h))))
}

func (b *Bitmap) Clone() *Bitmap {
	return (*Bitmap)(unsafe.Pointer(C.al_clone_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)))))
}

func (b *Bitmap) Convert() {
	C.al_convert_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)))
}

func ConvertBitmaps() {
	C.al_convert_bitmaps()
}

func GetNewBitmapFlags() int32 {
	return int32(C.al_get_new_bitmap_flags())
}

func SetNewBitmapFlags(flags int32) {
	C.al_set_new_bitmap_flags(C.int(flags))
}

func AddNewBitmapFlag(flag int32) {
	C.al_add_new_bitmap_flag(C.int(flag))
}

func GetNewBitmapFormat() int32 {
	return int32(C.al_get_new_bitmap_format())
}

func SetNewBitmapFormat(format int32) {
	C.al_set_new_bitmap_format(C.int(format))
}

/*********************/
/* Bitmap Properties */
/*********************/

func (b *Bitmap) GetFlags() int32 {
	return int32(C.al_get_bitmap_flags((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) GetFormat() int32 {
	return int32(C.al_get_bitmap_format((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) GetHeight() int32 {
	return int32(C.al_get_bitmap_height((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) GetWidth() int32 {
	return int32(C.al_get_bitmap_width((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) GetPixel(x, y int32) Color {
	return (Color)(C.al_get_pixel((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.int(x), C.int(y)))
}

func (b *Bitmap) IsLocked() bool {
	return bool(C.al_is_bitmap_locked((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) IsCompatible() bool {
	return bool(C.al_is_compatible_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) IsSub() bool {
	return bool(C.al_is_sub_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) GetParent() *Bitmap {
	return (*Bitmap)(unsafe.Pointer(C.al_get_parent_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)))))
}

/**********************/
/* Drawing Operations */
/**********************/
func ClearToColor(c *Color) {
	C.al_clear_to_color((C.ALLEGRO_COLOR)(*c))
}

func ClearDepthBuffer(z float32) {
	C.al_clear_depth_buffer(C.float(z))
}

func (b *Bitmap) Draw(dx, dy float32, flags int32) {
	C.al_draw_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawTinted(tint *Color, dx, dy float32, flags int32) {
	C.al_draw_tinted_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), (C.ALLEGRO_COLOR)(*tint), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawRegion(sx, sy, sw, sh, dx, dy float32, flags int32) {
	C.al_draw_bitmap_region((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawTintedRegion(tint *Color, sx, sy, sw, sh, dx, dy float32, flags int32) {
	C.al_draw_tinted_bitmap_region((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), (C.ALLEGRO_COLOR)(*tint), C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (b *Bitmap) DrawPixel(x, y float32, c *Color) {
	C.al_draw_pixel(C.float(x), C.float(y), (C.ALLEGRO_COLOR)(*c))
}

func (b *Bitmap) DrawRotated(cx, cy, dx, dy, angle float32, flags int32) {
	C.al_draw_rotated_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawTintedRotated(c *Color, cx, cy, dx, dy, angle float32, flags int32) {
	C.al_draw_tinted_rotated_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), (C.ALLEGRO_COLOR)(*c), C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawScaledRotated(cx, cy, dx, dy, xscale, yscale, angle float32, flags int32) {
	C.al_draw_scaled_rotated_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawTintedScaledRotated(c *Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags int32) {
	C.al_draw_tinted_scaled_rotated_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), (C.ALLEGRO_COLOR)(*c), C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawTintedScaledRotatedRegion(sx, sy, sw, sh float32, c *Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
	C.al_draw_tinted_scaled_rotated_bitmap_region((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.float(sx), C.float(sy), C.float(sw), C.float(sh), (C.ALLEGRO_COLOR)(*c), C.float(cx), C.float(cy), C.float(dx), C.float(dy), C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

func (b *Bitmap) DrawScaled(sx, sy, sw, sh, dx, dy, dw, dh float32, flags int32) {
	C.al_draw_scaled_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.float(dw), C.float(dh), C.int(flags))
}

func (b *Bitmap) DrawTintedScaled(c *Color, sx, sy, sw, sh, dx, dy, dw, dh float32, flags int32) {
	C.al_draw_tinted_scaled_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), (C.ALLEGRO_COLOR)(*c), C.float(sx), C.float(sy), C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.float(dw), C.float(dh), C.int(flags))
}

func GetTargetBitmap() *Bitmap {
	return (*Bitmap)(unsafe.Pointer(C.al_get_target_bitmap()))
}

func (b *Bitmap) SetTarget() {
	C.al_set_target_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)))
}

func PutPixel(x, y int32, c *Color) {
	C.al_put_pixel(C.int(x), C.int(y), (C.ALLEGRO_COLOR)(*c))
}

func PutBlendedPixel(x, y int32, c *Color) {
	C.al_put_blended_pixel(C.int(x), C.int(y), (C.ALLEGRO_COLOR)(*c))
}

func (d *Display) SetTargetBackbuffer() {
	C.al_set_target_backbuffer((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))
}

func GetCurrentDisplay() *Display {
	return (*Display)(unsafe.Pointer(C.al_get_current_display()))
}

/******************/
/* Blending modes */
/******************/

func GetBlender() (int32, int32, int32) {
	var op, src, dst C.int
	C.al_get_blender(&op, &src, &dst)
	return int32(op), int32(src), int32(dst)
}

func SetBlender(op, src, dst int32) {
	C.al_set_blender(C.int(op), C.int(src), C.int(dst))
}

func GetSeparateBlender() (int32, int32, int32, int32, int32, int32) {
	var op, src, dst, alpha_op, alpha_src, alpha_dst C.int
	C.al_get_separate_blender(&op, &src, &dst, &alpha_op, &alpha_src, &alpha_dst)
	return int32(op), int32(src), int32(dst), int32(alpha_op), int32(alpha_src), int32(alpha_dst)
}

func SetSeparateBlender(op, src, dst, alpha_op, alpha_src, alpha_dst int32) {
	C.al_set_separate_blender(C.int(op), C.int(src), C.int(dst), C.int(alpha_op), C.int(alpha_src), C.int(alpha_dst))
}

/************/
/* Clipping */
/************/

func GetClippingRectangle() (int32, int32, int32, int32) {
	var x, y, w, h C.int
	C.al_get_clipping_rectangle(&x, &y, &w, &h)
	return int32(x), int32(y), int32(w), int32(h)
}

func SetClippingRectangle(x, y, w, h int32) {
	C.al_set_clipping_rectangle(C.int(x), C.int(y), C.int(w), C.int(h))
}

func ResetClippingRectangle() {
	C.al_reset_clipping_rectangle()
}

/******************************/
/* Graphics Utility Functions */
/******************************/

func (b *Bitmap) ConvertMaskToAlpha(maskColor *Color) {
	C.al_convert_mask_to_alpha((*C.ALLEGRO_BITMAP)(unsafe.Pointer(b)), (C.ALLEGRO_COLOR)(*maskColor))
}

/********************/
/* Deferred Drawing */
/********************/

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
	return (*Bitmap)(unsafe.Pointer(C.al_load_bitmap(f)))
}

func LoadBitmapF(fp *File, ident string) *Bitmap {
	i := C.CString(ident)
	defer C.free(unsafe.Pointer(i))
	return (*Bitmap)(unsafe.Pointer(C.al_load_bitmap_f((*C.ALLEGRO_FILE)(unsafe.Pointer(fp)), i)))
}

func LoadBitmapFlagsF(fp *File, ident string, flags int32) *Bitmap {
	i := C.CString(ident)
	defer C.free(unsafe.Pointer(i))
	return (*Bitmap)(unsafe.Pointer(C.al_load_bitmap_flags_f((*C.ALLEGRO_FILE)(unsafe.Pointer(fp)), i, C.int(flags))))
}

func (b *Bitmap) Save(fileName string) bool {
	f := C.CString(fileName)
	defer C.free(unsafe.Pointer(f))
	return bool(C.al_save_bitmap(f, (*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

func (b *Bitmap) SaveF(fp *File, ident string) bool {
	i := C.CString(ident)
	defer C.free(unsafe.Pointer(i))
	return bool(C.al_save_bitmap_f((*C.ALLEGRO_FILE)(unsafe.Pointer(fp)), i, (*C.ALLEGRO_BITMAP)(unsafe.Pointer(b))))
}

/****************/
/* Render State */
/****************/

func SetRenderState(state uint32, value int32) {
	C.al_set_render_state(C.ALLEGRO_RENDER_STATE(state), C.int(value))
}

/************************/
/* LockedRegion getters */
/************************/

func (l *LockedRegion) Data() unsafe.Pointer {
	lc := (*C.ALLEGRO_LOCKED_REGION)(unsafe.Pointer(l))
	return unsafe.Pointer(lc.data)
}

func (l *LockedRegion) Format() int32 {
	lc := (*C.ALLEGRO_LOCKED_REGION)(unsafe.Pointer(l))
	return int32(lc.format)
}

func (l *LockedRegion) Pitch() int32 {
	lc := (*C.ALLEGRO_LOCKED_REGION)(unsafe.Pointer(l))
	return int32(lc.pitch)
}

func (l *LockedRegion) PixelSize() int32 {
	lc := (*C.ALLEGRO_LOCKED_REGION)(unsafe.Pointer(l))
	return int32(lc.pixel_size)
}
