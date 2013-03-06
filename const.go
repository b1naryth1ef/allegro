package allegro

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_shader.h>
// #include <allegro5/allegro_opengl.h>
import "C"

// system.go
const (
	ResourcesPath     = C.ALLEGRO_RESOURCES_PATH
	TempPath          = C.ALLEGRO_TEMP_PATH
	UserDataPath      = C.ALLEGRO_USER_DATA_PATH
	UserHomePath      = C.ALLEGRO_USER_HOME_PATH
	UserSettingsPath  = C.ALLEGRO_USER_SETTINGS_PATH
	UserDocumentsPath = C.ALLEGRO_USER_DOCUMENTS_PATH
	ExenamePath       = C.ALLEGRO_EXENAME_PATH
)

// display.go
const (
	Windowed                = C.ALLEGRO_WINDOWED
	Fullscreen              = C.ALLEGRO_FULLSCREEN
	FullscreenWindow        = C.ALLEGRO_FULLSCREEN_WINDOW
	Resizable               = C.ALLEGRO_RESIZABLE
	Opengl                  = C.ALLEGRO_OPENGL
	Opengl30                = C.ALLEGRO_OPENGL_3_0
	OpenglForwardCompatible = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
	Frameless               = C.ALLEGRO_FRAMELESS
	Noframe                 = C.ALLEGRO_NOFRAME
	GenerateExposeEvents    = C.ALLEGRO_GENERATE_EXPOSE_EVENTS
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

const (
	DisplayOrientationUnknown    = C.ALLEGRO_DISPLAY_ORIENTATION_UNKNOWN
	DisplayOrientation0Degrees   = C.ALLEGRO_DISPLAY_ORIENTATION_0_DEGREES
	DisplayOrientation90Degrees  = C.ALLEGRO_DISPLAY_ORIENTATION_90_DEGREES
	DisplayOrientation180Degrees = C.ALLEGRO_DISPLAY_ORIENTATION_180_DEGREES
	DisplayOrientation270Degrees = C.ALLEGRO_DISPLAY_ORIENTATION_270_DEGREES
	DisplayOrientationFaceUp     = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_UP
	DisplayOrientationFaceDown   = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_DOWN
)

// events.go
const (
	EventJoystickAxis          = C.ALLEGRO_EVENT_JOYSTICK_AXIS
	EventJoystickButtonDown    = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN
	EventJoystickButtonUp      = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP
	EventJoystickConfiguration = C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION
	EventKeyDown               = C.ALLEGRO_EVENT_KEY_DOWN
	EventKeyChar               = C.ALLEGRO_EVENT_KEY_CHAR
	EventKeyUp                 = C.ALLEGRO_EVENT_KEY_UP
	EventMouseAxes             = C.ALLEGRO_EVENT_MOUSE_AXES
	EventMouseButtonDown       = C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN
	EventMouseButtonUp         = C.ALLEGRO_EVENT_MOUSE_BUTTON_UP
	EventMouseEnterDisplay     = C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY
	EventMouseLeaveDisplay     = C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY
	EventMouseWarped           = C.ALLEGRO_EVENT_MOUSE_WARPED
	EventTimer                 = C.ALLEGRO_EVENT_TIMER
	EventDisplayExpose         = C.ALLEGRO_EVENT_DISPLAY_EXPOSE
	EventDisplayResize         = C.ALLEGRO_EVENT_DISPLAY_RESIZE
	EventDisplayClose          = C.ALLEGRO_EVENT_DISPLAY_CLOSE
	EventDisplayLost           = C.ALLEGRO_EVENT_DISPLAY_LOST
	EventDisplayFound          = C.ALLEGRO_EVENT_DISPLAY_FOUND
	EventDisplaySwitchIn       = C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN
	EventDisplaySwitchOut      = C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT
	EventDisplayOrientation    = C.ALLEGRO_EVENT_DISPLAY_ORIENTATION
	EventDisplayHaltDrawing    = C.ALLEGRO_EVENT_DISPLAY_HALT_DRAWING
	EventDisplayResumeDrawing  = C.ALLEGRO_EVENT_DISPLAY_RESUME_DRAWING
	EventDisplayConnected      = C.ALLEGRO_EVENT_DISPLAY_CONNECTED
	EventDisplayDisconnected   = C.ALLEGRO_EVENT_DISPLAY_DISCONNECTED
	EventTouchBegin            = C.ALLEGRO_EVENT_TOUCH_BEGIN
	EventTouchEnd              = C.ALLEGRO_EVENT_TOUCH_END
	EventTouchMove             = C.ALLEGRO_EVENT_TOUCH_MOVE
	EventTouchCancel           = C.ALLEGRO_EVENT_TOUCH_CANCEL
)

// file.go
const (
	SeekSet = C.ALLEGRO_SEEK_SET
	SeekCur = C.ALLEGRO_SEEK_CUR
	SeekEnd = C.ALLEGRO_SEEK_END
)

// filesystem.go
const (
	FilemodeRead    = C.ALLEGRO_FILEMODE_READ
	FilemodeWrite   = C.ALLEGRO_FILEMODE_WRITE
	FilemodeExecute = C.ALLEGRO_FILEMODE_EXECUTE
	FilemodeHidden  = C.ALLEGRO_FILEMODE_HIDDEN
	FilemodeIsfile  = C.ALLEGRO_FILEMODE_ISFILE
	FilemodeIsdir   = C.ALLEGRO_FILEMODE_ISDIR
)

// fixed.go
const (
	FixToRad = int32(1608)
	RadToFix = int32(2670177)
)

// graphics.go
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

const (
	LockReadonly  = C.ALLEGRO_LOCK_READONLY
	LockWriteOnly = C.ALLEGRO_LOCK_WRITEONLY
	LockReadWrite = C.ALLEGRO_LOCK_READWRITE
)

const (
	VideoBitmap       = C.ALLEGRO_VIDEO_BITMAP
	MemoryBitmap      = C.ALLEGRO_MEMORY_BITMAP
	ConvertBitmap     = C.ALLEGRO_CONVERT_BITMAP
	ForceLocking      = C.ALLEGRO_FORCE_LOCKING
	NoPreserveTexture = C.ALLEGRO_NO_PRESERVE_TEXTURE
	MinLinear         = C.ALLEGRO_MIN_LINEAR
	MagLinear         = C.ALLEGRO_MAG_LINEAR
	Mipmap            = C.ALLEGRO_MIPMAP
)
const (
	FlipHorizontal = C.ALLEGRO_FLIP_HORIZONTAL
	FlipVertical   = C.ALLEGRO_FLIP_VERTICAL
)

const (
	Zero            = C.ALLEGRO_ZERO
	One             = C.ALLEGRO_ONE
	Alpha           = C.ALLEGRO_ALPHA
	InverseAlpha    = C.ALLEGRO_INVERSE_ALPHA
	SrcColor        = C.ALLEGRO_SRC_COLOR
	DstColor        = C.ALLEGRO_DST_COLOR
	InverseSrcColor = C.ALLEGRO_INVERSE_SRC_COLOR
	InverseDstColor = C.ALLEGRO_INVERSE_DST_COLOR
)

const (
	Add          = C.ALLEGRO_ADD
	SrcMinusDest = C.ALLEGRO_SRC_MINUS_DEST
	DestMinusSrc = C.ALLEGRO_DEST_MINUS_SRC
)

const (
	KeepIndex            = C.ALLEGRO_KEEP_INDEX
	KeepBitmapFormat     = C.ALLEGRO_KEEP_BITMAP_FORMAT
	NoPremultipliedAlpha = C.ALLEGRO_NO_PREMULTIPLIED_ALPHA
)

const (
	AlphaTest      = C.ALLEGRO_ALPHA_TEST
	WriteMask      = C.ALLEGRO_WRITE_MASK
	DepthTest      = C.ALLEGRO_DEPTH_TEST
	DepthFunction  = C.ALLEGRO_DEPTH_FUNCTION
	AlphaFunction  = C.ALLEGRO_ALPHA_FUNCTION
	AlphaTestValue = C.ALLEGRO_ALPHA_TEST_VALUE
)

const (
	RenderNever        = C.ALLEGRO_RENDER_NEVER
	RenderAlways       = C.ALLEGRO_RENDER_ALWAYS
	RenderLess         = C.ALLEGRO_RENDER_LESS
	RenderEqual        = C.ALLEGRO_RENDER_EQUAL
	RenderLessEqual    = C.ALLEGRO_RENDER_LESS_EQUAL
	RenderGreater      = C.ALLEGRO_RENDER_GREATER
	RenderNotEqual     = C.ALLEGRO_RENDER_NOT_EQUAL
	RenderGreaterEqual = C.ALLEGRO_RENDER_GREATER_EQUAL
)

const (
	MaskRed   = C.ALLEGRO_MASK_RED
	MaskGreen = C.ALLEGRO_MASK_GREEN
	MaskBlue  = C.ALLEGRO_MASK_BLUE
	MaskAlpha = C.ALLEGRO_MASK_ALPHA
	MaskDepth = C.ALLEGRO_MASK_DEPTH
	MaskRgb   = C.ALLEGRO_MASK_RGB
	MaskRgba  = C.ALLEGRO_MASK_RGBA
)

// joystick.go
const (
	JoyflagDigital  = C.ALLEGRO_JOYFLAG_DIGITAL
	JoyflagAnalogue = C.ALLEGRO_JOYFLAG_DIGITAL
)

// keyboard.go
const (
	KeyA             = C.ALLEGRO_KEY_A
	KeyB             = C.ALLEGRO_KEY_B
	KeyC             = C.ALLEGRO_KEY_C
	KeyD             = C.ALLEGRO_KEY_D
	KeyE             = C.ALLEGRO_KEY_E
	KeyF             = C.ALLEGRO_KEY_F
	KeyG             = C.ALLEGRO_KEY_G
	KeyH             = C.ALLEGRO_KEY_H
	KeyI             = C.ALLEGRO_KEY_I
	KeyJ             = C.ALLEGRO_KEY_J
	KeyK             = C.ALLEGRO_KEY_K
	KeyL             = C.ALLEGRO_KEY_L
	KeyM             = C.ALLEGRO_KEY_M
	KeyN             = C.ALLEGRO_KEY_N
	KeyO             = C.ALLEGRO_KEY_O
	KeyP             = C.ALLEGRO_KEY_P
	KeyQ             = C.ALLEGRO_KEY_Q
	KeyR             = C.ALLEGRO_KEY_R
	KeyS             = C.ALLEGRO_KEY_S
	KeyT             = C.ALLEGRO_KEY_T
	KeyU             = C.ALLEGRO_KEY_U
	KeyV             = C.ALLEGRO_KEY_V
	KeyW             = C.ALLEGRO_KEY_W
	KeyX             = C.ALLEGRO_KEY_X
	KeyY             = C.ALLEGRO_KEY_Y
	KeyZ             = C.ALLEGRO_KEY_Z
	Key0             = C.ALLEGRO_KEY_0
	Key1             = C.ALLEGRO_KEY_1
	Key2             = C.ALLEGRO_KEY_2
	Key3             = C.ALLEGRO_KEY_3
	Key4             = C.ALLEGRO_KEY_4
	Key5             = C.ALLEGRO_KEY_5
	Key6             = C.ALLEGRO_KEY_6
	Key7             = C.ALLEGRO_KEY_7
	Key8             = C.ALLEGRO_KEY_8
	Key9             = C.ALLEGRO_KEY_9
	KeyPad0          = C.ALLEGRO_KEY_PAD_0
	KeyPad1          = C.ALLEGRO_KEY_PAD_1
	KeyPad2          = C.ALLEGRO_KEY_PAD_2
	KeyPad3          = C.ALLEGRO_KEY_PAD_3
	KeyPad4          = C.ALLEGRO_KEY_PAD_4
	KeyPad5          = C.ALLEGRO_KEY_PAD_5
	KeyPad6          = C.ALLEGRO_KEY_PAD_6
	KeyPad7          = C.ALLEGRO_KEY_PAD_7
	KeyPad8          = C.ALLEGRO_KEY_PAD_8
	KeyPad9          = C.ALLEGRO_KEY_PAD_9
	KeyF1            = C.ALLEGRO_KEY_F1
	KeyF2            = C.ALLEGRO_KEY_F2
	KeyF3            = C.ALLEGRO_KEY_F3
	KeyF4            = C.ALLEGRO_KEY_F4
	KeyF5            = C.ALLEGRO_KEY_F5
	KeyF6            = C.ALLEGRO_KEY_F6
	KeyF7            = C.ALLEGRO_KEY_F7
	KeyF8            = C.ALLEGRO_KEY_F8
	KeyF9            = C.ALLEGRO_KEY_F9
	KeyF10           = C.ALLEGRO_KEY_F10
	KeyF11           = C.ALLEGRO_KEY_F11
	KeyF12           = C.ALLEGRO_KEY_F12
	KeyEscape        = C.ALLEGRO_KEY_ESCAPE
	KeyTilde         = C.ALLEGRO_KEY_TILDE
	KeyMinus         = C.ALLEGRO_KEY_MINUS
	KeyEquals        = C.ALLEGRO_KEY_EQUALS
	KeyBackspace     = C.ALLEGRO_KEY_BACKSPACE
	KeyTab           = C.ALLEGRO_KEY_TAB
	KeyOpenbrace     = C.ALLEGRO_KEY_OPENBRACE
	KeyClosebrace    = C.ALLEGRO_KEY_CLOSEBRACE
	KeyEnter         = C.ALLEGRO_KEY_ENTER
	KeySemicolon     = C.ALLEGRO_KEY_SEMICOLON
	KeyQuote         = C.ALLEGRO_KEY_QUOTE
	KeyBackslash     = C.ALLEGRO_KEY_BACKSLASH
	KeyBackslash2    = C.ALLEGRO_KEY_BACKSLASH2
	KeyComma         = C.ALLEGRO_KEY_COMMA
	KeyFullstop      = C.ALLEGRO_KEY_FULLSTOP
	KeySlash         = C.ALLEGRO_KEY_SLASH
	KeySpace         = C.ALLEGRO_KEY_SPACE
	KeyInsert        = C.ALLEGRO_KEY_INSERT
	KeyDelete        = C.ALLEGRO_KEY_DELETE
	KeyHome          = C.ALLEGRO_KEY_HOME
	KeyEnd           = C.ALLEGRO_KEY_END
	KeyPgup          = C.ALLEGRO_KEY_PGUP
	KeyPgdn          = C.ALLEGRO_KEY_PGDN
	KeyLeft          = C.ALLEGRO_KEY_LEFT
	KeyRight         = C.ALLEGRO_KEY_RIGHT
	KeyUp            = C.ALLEGRO_KEY_UP
	KeyDown          = C.ALLEGRO_KEY_DOWN
	KeyPadSlash      = C.ALLEGRO_KEY_PAD_SLASH
	KeyPadAsterisk   = C.ALLEGRO_KEY_PAD_ASTERISK
	KeyPadMinus      = C.ALLEGRO_KEY_PAD_MINUS
	KeyPadPlus       = C.ALLEGRO_KEY_PAD_PLUS
	KeyPadDelete     = C.ALLEGRO_KEY_PAD_DELETE
	KeyPadEnter      = C.ALLEGRO_KEY_PAD_ENTER
	KeyPrintscreen   = C.ALLEGRO_KEY_PRINTSCREEN
	KeyPause         = C.ALLEGRO_KEY_PAUSE
	KeyAbntC1        = C.ALLEGRO_KEY_ABNT_C1
	KeyYen           = C.ALLEGRO_KEY_YEN
	KeyKana          = C.ALLEGRO_KEY_KANA
	KeyConvert       = C.ALLEGRO_KEY_CONVERT
	KeyNoconvert     = C.ALLEGRO_KEY_NOCONVERT
	KeyAt            = C.ALLEGRO_KEY_AT
	KeyCircumflex    = C.ALLEGRO_KEY_CIRCUMFLEX
	KeyColon2        = C.ALLEGRO_KEY_COLON2
	KeyKanji         = C.ALLEGRO_KEY_KANJI
	KeyPadEquals     = C.ALLEGRO_KEY_PAD_EQUALS
	KeyBackquote     = C.ALLEGRO_KEY_BACKQUOTE
	KeySemicolon2    = C.ALLEGRO_KEY_SEMICOLON2
	KeyCommand       = C.ALLEGRO_KEY_COMMAND
	KeyLshift        = C.ALLEGRO_KEY_LSHIFT
	KeyRshift        = C.ALLEGRO_KEY_RSHIFT
	KeyLctrl         = C.ALLEGRO_KEY_LCTRL
	KeyRctrl         = C.ALLEGRO_KEY_RCTRL
	KeyAlt           = C.ALLEGRO_KEY_ALT
	KeyAltgr         = C.ALLEGRO_KEY_ALTGR
	KeyLwin          = C.ALLEGRO_KEY_LWIN
	KeyRwin          = C.ALLEGRO_KEY_RWIN
	KeyMenu          = C.ALLEGRO_KEY_MENU
	KeyScrolllock    = C.ALLEGRO_KEY_SCROLLLOCK
	KeyNumlock       = C.ALLEGRO_KEY_NUMLOCK
	KeyCapslock      = C.ALLEGRO_KEY_CAPSLOCK
	KeyMax           = C.ALLEGRO_KEY_MAX
	KeymodShift      = C.ALLEGRO_KEYMOD_SHIFT
	KeymodCtrl       = C.ALLEGRO_KEYMOD_CTRL
	KeymodAlt        = C.ALLEGRO_KEYMOD_ALT
	KeymodLwin       = C.ALLEGRO_KEYMOD_LWIN
	KeymodRwin       = C.ALLEGRO_KEYMOD_RWIN
	KeymodMenu       = C.ALLEGRO_KEYMOD_MENU
	KeymodAltgr      = C.ALLEGRO_KEYMOD_ALTGR
	KeymodCommand    = C.ALLEGRO_KEYMOD_COMMAND
	KeymodScrolllock = C.ALLEGRO_KEYMOD_SCROLLLOCK
	KeymodNumlock    = C.ALLEGRO_KEYMOD_NUMLOCK
	KeymodCapslock   = C.ALLEGRO_KEYMOD_CAPSLOCK
	KeymodInaltseq   = C.ALLEGRO_KEYMOD_INALTSEQ
	KeymodAccent1    = C.ALLEGRO_KEYMOD_ACCENT1
	KeymodAccent2    = C.ALLEGRO_KEYMOD_ACCENT2
	KeymodAccent3    = C.ALLEGRO_KEYMOD_ACCENT3
	KeymodAccent4    = C.ALLEGRO_KEYMOD_ACCENT4
)

// mouse.go
const (
	SystemMouseCursorNone        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_NONE
	SystemMouseCursorDefault     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_DEFAULT
	SystemMouseCursorArrow       = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ARROW
	SystemMouseCursorBusy        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_BUSY
	SystemMouseCursorQuestion    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_QUESTION
	SystemMouseCursorEdit        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_EDIT
	SystemMouseCursorMove        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_MOVE
	SystemMouseCursorResizeN     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_N
	SystemMouseCursorResizeW     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_W
	SystemMouseCursorResizeS     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_S
	SystemMouseCursorResizeE     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_E
	SystemMouseCursorResizeNW    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NW
	SystemMouseCursorResizeSW    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SW
	SystemMouseCursorResizeSE    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SE
	SystemMouseCursorResizeNE    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NE
	SystemMouseCursorProgress    = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PROGRESS
	SystemMouseCursorPrecision   = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PRECISION
	SystemMouseCursorLink        = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_LINK
	SystemMouseCursorSelect      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ALT_SELECT
	SystemMouseCursorUnavailable = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_UNAVAILABLE
)

// shader.go
const (
	VertexShader = C.ALLEGRO_VERTEX_SHADER
	PixelShader  = C.ALLEGRO_PIXEL_SHADER
)

const (
	ShaderAuto = C.ALLEGRO_SHADER_AUTO
	ShaderGlsl = C.ALLEGRO_SHADER_GLSL
	ShaderHlsl = C.ALLEGRO_SHADER_HLSL
)

// state.go
const (
	StateNewDisplayParameters = C.ALLEGRO_STATE_NEW_DISPLAY_PARAMETERS
	StateNewBitmapParameters  = C.ALLEGRO_STATE_NEW_BITMAP_PARAMETERS
	StateDisplay              = C.ALLEGRO_STATE_DISPLAY
	StateBlender              = C.ALLEGRO_STATE_BLENDER
	StteTransform             = C.ALLEGRO_STATE_TRANSFORM
	StateNewFileInterface     = C.ALLEGRO_STATE_NEW_FILE_INTERFACE
	StateBitmap               = C.ALLEGRO_STATE_BITMAP
	StateTargetBitmap         = C.ALLEGRO_STATE_TARGET_BITMAP
	StateAll                  = C.ALLEGRO_STATE_ALL
)

// touchinput.go
const TouchInputMaxTouchCount = C.ALLEGRO_TOUCH_INPUT_MAX_TOUCH_COUNT

const (
	MouseEmulationNone        = C.ALLEGRO_MOUSE_EMULATION_NONE
	MouseEmulationTransparent = C.ALLEGRO_MOUSE_EMULATION_TRANSPARENT
	MouseEmulationInclusive   = C.ALLEGRO_MOUSE_EMULATION_INCLUSIVE
	MouseEmulationExclusive   = C.ALLEGRO_MOUSE_EMULATION_EXCLUSIVE
	MouseEmulation50X         = C.ALLEGRO_MOUSE_EMULATION_5_0_x
)

// opengl.go
const (
	DesktopOpengl = C.ALLEGRO_DESKTOP_OPENGL
	OpenglEs      = C.ALLEGRO_OPENGL_ES
)
