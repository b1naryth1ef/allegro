package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

const (
	Windowed                = C.ALLEGRO_WINDOWED
	Fullscreen              = C.ALLEGRO_FULLSCREEN
	FullscreenWindow        = C.ALLEGRO_FULLSCREEN_WINDOW
	Resizable               = C.ALLEGRO_RESIZABLE
	Opengl                  = C.ALLEGRO_OPENGL
	Opengl30                = C.ALLEGRO_OPENGL_3_0
	OpenglForwardCompatible = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
	//	Direct3d                = C.ALLEGRO_DIRECT3D
	Frameless            = C.ALLEGRO_FRAMELESS
	Noframe              = C.ALLEGRO_NOFRAME
	GenerateExposeEvents = C.ALLEGRO_GENERATE_EXPOSE_EVENTS
)

const (
	ColorSize               = C.ALLEGRO_COLOR_SIZE
	RedSize                 = C.ALLEGRO_RED_SIZE
	GreenSize               = C.ALLEGRO_GREEN_SIZE
	BlueSize                = C.ALLEGRO_BLUE_SIZE
	AlphaSize               = C.ALLEGRO_ALPHA_SIZE
	RedShift                = C.ALLEGRO_RED_SHIFT
	GreenShift              = C.ALLEGRO_GREEN_SHIFT
	BlueShift               = C.ALLEGRO_BLUE_SHIFT
	AlphaShift              = C.ALLEGRO_ALPHA_SHIFT
	AccRedSize              = C.ALLEGRO_ACC_RED_SIZE
	AccGreenSize            = C.ALLEGRO_ACC_GREEN_SIZE
	AccBlueSize             = C.ALLEGRO_ACC_BLUE_SIZE
	AccAlphaSize            = C.ALLEGRO_ACC_ALPHA_SIZE
	Stereo                  = C.ALLEGRO_STEREO
	AuxBuffers              = C.ALLEGRO_AUX_BUFFERS
	DepthSize               = C.ALLEGRO_DEPTH_SIZE
	StencilSize             = C.ALLEGRO_STENCIL_SIZE
	SampleBuffers           = C.ALLEGRO_SAMPLE_BUFFERS
	Samples                 = C.ALLEGRO_SAMPLES
	RenderMethod            = C.ALLEGRO_RENDER_METHOD
	FloatColor              = C.ALLEGRO_FLOAT_COLOR
	FloatDepth              = C.ALLEGRO_FLOAT_DEPTH
	SingleBuffer            = C.ALLEGRO_SINGLE_BUFFER
	SwapMethod              = C.ALLEGRO_SWAP_METHOD
	CompatibleDisplay       = C.ALLEGRO_COMPATIBLE_DISPLAY
	UpdateDisplayRegionFlag = C.ALLEGRO_UPDATE_DISPLAY_REGION
	Vsync                   = C.ALLEGRO_VSYNC
	MaxBitmapSize           = C.ALLEGRO_MAX_BITMAP_SIZE
	SupportNpotBitmap       = C.ALLEGRO_SUPPORT_NPOT_BITMAP
	CanDrawIntoBitmap       = C.ALLEGRO_CAN_DRAW_INTO_BITMAP
	SupportSeparateAlpha    = C.ALLEGRO_SUPPORT_SEPARATE_ALPHA
)

const (
	Require  = C.ALLEGRO_REQUIRE
	Suggest  = C.ALLEGRO_SUGGEST
	DontCare = C.ALLEGRO_DONTCARE
)

/********************/
/* Display Creation */
/********************/

type Display C.ALLEGRO_DISPLAY

func CreateDisplay(w, h int32) *Display {
	return (*Display)(unsafe.Pointer(C.al_create_display(C.int(w), C.int(h))))
}

func (d *Display) Destroy() {
	C.al_destroy_display((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))
}

func GetNewDisplayFlags() int32 {
	return int32(C.al_get_new_display_flags())
}

func SetNewDisplayFlags(flags int32) {
	C.al_set_new_display_flags(C.int(flags))
}

func GetNewDisplayOption(option int32) (int32, int32) {
	var importance C.int
	v := int32(C.al_get_new_display_option(C.int(option), &importance))
	return v, int32(importance)
}

func SetNewDisplayOption(option, value, importance int32) {
	C.al_set_new_display_option(C.int(option), C.int(value), C.int(importance))
}

func ResetNewDisplayOptions() {
	C.al_reset_new_display_options()
}

func GetNewWindowPosition() (int32, int32) {
	var x, y C.int
	C.al_get_new_window_position(&x, &y)
	return int32(x), int32(y)
}

func SetNewWindowPosition(x, y int32) {
	C.al_set_new_window_position(C.int(x), C.int(y))
}

func GetNewDisplayRefreshRate() int32 {
	return int32(C.al_get_new_display_refresh_rate())
}

func SetNewDisplayRefreshRate(refreshRate int32) {
	C.al_set_new_display_refresh_rate(C.int(refreshRate))
}

/**********************/
/* Display Operations */
/**********************/

func (d *Display) GetEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_display_event_source((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))))
}

func (d *Display) GetBackbuffer() *Bitmap {
	return (*Bitmap)(unsafe.Pointer(C.al_get_backbuffer((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))))
}

func FlipDisplay() {
	C.al_flip_display()
}

func UpdateDisplayRegion(x, y, width, height int32) {
	C.al_update_display_region(C.int(x), C.int(y), C.int(width), C.int(height))
}

func WaitForVsync() bool {
	return bool(C.al_wait_for_vsync())
}

/*****************************/
/* Display Size and Position */
/*****************************/

func (d *Display) GetWidth() int32 {
	return int32(C.al_get_display_width((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) GetHeight() int32 {
	return int32(C.al_get_display_height((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) Resize(width, height int32) bool {
	return bool(C.al_resize_display((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(width), C.int(height)))
}

func (d *Display) AcknowledgeResize() bool {
	return bool(C.al_acknowledge_resize((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) GetWindowPosition() (int32, int32) {
	var x, y C.int
	C.al_get_window_position((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), &x, &y)
	return int32(x), int32(y)
}

func (d *Display) SetWindowPosition(x, y int32) {
	C.al_set_window_position((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(x), C.int(y))
}

/********************/
/* Display Settings */
/********************/

func (d *Display) GetFlags() int32 {
	return int32(C.al_get_display_flags((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) SetDisplayFlag(flag int, onoff bool) bool {
	return bool(C.al_set_display_flag((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(flag), C.bool(onoff)))
}

func (d *Display) GetOption(option int32) int32 {
	return int32(C.al_get_display_option((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), C.int(option)))
}

func (d *Display) GetFormat() int32 {
	return int32(C.al_get_display_format((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) GetRefreshRate() int32 {
	return int32(C.al_get_display_refresh_rate((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d))))
}

func (d *Display) SetWindowTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.al_set_window_title((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), t)
}

func (d *Display) SetIcon(icon *Bitmap) {
	C.al_set_display_icon((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)), (*C.ALLEGRO_BITMAP)(unsafe.Pointer(icon)))
}

//func (d *Display) SetIcons(numIcons int, icons [](*Bitmap)) {
//	C.al_set_display_icons(C.int(numIcons), [](*C.ALLEGRO_BITMAP)(icons))
//}

/***************/
/* Screensaver */
/***************/

func InhibitScreensaver(inhibit bool) bool {
	return bool(C.al_inhibit_screensaver(C.bool(inhibit)))
}
