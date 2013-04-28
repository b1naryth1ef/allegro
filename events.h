#include <allegro5/allegro.h>

ALLEGRO_EVENT_TYPE GetEventType(ALLEGRO_EVENT *str) {
	return str->type;
}

/*******/
/* Any */
/*******/

ALLEGRO_EVENT_TYPE GetAnyType(ALLEGRO_EVENT *str) {
	return str->any.type;
}

ALLEGRO_EVENT_SOURCE *GetAnySource(ALLEGRO_EVENT *str) {
	return str->any.source;
}

double GetAnyTimestamp(ALLEGRO_EVENT *str) {
	return str->any.timestamp;
}

/***********/
/* Display */
/***********/

ALLEGRO_EVENT_TYPE GetDisplayType(ALLEGRO_EVENT *str) {
	return str->display.type;
}

ALLEGRO_DISPLAY *GetDisplaySource(ALLEGRO_EVENT *str) {
	return str->display.source;
}

double GetDisplayTimestamp(ALLEGRO_EVENT *str) {
	return str->display.timestamp;
}

int GetDisplayX(ALLEGRO_EVENT *str) {
	return str->display.x;
}

int GetDisplayY(ALLEGRO_EVENT *str) {
	return str->display.y;
}

int GetDisplayWidth(ALLEGRO_EVENT *str) {
	return str->display.width;
}

int GetDisplayHeight(ALLEGRO_EVENT *str) {
	return str->display.height;
}

int GetDisplayOrientation(ALLEGRO_EVENT *str) {
	return str->display.orientation;
}

/************/
/* Joystick */
/************/

ALLEGRO_EVENT_TYPE GetJoystickType(ALLEGRO_EVENT *str) {
	return str->joystick.type;
}

ALLEGRO_JOYSTICK *GetJoystickSource(ALLEGRO_EVENT *str) {
	return str->joystick.source;
}

double GetJoystickTimestamp(ALLEGRO_EVENT *str) {
	return str->joystick.timestamp;
}

ALLEGRO_JOYSTICK *GetJoystickId(ALLEGRO_EVENT *str) {
	return str->joystick.id;
}

int GetJoystickStick(ALLEGRO_EVENT *str) {
	return str->joystick.stick;
}

int GetJoystickAxis(ALLEGRO_EVENT *str) {
	return str->joystick.axis;
}

float GetJoystickPos(ALLEGRO_EVENT *str) {
	return str->joystick.pos;
}

int GetJoystickButton(ALLEGRO_EVENT *str) {
	return str->joystick.button;
}

/************/
/* Keyboard */
/************/

ALLEGRO_EVENT_TYPE GetKeyboardType(ALLEGRO_EVENT *str) {
	return str->keyboard.type;
}

ALLEGRO_KEYBOARD *GetKeyboardSource(ALLEGRO_EVENT *str) {
	return str->keyboard.source;
}

double GetKeyboardTimestamp(ALLEGRO_EVENT *str) {
	return str->keyboard.timestamp;
}

ALLEGRO_DISPLAY *GetKeyboardDisplay(ALLEGRO_EVENT *str) {
	return str->keyboard.display;
}

int GetKeyboardKeycode(ALLEGRO_EVENT *str) {
	return str->keyboard.keycode;
}

int GetKeyboardUnichar(ALLEGRO_EVENT *str) {
	return str->keyboard.unichar;
}

unsigned int GetKeyboardModifiers(ALLEGRO_EVENT *str) {
	return str->keyboard.modifiers;
}

bool GetKeyboardRepeat(ALLEGRO_EVENT *str) {
	return str->keyboard.repeat;
}

/*********/
/* Mouse */
/*********/

ALLEGRO_EVENT_TYPE GetMouseType(ALLEGRO_EVENT *str) {
	return str->mouse.type;
}

ALLEGRO_MOUSE *GetMouseSource(ALLEGRO_EVENT *str) {
	return str->mouse.source;
}

