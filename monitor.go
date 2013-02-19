package allegro

// #include <allegro5/allegro.h>
import "C"

import "unsafe"

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

func GetMonitorInfo(adapter int) *MonitorInfo {
	var m C.ALLEGRO_MONITOR_INFO
	b := bool(C.al_get_monitor_info(C.int(adapter), &m))
	if b == false {
		return nil
	}
	return (*MonitorInfo)(unsafe.Pointer(&m))
}

func GetNumVideoAdapters() int {
	return int(C.al_get_num_video_adapters())
}
