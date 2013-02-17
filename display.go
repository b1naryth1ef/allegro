package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

/********************/
/* Display Creation */
/********************/

type Display struct {
	allegroDisplay *C.ALLEGRO_DISPLAY
}

func CreateDisplay(w, h int) *Display {
	return &Display{C.al_create_display(C.int(w), C.int(h))}
}

func (d *Display) DestroyDisplay() {
	C.al_destroy_display(d.allegroDisplay)
}

func GetNewDisplayFlags() int {
	return int(C.al_get_new_display_flags())
}

func GetNewDisplayRefreshRate() int {
	return int(C.al_get_new_display_refresh_rate())
}

func GetNewWindowPosition() (int, int) {
	var x, y C.int
	C.al_get_new_window_position(&x, &y)
	return int(x), int(y)
}

// importance
const (
	Require  = C.ALLEGRO_REQUIRE
	Suggest  = C.ALLEGRO_SUGGEST
	DontCare = C.ALLEGRO_DONTCARE
)

// option
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

func SetNewDisplayOption(option, value, importance int) {
	C.al_set_new_display_option(C.int(option), C.int(value), C.int(importance))
}

func ResetNewDisplayOptions() {
	C.al_reset_new_display_options()
}

// flags
const (
	Windowed                = C.ALLEGRO_WINDOWED
	Fullscreen              = C.ALLEGRO_FULLSCREEN
	FullscreenWindow        = C.ALLEGRO_FULLSCREEN_WINDOW
	Resizable               = C.ALLEGRO_RESIZABLE
	Opengl                  = C.ALLEGRO_OPENGL
	Opengl30                = C.ALLEGRO_OPENGL_3_0
	OpenglForwardCompatible = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
	//	Direct3d = C.ALLEGRO_DIRECT3D
	//	Frameless = C.ALLEGRO_FRAMELESS
	Noframe              = C.ALLEGRO_NOFRAME
	GenerateExposeEvents = C.ALLEGRO_GENERATE_EXPOSE_EVENTS
)

func SetNewDisplayFlags(flags int) {
	C.al_set_new_display_flags(C.int(flags))
}

func SetNewDisplayRefreshRate(refreshRate int) {
	C.al_set_new_display_refresh_rate(C.int(refreshRate))
}

func SetNewWindowPosition(x, y int) {
	C.al_set_new_window_position(C.int(x), C.int(y))
}

/**********************/
/* Display Operations */
/**********************/

func (d *Display) AcknowledgeResize() bool {
	return bool(C.al_acknowledge_resize(d.allegroDisplay))
}

func FlipDisplay() {
	C.al_flip_display()
}

func (d *Display) GetBackbuffer() *Bitmap {
	return &Bitmap{C.al_get_backbuffer(d.allegroDisplay)}
}

func (d *Display) GetDisplayFlags() int {
	return int(C.al_get_display_flags(d.allegroDisplay))
}

func (d *Display) GetDisplayFormat() int {
	return int(C.al_get_display_format(d.allegroDisplay))
}

func (d *Display) GetDisplayHeight() int {
	return int(C.al_get_display_height(d.allegroDisplay))
}

func (d *Display) GetDisplayRefreshRate() int {
	return int(C.al_get_display_refresh_rate(d.allegroDisplay))
}

func (d *Display) GetDisplayWidth() int {
	return int(C.al_get_display_width(d.allegroDisplay))
}

func (d *Display) GetWindowPosition() (int, int) {
	var x, y C.int
	C.al_get_window_position(d.allegroDisplay, &x, &y)
	return int(x), int(y)
}

func InhibitScreensaver(inhibit bool) bool {
	return bool(C.al_inhibit_screensaver(C.bool(inhibit)))
}

func (d *Display) ResizeDisplay(width, height int) bool {
	return bool(C.al_resize_display(d.allegroDisplay, C.int(width), C.int(height)))
}

func (d *Display) SetDisplayIcon(icon *Bitmap) {
	C.al_set_display_icon(d.allegroDisplay, icon.allegroBitmap)
}

func (d *Display) GetDisplayOption(option int) int {
	return int(C.al_get_display_option(d.allegroDisplay, C.int(option)))
}

func (d *Display) SetWindowPosition(x, y int) {
	C.al_set_window_position(d.allegroDisplay, C.int(x), C.int(y))
}

func (d *Display) SetWindowTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.al_set_window_title(d.allegroDisplay, t)
}

//func (d *Display) SetDisplayFlag(flag int, onoff bool) bool {
//	return bool(C.al_set_display_flag(d.allegroDisplay, C.int(flag), C.bool(onoff)))
//}

func (d *Display) ToggleDisplayFlag(flag int, onoff bool) bool {
	return bool(C.al_toggle_display_flag(d.allegroDisplay, C.int(flag), C.bool(onoff)))
}

func UpdateDisplayRegion(x, y, width, height int) {
	C.al_update_display_region(C.int(x), C.int(y), C.int(width), C.int(height))
}

func WaitForVsync() bool {
	return bool(C.al_wait_for_vsync())
}

/****************************/
/* Fullscreen Display Modes */
/****************************/

type DisplayMode struct {
	Width       int
	Height      int
	Format      int
	RefreshRate int
}

func GetNumDisplayModes() int {
	return int(C.al_get_num_display_modes())
}

/************/
/* Monitors */
/************/

type MonitorInfo struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func GetNewDisplayAdapter() int {
	return int(C.al_get_new_display_adapter())
}

func SetNewDisplayAdapter(adapter int) {
	C.al_set_new_display_adapter(C.int(adapter))
}

func GetNumVideoAdapters() int {
	return int(C.al_get_num_video_adapters())
}
