package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

const (
	JoyflagDigital  = C.ALLEGRO_JOYFLAG_DIGITAL
	JoyflagAnalogue = C.ALLEGRO_JOYFLAG_DIGITAL
)

type Joystick C.ALLEGRO_JOYSTICK

type JoystickState C.ALLEGRO_JOYSTICK_STATE

func InstallJoystick() bool {
	return bool(C.al_install_joystick())
}

func UninstallJoystick() {
	C.al_uninstall_joystick()
}

func IsJoystickInstalled() bool {
	return bool(C.al_is_joystick_installed())
}

func ReconfigureJoysticks() bool {
	return bool(C.al_reconfigure_joysticks())
}

func GetNumJoysticks() int32 {
	return int32(C.al_get_num_joysticks())
}

func GetJoystick(num int32) *Joystick {
	return (*Joystick)(unsafe.Pointer(C.al_get_joystick(C.int(num))))
}

func (j *Joystick) Release() {
	C.al_release_joystick((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j)))
}

func (j *Joystick) GetActive() bool {
	return bool(C.al_get_joystick_active((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j))))
}

func (j *Joystick) GetName() string {
	return C.GoString(C.al_get_joystick_name((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j))))
}

func (j *Joystick) GetStickName(stick int32) string {
	return C.GoString(C.al_get_joystick_stick_name((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j)), C.int(stick)))
}

func (j *Joystick) GetAxisName(stick, axis int32) string {
	return C.GoString(C.al_get_joystick_axis_name((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j)), C.int(stick), C.int(axis)))
}

func (j *Joystick) GetButtonName(button int32) string {
	return C.GoString(C.al_get_joystick_button_name((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j)), C.int(button)))
}

func (j *Joystick) GetStickFlags(stick int32) int32 {
	return int32(C.al_get_joystick_stick_flags((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j)), C.int(stick)))
}

func (j *Joystick) GetNumSticks() int32 {
	return int32(C.al_get_joystick_num_sticks((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j))))
}

func (j *Joystick) GetNumAxes(stick int32) int32 {
	return int32(C.al_get_joystick_num_axes((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j)), C.int(stick)))
}

func (j *Joystick) GetNumButtons() int32 {
	return int32(C.al_get_joystick_num_buttons((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j))))
}

func (j *Joystick) GetJoystickState() *JoystickState {
	js := new(C.ALLEGRO_JOYSTICK_STATE)
	C.al_get_joystick_state((*C.ALLEGRO_JOYSTICK)(unsafe.Pointer(j)), js)
	return (*JoystickState)(unsafe.Pointer(js))
}

func GetJoystickEventSource() *EventSource {
	return (*EventSource)(unsafe.Pointer(C.al_get_joystick_event_source()))
}