double GetMouseTimestamp(ALLEGRO_EVENT *str) {
	return str->mouse.timestamp;
}

ALLEGRO_DISPLAY *GetMouseDisplay(ALLEGRO_EVENT *str) {
	return str->mouse.display;
}

int GetMouseX(ALLEGRO_EVENT *str) {
	return str->mouse.x;
}

int GetMouseY(ALLEGRO_EVENT *str) {
	return str->mouse.y;
}

int GetMouseZ(ALLEGRO_EVENT *str) {
	return str->mouse.z;
}

int GetMouseW(ALLEGRO_EVENT *str) {
	return str->mouse.w;
}

int GetMouseDx(ALLEGRO_EVENT *str) {
	return str->mouse.dx;
}

int GetMouseDy(ALLEGRO_EVENT *str) {
	return str->mouse.dy;
}

int GetMouseDz(ALLEGRO_EVENT *str) {
	return str->mouse.dz;
}

int GetMouseDw(ALLEGRO_EVENT *str) {
	return str->mouse.dw;
}

unsigned int GetMouseButton(ALLEGRO_EVENT *str) {
	return str->mouse.button;
}

float GetMousePressure(ALLEGRO_EVENT *str) {
	return str->mouse.pressure;
}

/*********/
/* Timer */
/*********/

ALLEGRO_EVENT_TYPE GetTimerType(ALLEGRO_EVENT *str) {
	return str->timer.type;
}

ALLEGRO_TIMER *GetTimerSource(ALLEGRO_EVENT *str) {
	return str->timer.source;
}

double GetTimerTimestamp(ALLEGRO_EVENT *str) {
	return str->timer.timestamp;
}

int64_t GetTimerCount(ALLEGRO_EVENT *str) {
	return str->timer.count;
}

double GetTimerError(ALLEGRO_EVENT *str) {
	return str->timer.error;
}

/*********/
/* Touch */
/*********/

ALLEGRO_EVENT_TYPE GetTouchType(ALLEGRO_EVENT *str) {
	return str->touch.type;
}

ALLEGRO_TOUCH_INPUT *GetTouchSource(ALLEGRO_EVENT *str) {
	return str->touch.source;
}

double GetTouchTimestamp(ALLEGRO_EVENT *str) {
	return str->touch.timestamp;
}

ALLEGRO_DISPLAY *GetTouchDisplay(ALLEGRO_EVENT *str) {
	return str->touch.display;
}

int GetTouchId(ALLEGRO_EVENT *str) {
	return str->touch.id;
}

float GetTouchX(ALLEGRO_EVENT *str) {
	return str->touch.x;
}

float GetTouchY(ALLEGRO_EVENT *str) {
	return str->touch.y;
}

float GetTouchDx(ALLEGRO_EVENT *str) {
	return str->touch.dx;
}

float GetTouchDy(ALLEGRO_EVENT *str) {
	return str->touch.dy;
}

bool GetTouchPrimary(ALLEGRO_EVENT *str) {
	return str->touch.primary;
}

/*************/
/* UserEvent */
/*************/

ALLEGRO_EVENT_TYPE GetUserType(ALLEGRO_EVENT *str) {
	return str->user.type;
}

ALLEGRO_EVENT_SOURCE *GetUserSource(ALLEGRO_EVENT *str) {
	return str->user.source;
}

double GetUserTimestamp(ALLEGRO_EVENT *str) {
	return str->user.timestamp;
}

intptr_t GetUserData1(ALLEGRO_EVENT *str) {
	return str->user.data1;
}

intptr_t GetUserData2(ALLEGRO_EVENT *str) {
	return str->user.data2;
}

intptr_t GetUserData3(ALLEGRO_EVENT *str) {
	return str->user.data3;
}

intptr_t GetUserData4(ALLEGRO_EVENT *str) {
	return str->user.data4;
}
