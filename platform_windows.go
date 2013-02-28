package allegro

// #include <allegro5/allegro.h>
import "C"

func (d *Display) GetWinWindowHandle() C.HWND {
	return C.al_get_win_window_handle((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(d)))
}

//func (d *Display) WinAddWindowCallback
//func (d *Display) WinRemoveWindowCallback
